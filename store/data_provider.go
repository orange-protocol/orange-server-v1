package store

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/orange-protocol/orange-server-v1/utils"

	"github.com/orange-protocol/orange-server-v1/graph/model"
)

func (this *DBCon) QueryDataProviderByDid(did string) (*DataProvider, error) {
	strsql := "select * from data_provider_info where did = ?"
	r, err := this.Dbconnect.Query(strsql, did)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	if r.Next() {
		t := &DataProvider{}
		err = r.Scan(&t.Did, &t.Introduction, &t.CreateTime, &t.DpType, &t.Name, &t.Apikey, &t.Title, &t.Provider,
			&t.InvokeFrequency, &t.ApiState, &t.Author, &t.Popularity, &t.Delay, &t.Icon)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, nil
}

func (this *DBCon) QueryDataProviderTitle(title string) (*DataProvider, error) {
	strsql := "select did,title from data_provider_info where title = ?"
	r, err := this.Dbconnect.Query(strsql, title)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &DataProvider{}
		err = r.Scan(&t.Did, &t.Title)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, nil
}

func (this *DBCon) QueryDataProviderMethodByDid(did string) ([]*DPMethod, error) {

	strsql := "select * from data_provider_method_info where did = ? and status != ? order by method desc"
	r, err := this.Dbconnect.Query(strsql, did, METHOD_REMOVED)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	res := make([]*DPMethod, 0)
	for r.Next() {
		t := &DPMethod{
			Labels: &LabelsInfo{},
		}
		var param, result, compositesetting sql.NullString
		err = r.Scan(&t.Did, &t.Method, &param, &result, &t.URL, &compositesetting, &t.Param, &t.Name, &t.Description, &t.Invoked, &t.Latency, &t.CreateTime, &t.HttpMethod, &t.Status, &t.Labels.BlockChain, &t.Labels.Category, &t.Labels.Scenario)
		if err != nil {
			return nil, err
		}
		if param.Valid {
			t.ParamSchema = param.String
		}
		if result.Valid {
			t.ResultSchema = result.String
		}
		if compositesetting.Valid {
			t.CompositeSetting = compositesetting.String
		}
		res = append(res, t)
	}
	return res, nil

}

func (this *DBCon) UpdateDPMethodLatency(did string, latency int64) error {
	strsql := "update data_provider_info set delay = ?,invokeFrequency = invokeFrequency+1 where did = ? "
	_, err := this.Dbconnect.Exec(strsql, latency, did)
	return err
}
func (this *DBCon) UpdateDPMethodInvoked(did, method string, latency int64) error {
	strsql := "update data_provider_method_info set latency = ?,invoked = invoked+1 where did = ? and method = ? "
	_, err := this.Dbconnect.Exec(strsql, latency, did, method)
	return err
}

func (this *DBCon) UpdateDPMethodStatus(did, method, status string) error {
	strsql := "update data_provider_method_info set status = ? where did = ? and method = ? "
	_, err := this.Dbconnect.Exec(strsql, status, did, method)
	return err
}

func (this *DBCon) QueryDPMethodByDIDAndMethod(did, method string) (*DPMethod, error) {
	strsql := "select * from data_provider_method_info where did = ? and method = ?"
	r, err := this.Dbconnect.Query(strsql, did, method)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &DPMethod{
			Labels: &LabelsInfo{},
		}
		var compositesetting sql.NullString
		err = r.Scan(&t.Did, &t.Method, &t.ParamSchema, &t.ResultSchema, &t.URL, &compositesetting, &t.Param, &t.Name, &t.Description, &t.Invoked, &t.Latency, &t.CreateTime, &t.HttpMethod, &t.Status, &t.Labels.BlockChain, &t.Labels.Category, &t.Labels.Scenario)
		if err != nil {
			return nil, err
		}
		if compositesetting.Valid {
			t.CompositeSetting = compositesetting.String
		}
		return t, nil
	}
	return nil, nil

}

func (this *DBCon) QueryAllDataProviders() ([]*DataProvider, error) {
	strsql := "select * from data_provider_info"
	r, err := this.Dbconnect.Query(strsql)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	res := make([]*DataProvider, 0)
	for r.Next() {
		t := &DataProvider{}
		err = r.Scan(&t.Did, &t.Introduction, &t.CreateTime, &t.DpType, &t.Name, &t.Apikey, &t.Title, &t.Provider,
			&t.InvokeFrequency, &t.ApiState, &t.Author, &t.Popularity, &t.Delay, &t.Icon)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (this *DBCon) QueryAllDataProvidersCountByCondition(where string) (int64, error) {
	strsql := "select count(*) from data_provider_info " + where
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

func (this *DBCon) QueryAllDataProvidersByCondition(where string) ([]*DataProvider, error) {
	strsql := "select * from data_provider_info " + where
	r, err := this.Dbconnect.Query(strsql)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	res := make([]*DataProvider, 0)
	for r.Next() {
		t := &DataProvider{}
		err = r.Scan(&t.Did, &t.Introduction, &t.CreateTime, &t.DpType, &t.Name, &t.Apikey, &t.Title, &t.Provider,
			&t.InvokeFrequency, &t.ApiState, &t.Author, &t.Popularity, &t.Delay, &t.Icon)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (this *DBCon) QueryAllDPMethodCountByCondition(where string) (int64, error) {
	strsql := "select count(*) from data_provider_method_info t1 " + where
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

func (this *DBCon) QueryAllDPMethodByCondition(where string) ([]*DPMethodWithDPInfo, error) {
	strsql := "select t1.*,t2.* from data_provider_method_info t1 left join data_provider_info t2 on t1.did = t2.did " + where
	r, err := this.Dbconnect.Query(strsql)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	res := make([]*DPMethodWithDPInfo, 0)
	for r.Next() {
		dp := &DataProvider{}
		dm := &DPMethod{
			Labels: &LabelsInfo{},
		}
		err = r.Scan(&dm.Did, &dm.Method, &dm.ParamSchema, &dm.ResultSchema, &dm.URL, &dm.CompositeSetting, &dm.Param, &dm.Name, &dm.Description, &dm.Invoked, &dm.Latency, &dm.CreateTime, &dm.HttpMethod, &dm.Status, &dm.Labels.BlockChain, &dm.Labels.Category, &dm.Labels.Scenario,
			&dp.Did, &dp.Introduction, &dp.CreateTime, &dp.DpType, &dp.Name, &dp.Apikey, &dp.Title, &dp.Provider, &dp.InvokeFrequency, &dp.ApiState, &dp.Author, &dp.Popularity, &dp.Delay, &dp.Icon)
		if err != nil {
			return nil, err
		}
		res = append(res, &DPMethodWithDPInfo{
			Dp:       dp,
			DpMethod: dm,
		})
	}
	return res, nil
}

func (this *DBCon) QueryDPAndMethodsByAP(apdid, apMethod string) ([]*DataProvider, []*DPMethod, error) {
	am, err := this.QueryAPMethodByDIDAndMethod(apdid, apMethod)
	if err != nil {
		return nil, nil, err
	}
	if am == nil {
		return nil, nil, fmt.Errorf("No AP found with AP DID:%s,apMethod:%s", apdid, apMethod)
	}

	strsql := "select t1.*,t2.* from data_provider_info t1 left join data_provider_method_info t2 on t1.did = t2.did where UPPER(t2.result_schema) = ?"
	r, err := this.Dbconnect.Query(strsql, am.ParamSchema)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	dps := make([]*DataProvider, 0)
	dms := make([]*DPMethod, 0)

	for r.Next() {
		tp := &DataProvider{}
		tm := &DPMethod{
			Labels: &LabelsInfo{},
		}
		err := r.Scan(&tp.Did, &tp.Introduction, &tp.CreateTime, &tp.DpType, &tp.Name, &tp.Apikey, &tp.Title, &tp.Provider, &tp.InvokeFrequency, &tp.ApiState, &tp.Author, &tp.Popularity, &tp.Delay, &tp.Icon,
			&tm.Did, &tm.Method, &tm.ParamSchema, &tm.ResultSchema, &tm.URL, &tm.CompositeSetting, &tm.Param, &tm.Name, &tm.Description, &tm.Invoked, &tm.Latency, &tm.CreateTime, &tm.HttpMethod, &tm.Status, &tm.Labels.BlockChain, &tm.Labels.Category, &tm.Labels.Scenario)
		if err != nil {
			return nil, nil, err
		}

		dps = append(dps, tp)
		dms = append(dms, tm)
	}
	if len(dps) != len(dms) {
		return nil, nil, fmt.Errorf("QueryDPAndMethodsByAP count not match")
	}
	return dps, dms, nil
}

func (this *DBCon) GetCompositeDpInfo(did, method string) ([]*model.MethodInfo, error) {
	param_schema, err := this.GetParamSchema(did, method)
	if err != nil {
		return nil, err
	}
	strsql := "select *from data_provider_method_info  where result_schema=? and status != ? order by create_time desc"
	r, err := this.Dbconnect.Query(strsql, param_schema, METHOD_REMOVED)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	mapDpInfo := make(map[string][]*model.MethodDetail)
	keys := make([]string, 0)
	for r.Next() {
		t := &model.MethodDetail{
			CompositeData: make([]*model.MethodDetail, 0),
			Labels: &model.LabelsInfos{
				BlockChain: make([]string, 0),
				Category:   make([]string, 0),
				Scenario:   make([]string, 0),
			},
		}
		var blockChainLabels string
		var categoryLabels string
		var scenarioLabels string
		err = r.Scan(&t.Did, &t.Method, &t.ParamSchema, &t.ResultSchema, &t.URL, &t.CompositeSetting, &t.Param,
			&t.Name, &t.Description, &t.Invoked, &t.Latency, &t.CreateTime, &t.HTTPMethod, &t.Status, &blockChainLabels, &categoryLabels, &scenarioLabels)
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
		res := make([]*model.MethodDetail, 0)
		if t.CompositeSetting == "NONE" {
			if dpInfo, present := mapDpInfo[t.Did]; present {
				dpInfo = append(dpInfo, t)
				mapDpInfo[t.Did] = dpInfo
			} else {
				res = append(res, t)
				mapDpInfo[t.Did] = res
				keys = append(keys, t.Did)
			}
		} else {
			dps := strings.Split(t.CompositeSetting, ";")
			re := make([]*model.MethodDetail, 0)
			for _, dp := range dps {
				results := strings.Split(dp, "#")
				if len(results) != 2 {
					return nil, fmt.Errorf("GetCompositeDpInfo split error:%d", len(results))
				} else {
					dps, err := this.GetDPMethodInfo(results[0], results[1])
					if err != nil {
						return nil, err
					}
					re = append(re, dps...)
				}
			}
			t.CompositeData = append(t.CompositeData, re...)
			if dpInfo, present := mapDpInfo[t.Did]; present {
				dpInfo = append(dpInfo, t)
				mapDpInfo[t.Did] = dpInfo
			} else {
				res = append(res, t)
				mapDpInfo[t.Did] = res
				keys = append(keys, t.Did)
			}
		}
	}
	dpInfos := make([]*model.MethodInfo, 0)
	sort.Strings(keys)
	for _, k := range keys {
		dataProvider, err := this.QueryDataProviderByDid(k)
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

func (this *DBCon) GetDPMethodInfo(did, method string) ([]*model.MethodDetail, error) {
	strsql := "select * from data_provider_method_info where did=? and method=?"
	r, err := this.Dbconnect.Query(strsql, did, method)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	res := make([]*model.MethodDetail, 0)
	for r.Next() {
		t := &model.MethodDetail{
			CompositeData: make([]*model.MethodDetail, 0),
			Labels: &model.LabelsInfos{
				BlockChain: make([]string, 0),
				Category:   make([]string, 0),
				Scenario:   make([]string, 0),
			},
		}
		var blockChainLabels string
		var categoryLabels string
		var scenarioLabels string
		err = r.Scan(&t.Did, &t.Method, &t.ParamSchema, &t.ResultSchema, &t.URL, &t.CompositeSetting, &t.Param,
			&t.Name, &t.Description, &t.Invoked, &t.Latency, &t.CreateTime, &t.HTTPMethod, &t.Status, &t.Labels.BlockChain, &t.Labels.Category, &t.Labels.Scenario)
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
		res = append(res, t)
	}
	return res, nil
}

func (db *DBCon) SaveDPInfo(userDID string, dpInfo *model.SubmitDpInfo) (bool, error) {
	strsql := ""
	var err error
	if dpInfo.DpInfoID == 0 {
		strsql = "insert into dp_info(user_did,dp_did,dp_name,dp_desc,avatar,status,create_time) values (?,?,?,?,?,?,?)"
		_, err = db.Dbconnect.Exec(strsql, userDID, dpInfo.DpDid, dpInfo.DpName, dpInfo.DpDesc, dpInfo.Avatar, utils.DRAFT, time.Now().Unix())
		if err != nil {
			return false, nil
		}
	} else {
		res, err := db.QueryDPInfo(userDID, utils.PUBLISHED)
		if err != nil {
			return false, err
		}
		if res == nil {
			strsql = "update dp_info set user_did=?,dp_did=?,dp_name=?,dp_desc=?,avatar=?,status=?,create_time=? where dp_info_id=?"
			_, err = db.Dbconnect.Exec(strsql, userDID, dpInfo.DpDid, dpInfo.DpName, dpInfo.DpDesc, dpInfo.Avatar, utils.VERIFYING, time.Now().Unix(), dpInfo.DpInfoID)
			if err != nil {
				return false, err
			}
		} else {
			strsql = "insert into dp_info(user_did,dp_did,dp_name,dp_desc,avatar,status,create_time) values (?,?,?,?,?,?,?)"
			_, err = db.Dbconnect.Exec(strsql, userDID, dpInfo.DpDid, dpInfo.DpName, dpInfo.DpDesc, dpInfo.Avatar, utils.DRAFT, time.Now().Unix())
			if err != nil {
				return false, nil
			}
		}
	}
	return true, nil
}

func (db *DBCon) QueryUserDP(userDID string) (*model.DpInfoRes, error) {
	strsql := "select dp_did,dp_name,dp_desc,avatar,status from dp_info where did=?"
	r, err := db.Dbconnect.Query(strsql, userDID)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &model.DpInfoRes{}
		err := r.Scan(&t.DpDid, &t.DpName, &t.DpDesc, &t.Avatar, &t.Status)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, nil
}

func (db *DBCon) CheckDpNameExist(dpName string) (*model.DpInfoRes, string, error) {
	strsql := "select user_did,dp_info_id,dp_did from dp_info where dp_name =?"
	r, err := db.Dbconnect.Query(strsql, dpName)
	if err != nil {
		return nil, "", err
	}
	defer r.Close()
	if r.Next() {
		t := &model.DpInfoRes{}
		userDId := ""
		err := r.Scan(&userDId, &t.DpInfoID, &t.DpDid)
		if err != nil {
			return nil, "", err
		}
		return t, userDId, nil
	}
	return nil, "", nil
}

func (db *DBCon) SubmitDPInfo(userDID string, dpInfo *model.SubmitDpInfo) (bool, error) {
	strsql := ""
	var err error
	if dpInfo.DpInfoID == 0 {
		strsql = "insert into dp_info(user_did,dp_did,dp_name,dp_desc,avatar,status,create_time) values (?,?,?,?,?,?,?)"
		_, err = db.Dbconnect.Exec(strsql, userDID, dpInfo.DpDid, dpInfo.DpName, dpInfo.DpDesc, dpInfo.Avatar, utils.VERIFYING, time.Now().Unix())
		if err != nil {
			return false, nil
		}
	} else {
		res, err := db.QueryDPInfo(userDID, utils.PUBLISHED)
		if err != nil {
			return false, err
		}
		if res == nil {
			strsql = "update dp_info set user_did=?,dp_did=?,dp_name=?,dp_desc=?,avatar=?,status=?,create_time=? where dp_info_id=?"
			_, err = db.Dbconnect.Exec(strsql, userDID, dpInfo.DpDid, dpInfo.DpName, dpInfo.DpDesc, dpInfo.Avatar, utils.VERIFYING, time.Now().Unix(), dpInfo.DpInfoID)
			if err != nil {
				return false, err
			}
		} else {
			draftResult, err := db.QueryDPInfo(userDID, utils.DRAFT)
			if err != nil {
				return false, err
			}
			failedResult, err := db.QueryDPInfo(userDID, utils.FAILED)
			if err != nil {
				return false, err
			}
			if draftResult != nil && draftResult.DpInfoID == dpInfo.DpInfoID {
				strsql = "update dp_info set user_did=?,dp_did=?,dp_name=?,dp_desc=?,avatar=?,status=?,create_time=? where dp_info_id=?"
				_, err = db.Dbconnect.Exec(strsql, userDID, dpInfo.DpDid, dpInfo.DpName, dpInfo.DpDesc, dpInfo.Avatar, utils.VERIFYING, time.Now().Unix(), dpInfo.DpInfoID)
				if err != nil {
					return false, err
				}
			} else if failedResult != nil && failedResult.DpInfoID == dpInfo.DpInfoID {
				strsql = "update dp_info set user_did=?,dp_did=?,dp_name=?,dp_desc=?,avatar=?,status=?,create_time=? where dp_info_id=?"
				_, err = db.Dbconnect.Exec(strsql, userDID, dpInfo.DpDid, dpInfo.DpName, dpInfo.DpDesc, dpInfo.Avatar, utils.VERIFYING, time.Now().Unix(), dpInfo.DpInfoID)
				if err != nil {
					return false, err
				}
			} else {
				strsql = "insert into dp_info(user_did,dp_did,dp_name,dp_desc,avatar,status,create_time) values (?,?,?,?,?,?,?)"
				_, err = db.Dbconnect.Exec(strsql, userDID, dpInfo.DpDid, dpInfo.DpName, dpInfo.DpDesc, dpInfo.Avatar, utils.VERIFYING, time.Now().Unix())
				if err != nil {
					return false, nil
				}
			}
		}
	}
	return true, nil
}
func (this *DBCon) RevokeDPInfo(userDID, afterStatus, beforeStatus string, dpInfoID int64) (bool, error) {
	strsql := "update dp_info set status=?,create_time=?  where dp_info_id=? and status=?"
	_, err := this.Dbconnect.Exec(strsql, afterStatus, time.Now().Unix(), dpInfoID, beforeStatus)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (this *DBCon) CancelDPInfo(dpDID, status string) error {
	strSql := "delete from dp_info where dp_did=? and status=?"
	_, err := this.Dbconnect.Exec(strSql, dpDID, status)
	return err
}

func (this *DBCon) QueryDPInfo(userDID, status string) (*model.DpInfoRes, error) {
	strsql := ""
	var err error
	var r *sql.Rows
	if status == "" {
		strsql = "select dp_info_id,dp_did,dp_name,dp_desc,avatar,status,res_desc from dp_info where user_did=? order by create_time desc limit 1"
		r, err = this.Dbconnect.Query(strsql, userDID)
	} else {
		strsql = "select dp_info_id, dp_did,dp_name,dp_desc,avatar,status,res_desc from dp_info where user_did=? and status=?"
		r, err = this.Dbconnect.Query(strsql, userDID, status)
	}
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &model.DpInfoRes{}
		err := r.Scan(&t.DpInfoID, &t.DpDid, &t.DpName, &t.DpDesc, &t.Avatar, &t.Status, &t.ResDesc)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, nil
}

func (this *DBCon) QueryUserDIDByDPDID(dpDID, status string) (string, error) {
	strsql := "select user_did from dp_info where dp_did=? and status=?"
	r, err := this.Dbconnect.Query(strsql, dpDID, status)
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

func (this *DBCon) CheckDuplicateDPDataSetName(dpDataSet *model.DpDataSetInfo, userDid string) (bool, error) {
	dpDataSetList, err := this.QueryUserDPDataSetList(userDid, "")
	if err != nil {
		return false, err
	}
	if len(dpDataSetList) == 0 {
		return true, nil
	}
	flag, err := this.CheckDuplicateDP(dpDataSetList, dpDataSet)
	if err != nil {
		return false, err
	}
	return flag, nil
}

func (this *DBCon) CheckDuplicateDP(dpDataSetList []*model.DpDataSetRes, dpDataSet *model.DpDataSetInfo) (bool, error) {
	dataSetName := make(map[string]bool, 0)
	methodName := make(map[string]bool, 0)
	for _, dataSet := range dpDataSetList {
		if dataSet.DataSetID != dpDataSet.DataSetID {
			dataSetName[dataSet.DataSetName] = true
			methodName[dataSet.DataSetMethodName] = true
		}
	}
	if dataSetName[dpDataSet.DataSetName] || methodName[dpDataSet.DataSetMethodName] {
		return false, fmt.Errorf("dataSetName or methodName duplicate")
	}
	return true, nil
}

func (this *DBCon) SubmitDPDataSet(dpDataSet *model.DpDataSetInfo, userDid, status string) (bool, error) {
	dpInfo, err := this.QueryDPInfo(userDid, utils.PUBLISHED)
	if err != nil {
		return false, err
	}
	dpDID := ""
	if dpInfo != nil {
		dpDID = *dpInfo.DpDid
		//return false, fmt.Errorf("dp info not published")
	}
	//dpDID := dpInfo.DpDid
	if status == utils.VERIFYING {
		if dpDID == "" {
			return false, fmt.Errorf("dp info not published")
		}
		if dpDataSet.DataSetID != 0 {
			dataSetInfo, err := this.QueryDPDataSetInfo(userDid, dpDataSet.DataSetID)
			if err != nil {
				return false, err
			}
			if dataSetInfo.Status == utils.PUBLISHED || dataSetInfo.Status == utils.VERIFYING || dataSetInfo.Status == utils.REVOKING {
				return false, fmt.Errorf("current data set status can not publish")
			}
		}
		dpInfo, err := this.QueryDPInfo(userDid, utils.PUBLISHED)
		if err != nil {
			return false, err
		}
		if dpInfo == nil || (status == utils.VERIFYING && dpInfo.Status != utils.PUBLISHED) {
			return false, fmt.Errorf("no dp released ")
		}
	}
	var strsql string
	strsql = "select count(user_did) from dp_data_set where user_did =? and status !=?"
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
	if dpDataSet.DataSetID == 0 {
		if res >= 10 {
			return false, fmt.Errorf("Number of dataset has reached 10")
		}
	}
	blockChainLabels := ""
	categoryLabels := ""
	scenarioLabels := ""
	if dpDataSet.Labels != nil {
		blockChainLabels = strings.Join(dpDataSet.Labels.BlockChain, ",")
		categoryLabels = strings.Join(dpDataSet.Labels.Category, ",")
		scenarioLabels = strings.Join(dpDataSet.Labels.Scenario, ",")
	}

	if dpDataSet.DataSetID == 0 {
		strsql = "insert into dp_data_set(user_did,dp_did,dataset_name,method_name,data_desc,http_method,http_url,params,status,create_time,block_chain_labels,category_labels,scenario_labels) values (?,?,?,?,?,?,?,?,?,?,?,?,?)"
		_, err = this.Dbconnect.Exec(strsql, userDid, dpDID, dpDataSet.DataSetName, dpDataSet.DataSetMethodName, dpDataSet.DataSetDesc,
			dpDataSet.HTTPMethod, dpDataSet.HTTPURL, dpDataSet.Params, status, time.Now().Unix(), blockChainLabels, categoryLabels, scenarioLabels)
		if err != nil {
			return false, err
		}
	} else {
		strsql = "update dp_data_set set user_did=?,dp_did=?,dataset_name=?,method_name=?,data_desc=?,http_method=?,http_url=?,params=?,status=?,res_desc=?,create_time=?,block_chain_labels=?,category_labels=?,scenario_labels=? where data_set_id=?"
		_, err = this.Dbconnect.Exec(strsql, userDid, dpDID, dpDataSet.DataSetName, dpDataSet.DataSetMethodName, dpDataSet.DataSetDesc,
			dpDataSet.HTTPMethod, dpDataSet.HTTPURL, dpDataSet.Params, status, "", time.Now().Unix(), blockChainLabels, categoryLabels, scenarioLabels, dpDataSet.DataSetID)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (this *DBCon) DeleteDPDataSet(userDid string, dataSetID int64) (bool, error) {
	strSql := "delete from dp_data_set where user_did=? and data_set_id=?"
	_, err := this.Dbconnect.Exec(strSql, userDid, dataSetID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (this *DBCon) RevokeDPDataSet(userDid, afterStatus, beforeStatus string, dataSetID int64) (bool, error) {
	res, err := this.QueryDPDataSetInfo(userDid, dataSetID)
	if err != nil {
		return false, err
	}
	if res == nil || res.Status != utils.VERIFYING {
		return false, fmt.Errorf("RevokeDPDataSet failed,current dpDataSet status not verifying")
	}
	return this.UpdateDPDataSetStatus(userDid, afterStatus, beforeStatus, dataSetID)
}

func (this *DBCon) UpdateDPDataSetStatus(userDid, afterStatus, beforeStatus string, dataSetID int64) (bool, error) {
	strsql := "update dp_data_set set status=? where user_did=? and  data_set_id=? and status=?"
	_, err := this.Dbconnect.Exec(strsql, afterStatus, userDid, dataSetID, beforeStatus)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (this *DBCon) QueryDPDataSetInfo(userDid string, dataSetID int64) (*model.DpDataSetRes, error) {
	var blockChainLabels string
	var categoryLabels string
	var scenarioLabels string
	strsql := "select data_set_id,dp_did, dataset_name,method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from dp_data_set where user_did=? and data_set_id=?"
	r, err := this.Dbconnect.Query(strsql, userDid, dataSetID)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &model.DpDataSetRes{
			Labels: &model.LabelsInfos{
				BlockChain: make([]string, 0),
				Category:   make([]string, 0),
				Scenario:   make([]string, 0),
			},
		}
		err := r.Scan(&t.DataSetID, &t.DpDid, &t.DataSetName, &t.DataSetMethodName, &t.DataSetDesc, &t.HTTPMethod, &t.HTTPURL, &t.Params, &t.Status, &t.ResDesc, &blockChainLabels, &categoryLabels, &scenarioLabels)
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

func (this *DBCon) QueryUserDPDataSetList(userDID, status string) ([]*model.DpDataSetRes, error) {
	var err error
	var r *sql.Rows
	if status == "" {
		strsql := "select data_set_id,dp_did, dataset_name, method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from dp_data_set where user_did=?"
		r, err = this.Dbconnect.Query(strsql, userDID)
	} else {
		strsql := "select data_set_id,dp_did, dataset_name, method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from dp_data_set where user_did=? and status=?"
		r, err = this.Dbconnect.Query(strsql, userDID, status)
	}
	if err != nil {
		return nil, err
	}
	dpDataSetReses := make([]*model.DpDataSetRes, 0)
	defer r.Close()
	for r.Next() {
		t := &model.DpDataSetRes{}
		var blockChainLabels string
		var categoryLabels string
		var scenarioLabels string
		err := r.Scan(&t.DataSetID, &t.DpDid, &t.DataSetName, &t.DataSetMethodName, &t.DataSetDesc, &t.HTTPMethod, &t.HTTPURL, &t.Params, &t.Status, &t.ResDesc, &blockChainLabels, &categoryLabels, &scenarioLabels)
		if err != nil {
			return nil, err
		}
		dpDataSetRes := &model.DpDataSetRes{
			DataSetID:         t.DataSetID,
			DpDid:             t.DpDid,
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
			dpDataSetRes.Labels.BlockChain = strings.Split(blockChainLabels, ",")
		}
		if categoryLabels != "" {
			dpDataSetRes.Labels.Category = strings.Split(categoryLabels, ",")
		}
		if scenarioLabels != "" {
			dpDataSetRes.Labels.Scenario = strings.Split(scenarioLabels, ",")
		}
		dpDataSetReses = append(dpDataSetReses, dpDataSetRes)
	}
	return dpDataSetReses, nil
}

func (this *DBCon) QueryUserDPDataSetLists(userDID, dataSetName, status string, page int64, size int64, labels model.LabelsInfo) (*model.DPDataSetList, error) {
	dPDataSetList := &model.DPDataSetList{
		DpDataSetData: make([]*model.DpDataSetRes, 0),
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
				strsql = "select data_set_id,dp_did, dataset_name, method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from dp_data_set where user_did=? order by create_time desc limit ? OFFSET ?"
			} else {
				strsql = "select data_set_id,dp_did, dataset_name, method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from dp_data_set where user_did=? and " + whereStr + "order by create_time desc limit ? OFFSET ?"
			}
			r, err = this.Dbconnect.Query(strsql, userDID, size, (page-1)*size)
		} else {
			strsql := "select data_set_id,dp_did, dataset_name, method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from dp_data_set where user_did=? and dataset_name like '%" + dataSetName + "%'" + whereStr +
				"order by create_time desc limit ? OFFSET ?"
			r, err = this.Dbconnect.Query(strsql, userDID, size, (page-1)*size)
		}
	} else {
		if dataSetName == "" {
			strsql := ""
			if whereStr == "" {
				strsql = "select data_set_id,dp_did,dataset_name, method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from dp_data_set where user_did=? and status=? order by create_time desc  limit ? OFFSET ?"
			} else {
				strsql = "select data_set_id,dp_did,dataset_name, method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from dp_data_set where user_did=? and status=? and " + whereStr + " order by create_time desc  limit ? OFFSET ?"
			}
			r, err = this.Dbconnect.Query(strsql, userDID, status, size, (page-1)*size)
		} else {
			strsql := "select data_set_id,dp_did, dataset_name, method_name,data_desc,http_method,http_url,params,status,res_desc,block_chain_labels,category_labels,scenario_labels from dp_data_set where user_did=? and status=? and dataset_name like '%" + dataSetName + "%'" + whereStr +
				"order by create_time desc  limit ? OFFSET ?"
			r, err = this.Dbconnect.Query(strsql, userDID, status, size, (page-1)*size)
		}
	}
	if err != nil {
		return nil, err
	}
	defer r.Close()
	for r.Next() {
		t := &model.DpDataSetRes{}
		var blockChainLabels string
		var categoryLabels string
		var scenarioLabels string
		err := r.Scan(&t.DataSetID, &t.DpDid, &t.DataSetName, &t.DataSetMethodName, &t.DataSetDesc, &t.HTTPMethod, &t.HTTPURL, &t.Params, &t.Status, &t.ResDesc, &blockChainLabels, &categoryLabels, &scenarioLabels)
		if err != nil {
			return nil, err
		}
		res := &model.DpDataSetRes{
			DataSetID:         t.DataSetID,
			DpDid:             t.DpDid,
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
		dPDataSetList.DpDataSetData = append(dPDataSetList.DpDataSetData, res)
	}
	var r1 *sql.Rows
	var err1 error
	if status == "" {
		if dataSetName == "" {
			strsql := "select count(data_set_id) from dp_data_set where user_did=?"
			r1, err1 = this.Dbconnect.Query(strsql, userDID)
		} else {
			strsql := "select count(data_set_id) from dp_data_set where user_did=? and dataset_name like '%" + dataSetName + "%'"
			r1, err1 = this.Dbconnect.Query(strsql, userDID)
		}
	} else {
		if dataSetName == "" {
			strsql := "select count(data_set_id)  from dp_data_set where user_did=? and status=?"
			r1, err1 = this.Dbconnect.Query(strsql, userDID, status)
		} else {
			strsql := "select count(data_set_id) from dp_data_set where user_did=? and status=? and dataset_name like '%" + dataSetName + "%'"
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
	dPDataSetList.CurPageNum = page
	dPDataSetList.TotalNum = totalNum
	return dPDataSetList, nil
}

func (this *DBCon) UpdateDPDataSetSchema(dpDid, dataSetName, inputSchema, outputSchema string) (bool, error) {
	strsql := "update dp_data_set set input_schema=?,output_schema=? " +
		"where  dp_did=? and method_name=?"
	_, err := this.Dbconnect.Exec(strsql, inputSchema, outputSchema, dpDid, dataSetName)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (this *DBCon) RevokePublishedDPDataSet(userDid string, dataSetID int64) (bool, error) {
	res, err := this.QueryDPDataSetInfo(userDid, dataSetID)
	if err != nil {
		return false, err
	}
	if res == nil {
		return false, fmt.Errorf("PublishedDPDataSet dataSetID:%d invalid", dataSetID)
	}
	if res.Status != utils.PUBLISHED {
		return false, fmt.Errorf("PublishedDPDataSet status:%s invalid", res.Status)
	}
	err = this.UpdateDPMethodStatus(res.DpDid, res.DataSetMethodName, utils.REVOKING)
	if err != nil {
		return false, err
	}
	return this.UpdateDPDataSetStatus(userDid, utils.REVOKING, utils.PUBLISHED, dataSetID)
}

func (db *DBCon) CheckDPDataSetNameExist(userDId, dataSetName string) (*model.DpDataSetRes, error) {
	strsql := "select data_set_id,dataset_name from dp_data_set where user_did=? and dataset_name=?"
	r, err := db.Dbconnect.Query(strsql, userDId, dataSetName)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &model.DpDataSetRes{}
		err := r.Scan(&t.DataSetID, &t.DataSetName)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, nil
}

func (db *DBCon) CheckDPDataSetMethodNameExist(userDId, methodName string) (*model.DpDataSetRes, error) {
	strsql := "select data_set_id,dataset_name from dp_data_set where user_did=? and method_name=?"
	r, err := db.Dbconnect.Query(strsql, userDId, methodName)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &model.DpDataSetRes{}
		err := r.Scan(&t.DataSetID, &t.DataSetName)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, nil
}

func (db *DBCon) CheckDPDataSetExistByDPDID(apDid string) (bool, error) {
	strsql := "select dataset_name from dp_data_set where dp_did=? and status=?"
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

func (this *DBCon) QueryDPDatatSetCountByCondition(userDId string) (int64, error) {
	strsql := "select count(user_did) from dp_data_set where user_did=? and status=?"
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
