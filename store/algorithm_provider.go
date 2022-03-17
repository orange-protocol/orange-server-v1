package store

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/orange-protocol/orange-server-v1/graph/model"
	"github.com/orange-protocol/orange-server-v1/utils"
)

func (this *DBCon) QueryAlgorithmProviderByDid(did string) (*AlgorithmProvider, error) {
	strsql := "select * from algorithm_provider_info where did = ?"
	r, err := this.Dbconnect.Query(strsql, did)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	if r.Next() {
		t := &AlgorithmProvider{}
		err = r.Scan(&t.Did, &t.APType, &t.Introduction, &t.CreateTime, &t.Name, &t.ApiKey, &t.Title, &t.Provider,
			&t.InvokeFrequency, &t.ApiState, &t.Author, &t.Popularity, &t.Delay, &t.Icon)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, nil
}

func (this *DBCon) QueryAlgorithmProviderTitle(title string) (*AlgorithmProvider, error) {
	strsql := "select did, title from algorithm_provider_info where title = ?"
	r, err := this.Dbconnect.Query(strsql, title)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &AlgorithmProvider{}
		err = r.Scan(&t.Did, &t.Title)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, nil
}

func (this *DBCon) QueryAlgorithmProviderMethodByDid(did string) ([]*APMethod, error) {
	strsql := "select * from algorithm_provider_method_info where did = ? and status != ?"
	r, err := this.Dbconnect.Query(strsql, did, METHOD_REMOVED)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	res := make([]*APMethod, 0)
	for r.Next() {
		t := &APMethod{
			Labels: &LabelsInfo{},
		}
		err = r.Scan(&t.Did, &t.Method, &t.ParamSchema, &t.ResultSchema, &t.URL, &t.Param, &t.Name, &t.Description, &t.Invoked, &t.Latency, &t.CreateTime, &t.HttpMethod, &t.Status, &t.Labels.BlockChain, &t.Labels.Category, &t.Labels.Scenario)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (this *DBCon) QueryAllAlgorithmProvidersCountByCondition(where string) (int64, error) {
	strsql := "select count(*) from algorithm_provider_info " + where
	r, err := this.Dbconnect.Query(strsql)
	if err != nil {
		return 0, err
	}
	defer r.Close()
	res := int64(0)
	if r.Next() {
		err = r.Scan(&res)
	}
	return res, err
}

func (this *DBCon) QueryAllAlgorithmProvidersByCondition(where string) ([]*AlgorithmProvider, error) {
	strsql := "select * from algorithm_provider_info " + where
	r, err := this.Dbconnect.Query(strsql)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	res := make([]*AlgorithmProvider, 0)
	for r.Next() {
		t := &AlgorithmProvider{}
		err = r.Scan(&t.Did, &t.APType, &t.Introduction, &t.CreateTime, &t.Name, &t.ApiKey, &t.Title, &t.Provider,
			&t.InvokeFrequency, &t.ApiState, &t.Author, &t.Popularity, &t.Delay, &t.Icon)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (this *DBCon) QueryAllAlgorithmProviderMethodsCountByCondition(where string) (int64, error) {
	strsql := "select count(*) from algorithm_provider_method_info t1" + where
	r, err := this.Dbconnect.Query(strsql)
	if err != nil {
		return 0, err
	}
	defer r.Close()
	res := int64(0)
	if r.Next() {
		err = r.Scan(&res)
	}
	return res, err
}

func (this *DBCon) QueryAllAlgorithmProviderMethodsByCondition(where string) ([]*APMethodWithAPInfo, error) {
	strsql := "select t1.*,t2.* from algorithm_provider_method_info t1 left join algorithm_provider_info t2 on t1.did = t2.did " + where
	r, err := this.Dbconnect.Query(strsql)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	res := make([]*APMethodWithAPInfo, 0)
	for r.Next() {
		ap := &AlgorithmProvider{}
		apm := &APMethod{
			Labels: &LabelsInfo{},
		}
		err = r.Scan(&apm.Did, &apm.Method, &apm.ParamSchema, &apm.ResultSchema, &apm.URL, &apm.Param, &apm.Name, &apm.Description, &apm.Invoked, &apm.Latency, &apm.CreateTime, &apm.HttpMethod, &apm.Status, &apm.Labels.BlockChain, &apm.Labels.Category, &apm.Labels.Scenario,
			&ap.Did, &ap.APType, &ap.Introduction, &ap.CreateTime, &ap.Name, &ap.ApiKey, &ap.Title, &ap.Provider, &ap.InvokeFrequency, &ap.ApiState, &ap.Author, &ap.Popularity, &ap.Delay, &ap.Icon)
		if err != nil {
			return nil, err
		}
		res = append(res, &APMethodWithAPInfo{
			Ap:       ap,
			ApMethod: apm,
		})
	}
	return res, nil
}

func (this *DBCon) QueryAllAlgorithmProviders() ([]*AlgorithmProvider, error) {
	strsql := "select * from algorithm_provider_info"
	r, err := this.Dbconnect.Query(strsql)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	res := make([]*AlgorithmProvider, 0)
	for r.Next() {
		t := &AlgorithmProvider{}
		err = r.Scan(&t.Did, &t.APType, &t.Introduction, &t.CreateTime, &t.Name, &t.ApiKey, &t.Title, &t.Provider,
			&t.InvokeFrequency, &t.ApiState, &t.Author, &t.Popularity, &t.Delay, &t.Icon)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (this *DBCon) QueryAPMethodByDIDAndMethod(did, method string) (*APMethod, error) {
	strsql := "select * from algorithm_provider_method_info where did = ? and method = ? and status != ?"
	r, err := this.Dbconnect.Query(strsql, did, method, METHOD_REMOVED)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &APMethod{
			Labels: &LabelsInfo{},
		}
		err = r.Scan(&t.Did, &t.Method, &t.ParamSchema, &t.ResultSchema, &t.URL, &t.Param, &t.Name, &t.Description, &t.Invoked, &t.Latency, &t.CreateTime, &t.HttpMethod, &t.Status, &t.Labels.BlockChain, &t.Labels.Category, &t.Labels.Scenario)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, nil

}

func (this *DBCon) AddAPInvokeFrenquency(did string, delay int64) error {
	strsql := "update algorithm_provider_info set delay = ?, invokeFrequency = invokeFrequency + 1 where did = ? "
	_, err := this.Dbconnect.Exec(strsql, delay, did)
	return err
}

func (this *DBCon) AddAPMethodInvoked(did, method string, delay int64) error {
	strsql := "update algorithm_provider_method_info set latency = ?, invoked = invoked + 1 where did = ? and method = ?"
	_, err := this.Dbconnect.Exec(strsql, delay, did, method)
	return err
}

func (this *DBCon) updateAPMethodStatus(did, method, status string) error {
	strsql := "update algorithm_provider_method_info set status=? where did = ? and method = ?"
	_, err := this.Dbconnect.Exec(strsql, status, did, method)
	return err
}

func (this *DBCon) GetAllTaskHistoryCount() (int64, error) {
	strsql := "select count(*) from task_history"
	r, err := this.Dbconnect.Query(strsql)
	if err != nil {
		return 0, err
	}
	defer r.Close()
	count := int64(0)
	if r.Next() {
		err = r.Scan(&count)
	}
	return count, err
}

func (this *DBCon) GetAPPopularity(did string) (int64, error) {

	totalcnt, err := this.GetAllTaskHistoryCount()
	if err != nil {
		return 0, err
	}
	if totalcnt == 0 {
		return 0, nil
	}

	strsql := "select count(*) from task_history where ap_did = ?"
	r, err := this.Dbconnect.Query(strsql, did)
	if err != nil {
		return 0, err
	}
	defer r.Close()
	count := int64(0)
	if r.Next() {
		err = r.Scan(&count)
	}

	return count * 100 / totalcnt, err
}
func (this *DBCon) GetDPPopularity(did string) (int64, error) {

	totalcnt, err := this.GetAllTaskHistoryCount()
	if err != nil {
		return 0, err
	}
	if totalcnt == 0 {
		return 0, nil
	}

	strsql := "select count(*) from task_history where dp_did = ?"
	r, err := this.Dbconnect.Query(strsql, did)
	if err != nil {
		return 0, err
	}
	defer r.Close()
	count := int64(0)
	if r.Next() {
		err = r.Scan(&count)
	}
	return count * 100 / totalcnt, err
}

func (this *DBCon) GetParamSchema(did, method string) (string, error) {
	strsql := "select param_schema from algorithm_provider_method_info where did=? and method=?"
	r, err := this.Dbconnect.Query(strsql, did, method)
	if err != nil {
		return "", err
	}
	defer r.Close()
	paramSchema := ""
	if r.Next() {
		err = r.Scan(&paramSchema)
		return paramSchema, err
	}
	return "", nil
}

func (this *DBCon) GetAllApInfo() ([]*model.MethodInfo, error) {
	strsql := "select * from algorithm_provider_method_info where status != ? order by did,create_time desc"
	r, err := this.Dbconnect.Query(strsql, METHOD_REMOVED)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	mapDpInfo := make(map[string][]*model.MethodDetail)
	keys := make([]string, 0)
	for r.Next() {
		t := &model.MethodDetail{
			Labels: &model.LabelsInfos{
				BlockChain: make([]string, 0),
				Category:   make([]string, 0),
				Scenario:   make([]string, 0),
			},
		}
		var blockChainLabels string
		var categoryLabels string
		var scenarioLabels string
		err = r.Scan(&t.Did, &t.Method, &t.ParamSchema, &t.ResultSchema, &t.URL, &t.Param, &t.Name, &t.Description, &t.Invoked, &t.Latency, &t.CreateTime, &t.HTTPMethod, &t.Status, &blockChainLabels, &categoryLabels, &scenarioLabels)
		if err != nil {
			return nil, err
		}
		res := make([]*model.MethodDetail, 0)
		if dpInfo, present := mapDpInfo[t.Did]; present {
			if blockChainLabels != "" {
				t.Labels.BlockChain = strings.Split(blockChainLabels, ",")
			}
			if categoryLabels != "" {
				t.Labels.Category = strings.Split(categoryLabels, ",")
			}
			if scenarioLabels != "" {
				t.Labels.Scenario = strings.Split(scenarioLabels, ",")
			}
			dpInfo = append(dpInfo, t)
			mapDpInfo[t.Did] = dpInfo
		} else {
			res = append(res, t)
			mapDpInfo[t.Did] = res
			keys = append(keys, t.Did)
		}
	}
	dpInfos := make([]*model.MethodInfo, 0)
	sort.Strings(keys)
	for _, k := range keys {
		dataProvider, err := this.QueryAlgorithmProviderByDid(k)
		if err != nil {
			return nil, err
		}
		dpMethodInfo := &model.MethodInfo{
			Did:   dataProvider.Did,
			Title: dataProvider.Title,
			Data:  mapDpInfo[k],
		}
		dpInfos = append(dpInfos, dpMethodInfo)
	}
	return dpInfos, nil
}

func (db *DBCon) SaveAPInfo(userDID string, apInfo *model.SubmitApInfo) (bool, error) {
	strsql := ""
	var err error
	if apInfo.ApInfoID == 0 {
		strsql = "insert into ap_info(user_did,ap_did,ap_name,ap_desc,avatar,status,create_time) values (?,?,?,?,?,?,?)"
		_, err = db.Dbconnect.Exec(strsql, userDID, apInfo.ApDid, apInfo.ApName, apInfo.ApDesc, apInfo.Avatar, utils.DRAFT, time.Now().Unix())
		if err != nil {
			return false, nil
		}
	} else {
		res, err := db.QueryAPInfo(userDID, utils.PUBLISHED)
		if err != nil {
			return false, err
		}
		if res == nil {
			strsql = "update ap_info set user_did=?,ap_did=?,ap_name=?,ap_desc=?,avatar=?,status=?,create_time=? where ap_info_id=?"
			_, err = db.Dbconnect.Exec(strsql, userDID, apInfo.ApDid, apInfo.ApName, apInfo.ApDesc, apInfo.Avatar, utils.VERIFYING, time.Now().Unix(), apInfo.ApInfoID)
			if err != nil {
				return false, err
			}
		} else {
			strsql = "insert into ap_info(user_did,ap_did,ap_name,ap_desc,avatar,status,create_time) values (?,?,?,?,?,?,?)"
			_, err = db.Dbconnect.Exec(strsql, userDID, apInfo.ApDid, apInfo.ApName, apInfo.ApDesc, apInfo.Avatar, utils.DRAFT, time.Now().Unix())
			if err != nil {
				return false, nil
			}
		}
	}
	return true, nil
}

func (db *DBCon) SubmitAPInfo(userDID string, apInfo *model.SubmitApInfo) (bool, error) {
	strsql := ""
	var err error
	if apInfo.ApInfoID == 0 {
		strsql = "insert into ap_info(user_did,ap_did,ap_name,ap_desc,avatar,status,create_time) values (?,?,?,?,?,?,?)"
		_, err = db.Dbconnect.Exec(strsql, userDID, apInfo.ApDid, apInfo.ApName, apInfo.ApDesc, apInfo.Avatar, utils.VERIFYING, time.Now().Unix())
		if err != nil {
			return false, nil
		}
	} else {
		res, err := db.QueryAPInfo(userDID, utils.PUBLISHED)
		if err != nil {
			return false, err
		}
		if res == nil {
			strsql = "update ap_info set user_did=?,ap_did=?,ap_name=?,ap_desc=?,avatar=?,status=?,create_time=? where ap_info_id=?"
			_, err = db.Dbconnect.Exec(strsql, userDID, apInfo.ApDid, apInfo.ApName, apInfo.ApDesc, apInfo.Avatar, utils.VERIFYING, time.Now().Unix(), apInfo.ApInfoID)
			if err != nil {
				return false, err
			}
		} else {
			draftResult, err := db.QueryAPInfo(userDID, utils.DRAFT)
			if err != nil {
				return false, err
			}
			failedResult, err := db.QueryAPInfo(userDID, utils.FAILED)
			if err != nil {
				return false, err
			}
			if draftResult != nil && draftResult.ApInfoID == apInfo.ApInfoID {
				strsql = "update ap_info set user_did=?,ap_did=?,ap_name=?,ap_desc=?,avatar=?,status=?,create_time=? where ap_info_id=?"
				_, err = db.Dbconnect.Exec(strsql, userDID, apInfo.ApDid, apInfo.ApName, apInfo.ApDesc, apInfo.Avatar, utils.VERIFYING, time.Now().Unix(), apInfo.ApInfoID)
				if err != nil {
					return false, err
				}
			} else if failedResult != nil && failedResult.ApInfoID == apInfo.ApInfoID {
				strsql = "update ap_info set user_did=?,ap_did=?,ap_name=?,ap_desc=?,avatar=?,status=?,create_time=? where ap_info_id=?"
				_, err = db.Dbconnect.Exec(strsql, userDID, apInfo.ApDid, apInfo.ApName, apInfo.ApDesc, apInfo.Avatar, utils.VERIFYING, time.Now().Unix(), apInfo.ApInfoID)
				if err != nil {
					return false, err
				}
			} else {
				strsql = "insert into ap_info(user_did,ap_did,ap_name,ap_desc,avatar,status,create_time) values (?,?,?,?,?,?,?)"
				_, err = db.Dbconnect.Exec(strsql, userDID, apInfo.ApDid, apInfo.ApName, apInfo.ApDesc, apInfo.Avatar, utils.VERIFYING, time.Now().Unix())
				if err != nil {
					return false, nil
				}
			}
		}
	}
	return true, nil
}

func (this *DBCon) UpdateApInfo(apInfo *model.SubmitApInfo) (bool, error) {
	strsql := "update ap_info set ap_name=?,ap_desc=?,avatar=? where ap_did=?"
	_, err := this.Dbconnect.Exec(strsql, apInfo.ApName, apInfo.ApDesc, apInfo.Avatar, apInfo.ApDid)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (this *DBCon) RevokeAPInfo(userDID, afterStatus, beforeStatus string, apInfoID int64) (bool, error) {
	strsql := "update ap_info set status=?,create_time=?  where ap_info_id=? and status=?"
	_, err := this.Dbconnect.Exec(strsql, afterStatus, time.Now().Unix(), apInfoID, beforeStatus)
	if err != nil {
		return false, err
	}
	return true, nil
	return true, nil
}

func (this *DBCon) QueryAPInfo(userDID, status string) (*model.ApInfoRes, error) {
	strsql := ""
	var err error
	var r *sql.Rows
	if status == "" {
		strsql = "select ap_info_id,ap_did,ap_name,ap_desc,avatar,status,res_desc from ap_info where user_did=? order by create_time desc limit 1"
		r, err = this.Dbconnect.Query(strsql, userDID)
	} else {
		strsql = "select ap_info_id,ap_did,ap_name,ap_desc,avatar,status,res_desc from ap_info where user_did=? and status=?"
		r, err = this.Dbconnect.Query(strsql, userDID, status)
	}
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &model.ApInfoRes{}
		err := r.Scan(&t.ApInfoID, &t.ApDid, &t.ApName, &t.ApDesc, &t.Avatar, &t.Status, &t.ResDesc)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, nil
}

func (db *DBCon) CheckApNameExist(apName string) (*model.ApInfoRes, string, error) {
	strsql := "select user_did, ap_info_id,ap_did from ap_info where ap_name =?"
	r, err := db.Dbconnect.Query(strsql, apName)
	if err != nil {
		return nil, "", err
	}
	defer r.Close()
	if r.Next() {
		t := &model.ApInfoRes{}
		userDId := ""
		err := r.Scan(&userDId, &t.ApInfoID, &t.ApDid)
		if err != nil {
			return nil, "", err
		}
		return t, userDId, nil
	}
	return nil, "", nil
}

func (this *DBCon) QueryUserDIDByAPDID(apDID, status string) (string, error) {
	strsql := "select user_did from ap_info where ap_did=? and status=?"
	r, err := this.Dbconnect.Query(strsql, apDID, status)
	if err != nil {
		return "", err
	}
	defer r.Close()
	userDID := ""
	if r.Next() {
		err := r.Scan(&userDID)
		if err != nil {
			return "", err
		}
	}
	return userDID, nil
}

func (this *DBCon) SubmitAPDataSet(apDataSet *model.ApDataSetInfo, userDid, status string) (bool, error) {
	apInfo, err := this.QueryAPInfo(userDid, utils.PUBLISHED)
	if err != nil {
		return false, err
	}
	apDID := ""
	if apInfo != nil {
		apDID = *apInfo.ApDid
		//return false, fmt.Errorf("ap info not published")
	}
	//apDID := apInfo.ApDid
	if status == utils.VERIFYING {
		if apDID == "" {
			return false, fmt.Errorf("AP Info not published")
		}
		if apDataSet.DataSetID != 0 {
			dataSetInfo, err := this.QueryAPDataSetInfo(userDid, apDataSet.DataSetID)
			if err != nil {
				return false, err
			}
			if dataSetInfo.Status == utils.PUBLISHED || dataSetInfo.Status == utils.VERIFYING || dataSetInfo.Status == utils.REVOKING {
				return false, fmt.Errorf("Dataset could not be published with current status")
			}
		}
		apInfo, err := this.QueryAPInfo(userDid, utils.PUBLISHED)
		if err != nil {
			return false, err
		}
		if apInfo == nil || (status == utils.VERIFYING && apInfo.Status != utils.PUBLISHED) {
			return false, fmt.Errorf("Please submit AP Info first")
		}
	}
	var strsql string
	strsql = "select count(user_did) from ap_data_set where user_did =? and status !=?"
	r, err := this.Dbconnect.Query(strsql, userDid, utils.PUBLISHED)
	if err != nil {
		return false, err
	}
	defer r.Close()
	res := int64(0)
	if r.Next() {
		err = r.Scan(&res)
	}
	if err != nil {
		return false, err
	}
	if apDataSet.DataSetID == 0 {
		if res >= 10 {
			return false, fmt.Errorf("Number of model has reached 10")
		}
	}
	blockChainLabels := ""
	categoryLabels := ""
	scenarioLabels := ""
	if apDataSet.Labels != nil {
		blockChainLabels = strings.Join(apDataSet.Labels.BlockChain, ",")
		categoryLabels = strings.Join(apDataSet.Labels.Category, ",")
		scenarioLabels = strings.Join(apDataSet.Labels.Scenario, ",")
	}

	if apDataSet.DataSetID == 0 {
		strsql = "insert into ap_data_set(user_did,ap_did,dataset_name,method_name,data_desc,http_method,http_url,params,status,create_time,block_chain_labels,category_labels,scenario_labels) values (?,?,?,?,?,?,?,?,?,?,?,?,?)"
		_, err = this.Dbconnect.Exec(strsql, userDid, apDID, apDataSet.DataSetName, apDataSet.DataSetMethodName, apDataSet.DataSetDesc,
			apDataSet.HTTPMethod, apDataSet.HTTPURL, apDataSet.Params, status, time.Now().Unix(), blockChainLabels, categoryLabels, scenarioLabels)
		if err != nil {
			return false, err
		}
	} else {
		strsql = "update ap_data_set set user_did=?, ap_did=?,dataset_name=?,method_name=?,data_desc=?,http_method=?,http_url=?,params=?,status=?,res_desc=?,create_time=?,block_chain_labels=?,category_labels=?,scenario_labels=? where data_set_id=?"
		_, err = this.Dbconnect.Exec(strsql, userDid, apDID, apDataSet.DataSetName, apDataSet.DataSetMethodName, apDataSet.DataSetDesc,
			apDataSet.HTTPMethod, apDataSet.HTTPURL, apDataSet.Params, status, "", time.Now().Unix(), blockChainLabels, categoryLabels, scenarioLabels, apDataSet.DataSetID)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (this *DBCon) DeleteAPDataSet(userDid string, dataSetID int64) (bool, error) {
	strSql := "delete from ap_data_set where user_did=? and data_set_id=?"
	_, err := this.Dbconnect.Exec(strSql, userDid, dataSetID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (this *DBCon) RevokeAPDataSet(userDid, afterStatus, beforeStatus string, dataSetID int64) (bool, error) {
	res, err := this.QueryAPDataSetInfo(userDid, dataSetID)
	if err != nil {
		return false, err
	}
	if res == nil || res.Status != utils.VERIFYING {
		return false, fmt.Errorf("RevokeAPDataSet failed, current apDataSet status not verifying")
	}
	return this.UpdateAPDataSetStatus(userDid, afterStatus, beforeStatus, dataSetID)
}

func (this *DBCon) UpdateAPDataSetStatus(userDid, afterStatus, beforeStatus string, dataSetID int64) (bool, error) {
	strsql := "update ap_data_set set status=? where user_did=? and  data_set_id=? and status=?"
	_, err := this.Dbconnect.Exec(strsql, afterStatus, userDid, dataSetID, beforeStatus)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (this *DBCon) QueryAPDataSetInfo(userDid string, dataSetID int64) (*model.ApDataSetRes, error) {
	var blockChainLabels string
	var categoryLabels string
	var scenarioLabels string
	strsql := "select data_set_id,ap_did, dataset_name,method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from ap_data_set where user_did=? and data_set_id=?"
	r, err := this.Dbconnect.Query(strsql, userDid, dataSetID)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &model.ApDataSetRes{
			Labels: &model.LabelsInfos{
				BlockChain: make([]string, 0),
				Category:   make([]string, 0),
				Scenario:   make([]string, 0),
			},
		}
		err := r.Scan(&t.DataSetID, &t.ApDid, &t.DataSetName, &t.DataSetMethodName, &t.DataSetDesc, &t.HTTPMethod, &t.HTTPURL, &t.Params, &t.Status, &t.ResDesc, &blockChainLabels, &categoryLabels, &scenarioLabels)
		if err != nil {
			return nil, err
		}
		if blockChainLabels != "" {
			t.Labels.BlockChain = strings.Split(blockChainLabels, ",")
		}
		if categoryLabels != "" {
			t.Labels.Category = strings.Split(categoryLabels, ",")
		}
		if scenarioLabels != "" {
			t.Labels.Scenario = strings.Split(scenarioLabels, ",")
		}
		return t, nil
	}
	return nil, nil
}

func (this *DBCon) QueryUserAPDataSetLists(userDID, dataSetName, status string, page int64, size int64, labels model.LabelsInfo) (*model.APDataSetList, error) {
	aPDataSetList := &model.APDataSetList{
		ApDataSetData: make([]*model.ApDataSetRes, 0),
		CurPageNum:    0,
		TotalNum:      0,
	}
	whereStr := ""
	for _, blockChainLabels := range labels.BlockChain {
		whereStr += "block_chain_labels like '%" + blockChainLabels + "%' or "
	}
	for _, catgegoryLabels := range labels.Category {
		whereStr += "category_labels like '%" + catgegoryLabels + "%' or "
	}
	for _, catgegoryLabels := range labels.Scenario {
		whereStr += "category_labels like '%" + catgegoryLabels + "%' or "
	}
	whereStr = strings.TrimRight(whereStr, "or ")
	var r *sql.Rows
	var err error
	if status == "" {
		if dataSetName == "" {
			strsql := ""
			if whereStr == "" {
				strsql = "select data_set_id,ap_did, dataset_name, method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from ap_data_set where user_did=? order by create_time desc limit ? OFFSET ?"
			} else {
				strsql = "select data_set_id,ap_did, dataset_name, method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from ap_data_set where user_did=? and " + whereStr + "order by create_time desc limit ? OFFSET ?"
			}
			r, err = this.Dbconnect.Query(strsql, userDID, size, (page-1)*size)
		} else {
			strsql := "select data_set_id,ap_did, dataset_name, method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from ap_data_set where user_did=? and dataset_name like '%" + dataSetName + "%'" + whereStr +
				"order by create_time desc limit ? OFFSET ?"
			r, err = this.Dbconnect.Query(strsql, userDID, size, (page-1)*size)
		}
	} else {
		if dataSetName == "" {
			strsql := ""
			if whereStr == "" {
				strsql = "select data_set_id,ap_did,dataset_name, method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from ap_data_set where user_did=? and status=? order by create_time desc limit ? OFFSET ?"
			} else {
				strsql = "select data_set_id,ap_did,dataset_name, method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from ap_data_set where user_did=? and status=? and " + whereStr + "order by create_time desc limit ? OFFSET ?"
			}
			r, err = this.Dbconnect.Query(strsql, userDID, status, size, (page-1)*size)
		} else {
			strsql := "select data_set_id,ap_did, dataset_name, method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from ap_data_set where user_did=? and status=? and dataset_name like '%" + dataSetName + "%'" + whereStr +
				"order by create_time desc limit ? OFFSET ?"
			r, err = this.Dbconnect.Query(strsql, userDID, status, size, (page-1)*size)

		}
	}
	if err != nil {
		return nil, err
	}
	defer r.Close()
	for r.Next() {
		t := &model.ApDataSetRes{}
		var blockChainLabels string
		var categoryLabels string
		var scenarioLabels string
		err := r.Scan(&t.DataSetID, &t.ApDid, &t.DataSetName, &t.DataSetMethodName, &t.DataSetDesc, &t.HTTPMethod, &t.HTTPURL, &t.Params, &t.Status, &t.ResDesc, &blockChainLabels, &categoryLabels, &scenarioLabels)
		if err != nil {
			return nil, err
		}
		res := &model.ApDataSetRes{
			DataSetID:         t.DataSetID,
			ApDid:             t.ApDid,
			DataSetName:       t.DataSetName,
			DataSetMethodName: t.DataSetMethodName,
			DataSetDesc:       t.DataSetDesc,
			HTTPMethod:        t.HTTPMethod,
			HTTPURL:           t.HTTPURL,
			Params:            t.Params,
			Status:            t.Status,
			ResDesc:           t.ResDesc,
			Labels: &model.LabelsInfos{
				BlockChain: make([]string, 0),
				Category:   make([]string, 0),
				Scenario:   make([]string, 0),
			},
		}
		if blockChainLabels != "" {
			res.Labels.BlockChain = strings.Split(blockChainLabels, ",")
		}
		if categoryLabels != "" {
			res.Labels.Category = strings.Split(categoryLabels, ",")
		}
		if scenarioLabels != "" {
			res.Labels.Scenario = strings.Split(scenarioLabels, ",")
		}
		aPDataSetList.ApDataSetData = append(aPDataSetList.ApDataSetData, res)
	}

	var r1 *sql.Rows
	var err1 error
	if status == "" {
		if dataSetName == "" {
			strsql := "select count(data_set_id)  from ap_data_set where user_did=?"
			r1, err1 = this.Dbconnect.Query(strsql, userDID)
		} else {
			strsql := "select count(data_set_id) from ap_data_set where user_did=? and dataset_name like '%" + dataSetName + "%'"
			r1, err1 = this.Dbconnect.Query(strsql, userDID)
		}
	} else {
		if dataSetName == "" {
			strsql := "select count(data_set_id)  from ap_data_set where user_did=? and status=?"
			r1, err1 = this.Dbconnect.Query(strsql, userDID, status)
		} else {
			strsql := "select count(data_set_id) from ap_data_set where user_did=? and status=? and dataset_name like '%" + dataSetName + "%'"
			r1, err1 = this.Dbconnect.Query(strsql, userDID, status)
		}
	}
	if err1 != nil {
		return nil, err
	}
	defer r1.Close()
	totalNum := int64(0)
	if r1.Next() {
		err = r1.Scan(&totalNum)
		if err != nil {
			return nil, err
		}
	}
	aPDataSetList.CurPageNum = page
	aPDataSetList.TotalNum = totalNum
	return aPDataSetList, nil
}

func (db *DBCon) QueryUserAP(userDID string) (*model.ApInfoRes, error) {
	strsql := "select *from ap_draft_info where did=?"
	r, err := db.Dbconnect.Query(strsql, userDID)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &model.ApInfoRes{}
		err := r.Scan(&t.ApDid, &t.ApName, &t.ApDesc, &t.Status)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, nil
}

func (this *DBCon) CancelAPInfo(apDID, status string) error {
	strSql := "delete from ap_info where ap_did=? and status=?"
	_, err := this.Dbconnect.Exec(strSql, apDID, status)
	return err
}

func (this *DBCon) UpdateAPDataSetSchema(apDid, dataSetName, inputSchema, outputSchema string) (bool, error) {
	strsql := "update ap_data_set set input_schema=?,output_schema=? " +
		"where ap_did=? and method_name=?"
	_, err := this.Dbconnect.Exec(strsql, inputSchema, outputSchema, apDid, dataSetName)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (this *DBCon) RevokePublishedAPDataSet(userDid string, dataSetID int64) (bool, error) {
	res, err := this.QueryAPDataSetInfo(userDid, dataSetID)
	if err != nil {
		return false, err
	}
	if res == nil {
		return false, fmt.Errorf("PublishedAPDataSet dataSetID:%d invalid", dataSetID)
	}
	if res.Status != utils.PUBLISHED {
		return false, fmt.Errorf("PublishedAPDataSet status:%s invalid", res.Status)
	}
	err = this.updateAPMethodStatus(res.ApDid, res.DataSetMethodName, utils.REVOKING)
	if err != nil {
		return false, err
	}
	return this.UpdateAPDataSetStatus(userDid, utils.REVOKING, utils.PUBLISHED, dataSetID)
}

func (this *DBCon) CheckDuplicateAPDataSetName(apDataSet *model.ApDataSetInfo, userDid string) (bool, error) {
	apDataSetList, err := this.QueryUserAPDataSetList(userDid, "")
	if err != nil {
		return false, err
	}
	if len(apDataSetList) == 0 {
		return true, nil
	}
	flag, err := this.CheckDuplicateAP(apDataSetList, apDataSet)
	if err != nil {
		return false, err
	}
	return flag, nil
}

func (this *DBCon) CheckDuplicateAP(apDataSetList []*model.ApDataSetRes, apDataSet *model.ApDataSetInfo) (bool, error) {
	dataSetName := make(map[string]bool, 0)
	methodName := make(map[string]bool, 0)
	for _, dataSet := range apDataSetList {
		if dataSet.DataSetID != apDataSet.DataSetID {
			dataSetName[dataSet.DataSetName] = true
			methodName[dataSet.DataSetMethodName] = true
		}
	}
	if dataSetName[apDataSet.DataSetName] || methodName[apDataSet.DataSetMethodName] {
		return false, fmt.Errorf("dataSetName or methodName duplicate")
	}
	return true, nil
}

func (this *DBCon) QueryUserAPDataSetList(userDID, status string) ([]*model.ApDataSetRes, error) {
	var err error
	var r *sql.Rows
	if status == "" {
		strsql := "select data_set_id,ap_did, dataset_name, method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from ap_data_set where user_did=?"
		r, err = this.Dbconnect.Query(strsql, userDID)
	} else {
		strsql := "select data_set_id,ap_did, dataset_name, method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from ap_data_set where user_did=? and status=?"
		r, err = this.Dbconnect.Query(strsql, userDID, status)
	}
	if err != nil {
		return nil, err
	}
	apDataSetReses := make([]*model.ApDataSetRes, 0)
	defer r.Close()
	for r.Next() {
		t := &model.ApDataSetRes{}
		var blockChainLabels string
		var categoryLabels string
		var scenarioLabels string
		err := r.Scan(&t.DataSetID, &t.ApDid, &t.DataSetName, &t.DataSetMethodName, &t.DataSetDesc, &t.HTTPMethod, &t.HTTPURL, &t.Params, &t.Status, &t.ResDesc, &blockChainLabels, &categoryLabels, &scenarioLabels)
		if err != nil {
			return nil, err
		}
		apDataSetRes := &model.ApDataSetRes{
			DataSetID:         t.DataSetID,
			ApDid:             t.ApDid,
			DataSetName:       t.DataSetName,
			DataSetMethodName: t.DataSetMethodName,
			DataSetDesc:       t.DataSetDesc,
			HTTPMethod:        t.HTTPMethod,
			HTTPURL:           t.HTTPURL,
			Params:            t.Params,
			Status:            t.Status,
			ResDesc:           t.ResDesc,
			Labels: &model.LabelsInfos{
				BlockChain: make([]string, 0),
				Category:   make([]string, 0),
				Scenario:   make([]string, 0),
			},
		}
		if blockChainLabels != "" {
			apDataSetRes.Labels.BlockChain = strings.Split(blockChainLabels, ",")
		}
		if categoryLabels != "" {
			apDataSetRes.Labels.Category = strings.Split(categoryLabels, ",")
		}
		if scenarioLabels != "" {
			apDataSetRes.Labels.Scenario = strings.Split(scenarioLabels, ",")
		}
		apDataSetReses = append(apDataSetReses, apDataSetRes)
	}
	return apDataSetReses, nil
}

func (db *DBCon) CheckAPDataSetNameExist(userDId, dataSetName string) (*model.ApDataSetRes, error) {
	strsql := "select data_set_id,dataset_name from ap_data_set where user_did=? and dataset_name=?"
	r, err := db.Dbconnect.Query(strsql, userDId, dataSetName)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &model.ApDataSetRes{}
		err := r.Scan(&t.DataSetID, &t.DataSetName)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, nil
}

func (db *DBCon) CheckAPDataSetMethodNameExist(userDId, methodName string) (*model.ApDataSetRes, error) {
	strsql := "select data_set_id,dataset_name from ap_data_set where user_did=? and method_name=?"
	r, err := db.Dbconnect.Query(strsql, userDId, methodName)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &model.ApDataSetRes{}
		err := r.Scan(&t.DataSetID, &t.DataSetName)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, nil
}

func (db *DBCon) CheckAPDataSetExistByAPDID(apDid string) (bool, error) {
	strsql := "select dataset_name from ap_data_set where ap_did=? and status=?"
	r, err := db.Dbconnect.Query(strsql, apDid, utils.PUBLISHED)
	if err != nil {
		return false, err
	}
	defer r.Close()
	name := ""
	if r.Next() {
		err := r.Scan(&name)
		if err != nil {
			return false, err
		}
	}
	return name != "", nil
}

func (this *DBCon) QueryAPDatatSetCountByCondition(userDId string) (int64, error) {
	strsql := "select count(user_did) from ap_data_set where user_did=? and status=?"
	r, err := this.Dbconnect.Query(strsql, userDId, utils.PUBLISHED)
	if err != nil {
		return 0, err
	}
	defer r.Close()
	res := int64(0)
	if r.Next() {
		err = r.Scan(&res)
	}
	return res, err
}
