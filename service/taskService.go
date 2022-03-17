package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/orange-protocol/orange-server-v1/config"
	"github.com/orange-protocol/orange-server-v1/did"
	"github.com/orange-protocol/orange-server-v1/log"
	"github.com/orange-protocol/orange-server-v1/provider"
	"github.com/orange-protocol/orange-server-v1/provider/common"
	"github.com/orange-protocol/orange-server-v1/store"
	"github.com/orange-protocol/orange-server-v1/wasm/types"
)

const (
	TASK_TICKER      = 10 * time.Second
	OSCORE_AP_METHOD = "calc30xWithDefi"
)

type TaskService struct {
	exitChan chan int
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (ts *TaskService) Run() {
	ticker := time.NewTicker(TASK_TICKER)
	for {
		select {
		case <-ticker.C:
			go ts.monitorNewTasks()
			go ts.MonitorDPFinishedTasks()
			go ts.monitorAPFinishedTasks()
			go ts.resovleTimeoutTasks()
			//todo monitor failed tasks & retry
		case <-ts.exitChan:
			return
		}

	}
}

func (ts *TaskService) resovleTimeoutTasks() {
	inqueryCnt, err := store.MySqlDB.GetInQueryTaskCount()
	if err != nil {
		log.Errorf("errors on GetInQueryTaskCount:%s", err.Error())
		return
	}
	cnt := config.GlobalConfig.BatchTaskCount - int(inqueryCnt)
	if cnt == 0 {
		return
	}
	err = store.MySqlDB.ResetTimeOutTasks(cnt)
	if err != nil {
		log.Errorf("ResetTimeOutTasks failed!%s", err.Error())
		return
	}

}

//FSM init -> solving
func (ts *TaskService) monitorNewTasks() {
	//todo add server id to lock the tasks
	inqueryCnt, err := store.MySqlDB.GetInQueryTaskCount()
	if err != nil {
		log.Errorf("errors on GetInQueryTaskCount:%s", err.Error())
		return
	}
	if inqueryCnt >= int64(config.GlobalConfig.BatchTaskCount) {
		return
	}

	_, err = store.MySqlDB.LockTaskStatus(store.TASK_STATUS_INIT, store.TASK_STATUS_RESOLVING, int64(config.GlobalConfig.BatchTaskCount)-inqueryCnt)
	if err != nil {
		log.Errorf("errors on LockTaskStatus:%s", err.Error())
		return
	}

	tasks, err := store.MySqlDB.QueryTaskByStatus(store.TASK_STATUS_RESOLVING)
	if err != nil {
		log.Errorf("errors on QueryTaskByStatus:%s", err.Error())
		return
	}

	for _, task := range tasks {
		//solving -> DP starting
		l, err := store.MySqlDB.ChangeTaskStatus(task.TaskId, store.TASK_STATUS_RESOLVING, store.TASK_STATUS_DP_QUERYING)
		if err != nil {
			log.Errorf("errors on QueryTaskByStatus:%s", err.Error())
			continue
		}
		if l != 1 {
			log.Debugf("fail to ChangeTaskStatus:taskid:%d", task.TaskId)
			continue
		}
		go ts.resovleDataProvider(task)
	}
}

func (ts *TaskService) resovleDataProvider(task *store.TaskInfo) error {
	dm, err := store.MySqlDB.QueryDPMethodByDIDAndMethod(task.DpDID, task.DpMethod)
	if err != nil {
		log.Errorf("errors on QueryDPMethodByDIDAndMethod did:%s", task.DpDID)
		return err
	}
	if dm == nil {
		return store.MySqlDB.SetDPResultFailed(task.TaskId, fmt.Sprintf("no dp method with did:%s,name:%s found ", task.DpDID, task.DpMethod))
	}
	if dm.CompositeSetting != "NONE" {
		err = ts.DealWithCompositeDP(dm, task)
	} else {
		err = ts.DealWithSingleDP(dm.Did, dm.Method, false, false, task)
	}
	if err != nil {
		return store.MySqlDB.SetDPResultFailed(task.TaskId, err.Error())
	}
	return nil
}

func (ts *TaskService) getParamString(userdid, dpdid, dpmethod string, isPOC bool, apdid string, dptype int, bind_info string) (string, string, error) {

	//param := ""

	dpm, err := store.MySqlDB.QueryDPMethodByDIDAndMethod(dpdid, dpmethod)
	if err != nil {
		return "", "", err
	}
	if dpm == nil {
		return "", "", fmt.Errorf("no dp method found with did:%s,method:%s", dpdid, dpmethod)
	}

	dp, err := store.MySqlDB.QueryDataProviderByDid(dpdid)
	if err != nil {
		return "", "", err
	}
	if dp == nil {
		return "", "", fmt.Errorf("no dp found with did:%s", dpdid)
	}
	param := dpm.Param

	var userAddressinfo []*store.UserAddressInfo
	if len(bind_info) == 0 {
		arr := strings.Split(bind_info, ",")
		for _, addr := range arr {
			uai := &store.UserAddressInfo{Chain: "eth", Address: addr}
			userAddressinfo = append(userAddressinfo, uai)
		}
	} else {
		userAddressinfo, err = store.MySqlDB.GetUserVisibleAddressInfo(userdid)
		if err != nil {
			return "", "", err
		}
		if len(userAddressinfo) == 0 {
			if strings.HasPrefix(userdid, "did:etho:") {
				userAddressinfo = append(userAddressinfo, &store.UserAddressInfo{
					Did:        userdid,
					Chain:      "eth",
					Address:    strings.ReplaceAll(userdid, "did:etho:", "0x"),
					Pubkey:     "",
					CreateTime: time.Now(),
					Visible:    true,
				})
			} else {
				return "", "", fmt.Errorf("no address info found under user did:%s", userdid)
			}
		}
	}

	//return ParseInputParam(param, dp.Apikey, userAddressinfo, isPOC, externalDID)
	return ParseInputParam(param, dp.Apikey, userAddressinfo, isPOC, apdid, dptype, dpm.HttpMethod)
}

func (ts *TaskService) DealWithSingleDP(dpdid, dpmethod string, isPOC bool, isLast bool, task *store.TaskInfo) error {

	dp, err := provider.GetDataProvider(dpdid)
	if err != nil {
		log.Errorf("errors on GetDataProvider did:%s,err:%s", task.DpDID, err.Error())
		return err
	}
	if dp == nil {
		return fmt.Errorf("no dp with did:%s found ", task.DpDID)
	}
	dpinfo, err := store.MySqlDB.QueryDataProviderByDid(dpdid)
	if err != nil {
		log.Errorf("errors on QueryDataProviderByDid did:%s", task.DpDID)
		return err
	}

	paramStr, comments, err := ts.getParamString(task.UserDID, dpdid, dpmethod, isPOC, task.ApDID, dpinfo.DpType, task.TaskBindInfo)
	if err != nil {
		log.Errorf("errors on getParamString:%s", err.Error())
		return err
	}
	log.Debugf("paramStr:%s\n", paramStr)
	if len(comments) > 0 {
		err := store.MySqlDB.SetTaskComments(task.TaskId, comments)
		if err != nil {
			log.Errorf("errors on SetTaskComments:%s", err.Error())
			return err
		}
	}

	t := time.Now()
	log.Debugf("== start to invoke dp:%s,method:%s\n", dpdid, dpmethod)
	//res, err := dp.InvokeMethod(dpmethod, paramMap)
	res, err := dp.InvokeMethodWithParamStr(dpmethod, paramStr)
	if err != nil {
		log.Errorf("errors on InvokeMethod err:%s", err.Error())
		err2 := store.MySqlDB.SetDPResultFailed(task.TaskId, err.Error())
		if err2 != nil {
			log.Errorf("errors on SetDPResultFailed:%d,err:%s,err2:%s", task.TaskId, err.Error(), err2.Error())
		}
		return err
	}
	latency := time.Now().Unix() - t.Unix()
	log.Debugf("== end to invoke dp:%s,method:%s,cost:%d\n", dpdid, dpmethod, latency)
	err = store.MySqlDB.UpdateDPMethodLatency(dpdid, latency)
	if err != nil {
		log.Errorf("errors on UpdateDPMethodLatency")
	}
	err = store.MySqlDB.UpdateDPMethodInvoked(dpdid, dpmethod, latency)
	if err != nil {
		log.Errorf("errors on UpdateDPMethodInvoked:%s", err.Error())
	}
	resdata := ""
	if dpinfo.DpType == common.DP_TYPE_CUSTOM {
		resdata, err = ts.solveCustomDPRes(res)
		if err != nil {
			log.Errorf("errors on solveCustomDPRes:%s,res:%s", err.Error(), res)
			return err
		}
	} else {
		resdata, err = ts.solveDPRes(res, dpmethod)
		if err != nil {
			log.Errorf("errors on solveXdaysRes:%s,res:%s", err.Error(), res)
			return err
		}
	}

	if isPOC {
		//append DP Result
		err = store.MySqlDB.AppendDPResult(task.TaskId, resdata, isLast, dpmethod)
		if err != nil {
			log.Errorf("errors on AppendDPResult err:%s", err.Error())
			return err
		}
	} else {
		err = store.MySqlDB.SetDPResult(task.TaskId, resdata)
		if err != nil {
			log.Errorf("errors on SetDPResult err:%s", err.Error())
			return err
		}
		err = store.MySqlDB.AddTaskHistoryByTaskInfo(task)
		if err != nil {
			log.Errorf("errors on AddTaskHistoryByTaskInfo err:%s", err.Error())
		}
		return err
	}

	return nil
}

func (ts *TaskService) DealWithCompositeDP(dm *store.DPMethod, task *store.TaskInfo) error {
	setting := dm.CompositeSetting
	//setting should be "did#method;did#method"

	arr := strings.Split(setting, ";")
	t := time.Now()
	for i, s := range arr {
		tmp := strings.Split(s, "#")
		if len(tmp) != 2 {
			return fmt.Errorf("composite setting format error")
		}
		did, method := tmp[0], tmp[1]
		isLast := false
		if i == len(arr)-1 {
			isLast = true
		}

		err := ts.DealWithSingleDP(did, method, true, isLast, task)
		if err != nil {
			log.Errorf("DealWithSingleDP failed:%s", err.Error())
			return err
		}
	}
	latency := time.Now().Unix() - t.Unix()

	err := store.MySqlDB.UpdateDPMethodLatency(task.DpDID, latency)
	if err != nil {
		log.Errorf("UpdateDPMethodLatency failed:%s", err.Error())
		//return err
	}
	err = store.MySqlDB.UpdateDPMethodInvoked(task.DpDID, task.DpMethod, latency)
	if err != nil {
		log.Errorf("errors on UpdateDPMethodInvoked %s", err.Error())
	}
	return nil
}

func (ts *TaskService) MonitorDPFinishedTasks() {
	_, err := store.MySqlDB.LockTaskStatus(store.TASK_STATUS_DP_FINISHED, store.TASK_STATUS_AP_RESOLVING, int64(config.GlobalConfig.BatchTaskCount))
	if err != nil {
		log.Errorf("errors on LockTaskStatus:%s", err.Error())
		return
	}

	//if cnt > 0 {
	tasks, err := store.MySqlDB.QueryTaskByStatus(store.TASK_STATUS_AP_RESOLVING)
	if err != nil {
		log.Errorf("errors on QueryTaskByStatus:%s", err.Error())
		return
	}

	for _, task := range tasks {
		l, err := store.MySqlDB.ChangeTaskStatus(task.TaskId, store.TASK_STATUS_AP_RESOLVING, store.TASK_STATUS_AP_QUERYING)

		if err != nil {
			log.Errorf("errors on QueryTaskByStatus:%s", err.Error())
			continue
		}
		if l != 1 {
			log.Debugf("fail to ChangeTaskStatus:taskid:%d from:%d, to%d", task.TaskId, store.TASK_STATUS_AP_RESOLVING, store.TASK_STATUS_AP_QUERYING)
			continue
		}
		go ts.resolveAlgorithmProvider(task)
	}

	//}
}

func (ts *TaskService) monitorAPFinishedTasks() {
	_, err := store.MySqlDB.LockTaskStatus(store.TASK_STATUS_AP_FINISHED, store.TASK_STATUS_VC_STARTING, int64(config.GlobalConfig.BatchTaskCount))
	if err != nil {
		log.Errorf("errors on LockTaskStatus:%s", err.Error())
		return
	}
	//if cnt > 0 {
	tasks, err := store.MySqlDB.QueryTaskByStatus(store.TASK_STATUS_VC_STARTING)
	if err != nil {
		log.Errorf("errors on QueryTaskByStatus:%s", err.Error())
		return
	}

	for _, task := range tasks {
		l, err := store.MySqlDB.ChangeTaskStatus(task.TaskId, store.TASK_STATUS_VC_STARTING, store.TASK_STATUS_VC_GENERATING)
		if err != nil {
			log.Errorf("errors on QueryTaskByStatus:%s", err.Error())
			continue
		}
		if l != 1 {
			log.Debugf("fail to ChangeTaskStatus:taskid:%d from:%d, to%d", task.TaskId, store.TASK_STATUS_AP_RESOLVING, store.TASK_STATUS_AP_QUERYING)
			continue
		}

		go ts.CreateCredential(task)

		//update latest oscore
		if task.ApMethod == OSCORE_AP_METHOD {
			latestscore, err := store.MySqlDB.QueryUserLatestOScoreInfo(task.UserDID)
			if err != nil {
				log.Errorf("fail to QueryUserLatestOScoreInfo:err:%s", err.Error())
				continue
			}
			score, err := strconv.Atoi(task.TaskResult)
			if err != nil {
				log.Errorf("fail to Atoi:result:%s,err:%s", task.TaskResult, err.Error())
				continue
			}
			param := &store.UserOscoreInfo{
				Did:        task.UserDID,
				Oscore:     score,
				ApDid:      task.ApDID,
				DpDid:      task.DpDID,
				CreateTime: time.Now(),
			}
			if latestscore == nil {
				err = store.MySqlDB.AddUserOScoreInfo(param)
			} else {
				err = store.MySqlDB.UpdateUserLatestOscore(param)
			}
			if err != nil {
				log.Errorf("fail to UpdateUserLatestOscore,did:%s,err:%s", task.UserDID, err.Error())
				continue
			}
		}
	}
}

func (ts *TaskService) resolveAlgorithmProvider(task *store.TaskInfo) error {

	//todo verify json schema ??
	ap, err := provider.GetAlgorithmProvider(task.ApDID)
	if err != nil {
		log.Errorf("errors on GetAlgorithmProvider:%s,%s", task.ApDID, err.Error())
		return nil
	}

	/*	fmt.Printf("resolveAlgorithmProvider:===dpresult:%s\n",task.DpResult)

		dpm ,err := store.MySqlDB.GetDPMethodInfo(task.DpDID,task.DpMethod)
		if err != nil {
			return err
		}
		if len(dpm)!= 1 {
			return fmt.Errorf("no dpmethod found for task:%d",task.TaskId)
		}

		if !strings.EqualFold(dpm[0].CompositeSetting , "NONE"){
			tmpArr := strings.Split(task.DpMethod,";;")
			m:= make(map[string]string)
			for _,t := range tmpArr {

			}
		}*/

	paramMap := make(map[string]interface{})
	paramMap["%input"] = task.DpResult
	t := time.Now()
	res, err := ap.Invoke(task.ApMethod, paramMap)
	if err != nil {
		log.Errorf("errors on Invoke:%s,err:%s", task.ApMethod, err.Error())
		store.MySqlDB.SetAPResultFailed(task.TaskId, err.Error())
		return err
	}

	err = store.MySqlDB.AddAPInvokeFrenquency(task.ApDID, time.Now().Unix()-t.Unix())
	if err != nil {
		log.Errorf("AddAPInvokeFrenquency failed:%s", err.Error())
	}
	err = store.MySqlDB.AddAPMethodInvoked(task.ApDID, task.ApMethod, time.Now().Unix()-t.Unix())
	if err != nil {
		log.Errorf("AddAPMethodInvoked failed:%s", err.Error())
	}

	apinfo, err := store.MySqlDB.QueryAlgorithmProviderByDid(task.ApDID)
	if err != nil {
		log.Errorf("errors on QueryAlgorithmProviderByDid:%s,err:%s", task.ApDID, err.Error())
		return err
	}
	apresult := ""
	if apinfo.APType == common.AP_TYPE_OUTER {
		m := make(map[string]interface{})
		err = json.Unmarshal(res.([]byte), &m)
		if err != nil {
			store.MySqlDB.SetAPResultFailed(task.TaskId, err.Error())
			return err
		}
		if es, ok := m["error"]; ok {
			store.MySqlDB.SetAPResultFailed(task.TaskId, es.(string))
			return err
		}
		apresult = fmt.Sprintf("%d", int64(m["score"].(float64)))
	} else {
		score := res.(*types.ScoreResult)
		apresult = fmt.Sprintf("%d", score.Score)
	}
	err = store.MySqlDB.SetAPResult(task.TaskId, apresult)
	if err != nil {
		log.Errorf("errors on SetAPResult:%s,err:%s", task.ApMethod, err.Error())
		return err
	}

	//task.TaskResult = fmt.Sprintf("%d", apresult)
	task.TaskResult = apresult
	task.TaskStatus = store.TASK_STATUS_AP_FINISHED

	err = store.MySqlDB.UpdateTaskHistoryByTaskInfo(task)
	if err != nil {
		log.Errorf("errors on UpdateTaskHistoryByTaskInfo,err:%s", err.Error())
		return err
	}
	return nil
}

func (ts *TaskService) CreateCredential(task *store.TaskInfo) error {
	apm, err := store.MySqlDB.QueryAPMethodByDIDAndMethod(task.ApDID, task.ApMethod)
	if err != nil {
		log.Errorf("QueryAPMethodByDIDAndMethod error:%s", err.Error())
		return err
	}
	dpm, err := store.MySqlDB.QueryDPMethodByDIDAndMethod(task.DpDID, task.DpMethod)
	if err != nil {
		log.Errorf("QueryDPMethodByDIDAndMethod error:%s", err.Error())
		return err
	}

	apMethodName, dpMethodName := "", ""
	if apm != nil {
		apMethodName = apm.Name
	}
	if dpm != nil {
		dpMethodName = dpm.Name
	}

	cred := &did.OrangeCredential{
		Data: &did.DataField{
			Provider_did: task.DpDID,
			Method:       task.DpMethod,
			Name:         dpMethodName,
			UserDid:      task.UserDID,
			BindData:     ts.getBindData(task),
			//Data:         task.DpResult,
		},
		Algorithm: &did.AlgorithmField{
			Provider_did: task.ApDID,
			Method:       task.ApMethod,
			Name:         apMethodName,
			Result:       task.TaskResult,
		},
	}
	credbts, err := json.Marshal(cred)
	if err != nil {
		log.Errorf("CreateCredential error:%s", err.Error())
		return err
	}

	commit := config.GlobalConfig.DidConf[0].Commit

	credstr, txhash, err := SysDidService.IssueCredential(config.GlobalConfig.DidConf[0].DID, string(credbts), commit)
	if err != nil {
		log.Errorf("CreateCredential error:%s", err.Error())
		return err
	}

	fileName := fmt.Sprintf("%d.cred", task.TaskId)
	err = ioutil.WriteFile("./files/"+fileName, []byte(credstr), 0644)
	if err != nil {
		log.Errorf("CreateCredential error:%s", err.Error())
		return err
	}

	err = store.MySqlDB.SetTaskCredFileAndTxhash(task.TaskId, fileName, txhash)
	if err != nil {
		log.Errorf("errors on SetTaskCredFileAndTxhash err:%s", err.Error())
		return err
	}

	err = store.MySqlDB.UpdateTaskHistoryTxhash(task.TaskId, txhash)
	if err != nil {
		log.Errorf("errors on UpdateTaskHistoryTxhash err:%s", err.Error())
		return err
	}

	return err
}

func (ts *TaskService) getBindData(task *store.TaskInfo) string {
	//todo ,add fix code
	ethoPre := "did:etho:"
	if strings.HasPrefix(task.UserDID, ethoPre) {
		return strings.ReplaceAll(task.UserDID, ethoPre, "0x")
	}
	return ""
}

func (ts *TaskService) solveXdaysRes(in []byte, method string) (string, error) {
	m := make(map[string]interface{})

	err := json.Unmarshal(in, &m)
	if err != nil {
		return "", err
	}
	if strings.Contains("errors", string(in)) {
		return "", fmt.Errorf(string(in))
	}

	if _, ok := m["errors"]; ok {
		log.Errorf("solveXdaysRes:error:%s", in)
		return "", fmt.Errorf(string(in))
	}

	t, ok := m["data"]
	if !ok {
		return "", fmt.Errorf("error result format no 'data' field")
	}

	if reflect.TypeOf(t).Kind() != reflect.Map {
		return "", fmt.Errorf("error result format 'data' field error")
	}
	t2, ok := t.(map[string]interface{})[method]
	if !ok {
		return "", fmt.Errorf("error result format no '%s' field", method)
	}

	res := ""
	if t2.(map[string]interface{})["encrypted"] != nil && len(t2.(map[string]interface{})["encrypted"].(string)) > 0 {
		res = t2.(map[string]interface{})["encrypted"].(string)
	} else {
		data := t2.(map[string]interface{})["data"]
		b, err := json.Marshal(data)
		if err != nil {
			return "", fmt.Errorf("json marshal failed")
		}
		res = string(b)
	}
	return res, nil
}
func (ts *TaskService) solveCustomDPRes(in []byte) (string, error) {
	m := make(map[string]interface{})

	err := json.Unmarshal(in, &m)
	if err != nil {
		return "", err
	}
	if strings.Contains("error", string(in)) {
		return "", fmt.Errorf(string(in))
	}

	if _, ok := m["error"]; ok {
		log.Errorf("solveCustomDPRes:error:%s", in)
		return "", fmt.Errorf(string(in))
	}
	return string(in), nil
}

func (ts *TaskService) solveDPRes(in []byte, method string) (string, error) {
	m := make(map[string]interface{})

	err := json.Unmarshal(in, &m)
	if err != nil {
		return "", err
	}
	if strings.Contains("errors", string(in)) {
		return "", fmt.Errorf(string(in))
	}

	if _, ok := m["errors"]; ok {
		log.Errorf("solveDPRes:error:%s", in)
		return "", fmt.Errorf(string(in))
	}

	t, ok := m["data"]
	if !ok {
		return "", fmt.Errorf("error result format no 'data' field")
	}

	if reflect.TypeOf(t).Kind() != reflect.Map {
		return "", fmt.Errorf("error result format 'data' field error")
	}
	t2, ok := t.(map[string]interface{})[method]
	if !ok {
		return "", fmt.Errorf("error result format no '%s' field", method)
	}

	//todo refactor me!!!
	if method == "queryAssets" {
		bts, err := json.Marshal(t2)
		if err != nil {
			log.Errorf("Marshal:error:%s", err.Error())
			return "", err
		}
		return string(bts), nil
	}

	res := ""
	if t2.(map[string]interface{})["encrypted"] != nil && len(t2.(map[string]interface{})["encrypted"].(string)) > 0 {
		res = t2.(map[string]interface{})["encrypted"].(string)
	} else {
		data := t2.(map[string]interface{})["data"]
		b, err := json.Marshal(data)
		if err != nil {
			return "", fmt.Errorf("json marshal failed")
		}
		res = string(b)
	}
	return res, nil
}
