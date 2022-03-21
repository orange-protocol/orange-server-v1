/*
 *
 *  * Copyright (C) 2022 The orange protocol Authors
 *  * This file is part of The orange library.
 *  *
 *  * The Orange is free software: you can redistribute it and/or modify
 *  * it under the terms of the GNU Lesser General Public License as published by
 *  * the Free Software Foundation, either version 3 of the License, or
 *  * (at your option) any later version.
 *  *
 *  * The orange is distributed in the hope that it will be useful,
 *  * but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  * GNU Lesser General Public License for more details.
 *  *
 *  * You should have received a copy of the GNU Lesser General Public License
 *  * along with The orange.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package store

import (
	"database/sql"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/orange-protocol/orange-server-v1/config"
	"github.com/orange-protocol/orange-server-v1/log"

	_ "github.com/go-sql-driver/mysql"
)

type DBCon struct {
	Dbconnect *sqlx.DB
}

var MySqlDB *DBCon
var NEW_TASK_RESOLVE_COUNT = 100

func InitMysql(cfg *config.DB) error {
	path := strings.Join([]string{cfg.UserName, ":", cfg.Password, "@tcp(", cfg.DBAddr, ")/", cfg.DbName, "?charset=utf8&parseTime=true"}, "")
	db, err := sqlx.Open("mysql", path)
	if err != nil {
		log.Errorf("open database failed:%s", err)
		return err
	}
	db.SetConnMaxLifetime(time.Hour)
	db.SetMaxIdleConns(10)
	if err := db.Ping(); err != nil {
		log.Errorf("ping database failed:%s", err)
		return err
	}
	log.Info("connect database success")
	dbcon := &DBCon{
		Dbconnect: db,
	}
	MySqlDB = dbcon
	return nil
}

func (this *DBCon) AddUserAddressInfo(info *UserAddressInfo) error {
	strsql := "insert into user_address_info(did, chain_name,address,pubkey,create_time,visible) values (?,?,?,?,sysdate(),?)"
	_, err := this.Dbconnect.Exec(strsql, info.Did, info.Chain, info.Address, info.Pubkey, info.Visible)
	if err != nil {
		log.Errorf("error on AddUserAddressInfo:%s", err.Error())
		return err
	}
	return nil
}

func (this *DBCon) DeleteUserAddressInfo(did, chain, address string) error {
	strsql := "delete from user_address_info where did = ? and chain_name = ? and address = ?"
	_, err := this.Dbconnect.Exec(strsql, did, chain, address)
	if err != nil {
		log.Errorf("error on DeleteUserAddressInfo:%s", err.Error())
		return err
	}
	return nil
}

func (this *DBCon) GetUserAddressInfo(did string) ([]*UserAddressInfo, error) {
	strsql := "select * from user_address_info where did = ? order by create_time desc"
	r, err := this.Dbconnect.Query(strsql, did)
	if err != nil {
		log.Errorf("error on GetUserAddressInfo:%s", err.Error())
		return nil, err
	}
	defer r.Close()

	res := make([]*UserAddressInfo, 0)
	for r.Next() {
		t := &UserAddressInfo{}
		err = r.Scan(&t.Did, &t.Chain, &t.Address, &t.Pubkey, &t.CreateTime, &t.Visible)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, err
}

func (this *DBCon) GetUserVisibleAddressInfo(did string) ([]*UserAddressInfo, error) {
	strsql := "select * from user_address_info where did = ? and visible=1 order by create_time desc"
	r, err := this.Dbconnect.Query(strsql, did)
	if err != nil {
		log.Errorf("error on GetUserAddressInfo:%s", err.Error())
		return nil, err
	}
	defer r.Close()

	res := make([]*UserAddressInfo, 0)
	for r.Next() {
		t := &UserAddressInfo{}
		err = r.Scan(&t.Did, &t.Chain, &t.Address, &t.Pubkey, &t.CreateTime, &t.Visible)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, err
}

func (this *DBCon) GetUserAddressInfoCount(did string) (int64, error) {
	strsql := "select count(*) from user_address_info where did=?"
	r, err := this.Dbconnect.Query(strsql, did)
	if err != nil {
		log.Errorf("error on GetUserAddressInfo:%s", err.Error())
		return 0, err
	}
	defer r.Close()

	res := int64(0)
	if r.Next() {
		err = r.Scan(&res)
	}

	return res, err

}

func (this *DBCon) IsUserBindLoginAddress(did string, addr string) (bool, error) {
	strsql := "select count(*) from user_address_info where did=? and address=?"
	r, err := this.Dbconnect.Query(strsql, did, addr)
	if err != nil {
		log.Errorf("error on GetUserAddressInfo:%s", err.Error())
		return false, err
	}
	defer r.Close()

	res := int64(0)
	if r.Next() {
		err = r.Scan(&res)
	}
	return res > 0, err
}

func (this *DBCon) AddTask(userdid, ap_did, ap_method, dp_did, dp_method, callerdid, bind_info string) (int64, error) {
	strsql := "insert into task_info(user_did,ap_did,ap_method,dp_did,dp_method,create_time,update_time,task_status,caller_did,task_bind_info) values (?,?,?,?,?,sysdate(),sysdate(),0,?,?)"
	r, err := this.Dbconnect.Exec(strsql, userdid, ap_did, ap_method, dp_did, dp_method, callerdid, bind_info)
	if err != nil {
		return -1, err
	}
	taskid, err := r.LastInsertId()
	if err != nil {
		return -1, err
	}
	return taskid, nil
}

func (this *DBCon) DeleteTask(userdid, ap_did, ap_method, dp_did, dp_method string) error {
	strsql := "delete from task_info where user_did = ? and ap_did = ? and ap_method = ? and dp_did = ? and dp_method = ?"
	_, err := this.Dbconnect.Exec(strsql, userdid, ap_did, ap_method, dp_did, dp_method)
	return err
}

func (this *DBCon) DeleteTaskById(taskid int64) error {
	strsql := "delete from task_info where task_id = ?"
	_, err := this.Dbconnect.Exec(strsql, taskid)
	return err
}

func (this *DBCon) QueryTaskByUniqueKey(userdid, ap_did, ap_method, dp_did, dp_method string) (*TaskInfo, error) {
	strsql := "select * from task_info where user_did = ? and ap_did = ? and ap_method = ? and dp_did = ? and dp_method = ?"
	r, err := this.Dbconnect.Query(strsql, userdid, ap_did, ap_method, dp_did, dp_method)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &TaskInfo{}
		var dpresult, result, rfile, bind_info sql.NullString
		err = r.Scan(&t.TaskId, &t.UserDID, &t.ApDID, &t.ApMethod, &t.DpDID, &t.DpMethod, &dpresult, &t.CreateTime, &t.UpdateTime, &t.TaskStatus, &result, &rfile, &t.IssueTxhash, &t.RevokeTxhash, &t.CallerDid, &t.Comments, &bind_info)
		if err != nil {
			return nil, err
		}
		if dpresult.Valid {
			t.DpResult = dpresult.String
		}

		if result.Valid {
			t.TaskResult = result.String
		}
		if rfile.Valid {
			t.ResultFile = rfile.String
		}
		if bind_info.Valid {
			t.TaskBindInfo = bind_info.String
		}

		return t, nil
	}
	return nil, nil
}

func (this *DBCon) QueryTaskByStatus(status int) ([]*TaskInfo, error) {
	strsql := "select * from task_info where task_status = ? order by create_time limit ?"
	r, err := this.Dbconnect.Query(strsql, status, config.GlobalConfig.BatchTaskCount)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	res := make([]*TaskInfo, 0)
	for r.Next() {
		t := &TaskInfo{}
		var dpresult, result, rfile, bind_info sql.NullString
		err = r.Scan(&t.TaskId, &t.UserDID, &t.ApDID, &t.ApMethod, &t.DpDID, &t.DpMethod, &dpresult, &t.CreateTime, &t.UpdateTime, &t.TaskStatus, &result, &rfile, &t.IssueTxhash, &t.RevokeTxhash, &t.CallerDid, &t.Comments, &bind_info)
		if err != nil {
			return nil, err
		}
		if dpresult.Valid {
			t.DpResult = dpresult.String
		}
		if result.Valid {
			t.TaskResult = result.String
		}
		if rfile.Valid {
			t.ResultFile = rfile.String
		}
		if bind_info.Valid {
			t.TaskBindInfo = bind_info.String
		}
		res = append(res, t)

	}
	return res, nil
}

func (this *DBCon) LockTaskStatus(fromStatus, toStatus int, count int64) (int64, error) {
	strsql := "update task_info set task_status = ?,update_time=sysdate() where task_status = ? order by create_time limit ?"
	r, err := this.Dbconnect.Exec(strsql, toStatus, fromStatus, count)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

func (this *DBCon) GetInQueryTaskCount() (int64, error) {
	strsql := "select count(*) from task_info where task_status in(?,?)"
	r, err := this.Dbconnect.Query(strsql, TASK_STATUS_RESOLVING, TASK_STATUS_DP_QUERYING)
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

func (this *DBCon) GetUserReputationCount(did string) (int64, error) {
	strsql := "select count(*) from task_info where user_did=?"
	r, err := this.Dbconnect.Query(strsql, did)
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

func (this *DBCon) ResetTimeOutTasks(cnt int) error {
	strsql := "update task_info set task_status = ? ,update_time = sysdate() where task_status in (?,?) and sysdate() -  update_time > ? order by create_time limit ? "
	_, err := this.Dbconnect.Exec(strsql, TASK_STATUS_INIT, TASK_STATUS_DP_QUERYING, TASK_STATUS_DP_FAILED, config.GlobalConfig.TimeOutSeconds, cnt)
	return err
}

func (this *DBCon) GetTimeOutTasks() ([]*TaskInfo, error) {
	strsql := "select * from task_info where task_status = ? and sysdate() -  update_time > ? order by create_time limit ?"
	r, err := this.Dbconnect.Query(strsql, TASK_STATUS_DP_QUERYING, config.GlobalConfig.TimeOutSeconds, config.GlobalConfig.BatchTaskCount)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	res := make([]*TaskInfo, 0)
	for r.Next() {
		t := &TaskInfo{}
		var dpresult, result, rfile, bind_info sql.NullString
		err = r.Scan(&t.TaskId, &t.UserDID, &t.ApDID, &t.ApMethod, &t.DpDID, &t.DpMethod, &dpresult, &t.CreateTime, &t.UpdateTime, &t.TaskStatus, &result, &rfile, &t.IssueTxhash, &t.RevokeTxhash, &t.CallerDid, &t.Comments, &bind_info)
		if err != nil {
			return nil, err
		}
		if dpresult.Valid {
			t.DpResult = dpresult.String
		}

		if result.Valid {
			t.TaskResult = result.String
		}
		if rfile.Valid {
			t.ResultFile = rfile.String
		}
		if bind_info.Valid {
			t.TaskBindInfo = bind_info.String
		}
		res = append(res, t)

	}
	return res, nil

}

func (this *DBCon) ChangeTaskStatus(taskid int64, fromStatus, toStatus int) (int64, error) {
	strsql := "update task_info set task_status = ?,update_time=sysdate() where task_id=? and task_status = ?"
	r, err := this.Dbconnect.Exec(strsql, toStatus, taskid, fromStatus)
	if err != nil {
		return 0, err
	}
	return r.RowsAffected()
}

func (this *DBCon) QueryTasksByUserDID(userdid string) ([]*TaskInfo, error) {
	//strsql := "select * from task_info where user_did = ? order by create_time "
	strsql := `select t1.*,t3.ap_name,t3.title,t3.icon,t2.name as ap_method_name,t5.dp_name,t5.title,t5.icon,t4.name as dp_method_name
			from (task_info t1
				left join algorithm_provider_method_info t2 on t1.ap_did = t2.did and t1.ap_method = t2.method
				left join algorithm_provider_info t3 on t1.ap_did = t3.did
				left join data_provider_method_info t4 on t1.dp_did = t4.did and t1.dp_method = t4.method
				left join data_provider_info t5 on t1.dp_did = t5.did)
			where t1.user_did = ?
			order by create_time desc`
	r, err := this.Dbconnect.Query(strsql, userdid)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	res := make([]*TaskInfo, 0)
	for r.Next() {
		t := &TaskInfo{}
		var dpresult, result, rfile, dmn, bind_info sql.NullString
		var apname, aptitle, apicon, apmethodname sql.NullString
		//err = r.Scan(&t.TaskId, &t.UserDID, &t.ApDID, &t.ApMethod, &t.DpDID, &t.DpMethod, &dpresult, &t.CreateTime, &t.UpdateTime, &t.TaskStatus, &result, &rfile, &t.IssueTxhash, &t.RevokeTxhash, &t.CallerDid, &t.ApName, &t.ApTitle, &t.ApIcon, &t.APMethodName, &t.DpName, &t.DpTitle, &t.DpIcon, &dmn)
		err = r.Scan(&t.TaskId, &t.UserDID, &t.ApDID, &t.ApMethod, &t.DpDID, &t.DpMethod, &dpresult, &t.CreateTime, &t.UpdateTime, &t.TaskStatus, &result, &rfile, &t.IssueTxhash, &t.RevokeTxhash, &t.CallerDid, &t.Comments, &bind_info, &apname, &aptitle, &apicon, &apmethodname, &t.DpName, &t.DpTitle, &t.DpIcon, &dmn)
		if err != nil {
			return nil, err
		}
		if dpresult.Valid {
			t.DpResult = dpresult.String
		}

		if result.Valid {
			t.TaskResult = result.String
		}
		if rfile.Valid {
			t.ResultFile = rfile.String
		}
		if bind_info.Valid {
			t.TaskBindInfo = bind_info.String
		}
		if dmn.Valid {
			t.DpMethodName = dmn.String
		}
		if apname.Valid {
			t.ApName = apname.String
		}
		if aptitle.Valid {
			t.ApTitle = aptitle.String
		}
		if apicon.Valid {
			t.APMethodName = apmethodname.String
		}

		res = append(res, t)

	}
	return res, nil
}

func (this *DBCon) QueryTaskByPK(taskid int64) (*TaskInfo, error) {
	//strsql := "select t1.*,t2.ap_name,t3.dp_name from (task_info t1 left join algorithm_provider_info t2 on t1.ap_did = t2.did) left join data_provider_info t3 on t1.dp_did = t3.did where t1.task_id = ? "
	strsql := `select t1.*,t3.ap_name,t3.title,t3.icon,t2.name as ap_method_name,t5.dp_name,t5.title,t5.icon,t4.name as dp_method_name
			from (task_info t1
				left join algorithm_provider_method_info t2 on t1.ap_did = t2.did and t1.ap_method = t2.method
				left join algorithm_provider_info t3 on t1.ap_did = t3.did
				left join data_provider_method_info t4 on t1.dp_did = t4.did and t1.dp_method = t4.method
				left join data_provider_info t5 on t1.dp_did = t5.did) where t1.task_id=?`
	r, err := this.Dbconnect.Query(strsql, taskid)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &TaskInfo{}
		var dpresult, result, rfile, bind_info, apname, dpname sql.NullString
		var title, icon sql.NullString
		err = r.Scan(&t.TaskId, &t.UserDID, &t.ApDID, &t.ApMethod, &t.DpDID, &t.DpMethod, &dpresult, &t.CreateTime, &t.UpdateTime, &t.TaskStatus, &result, &rfile, &t.IssueTxhash, &t.RevokeTxhash, &t.CallerDid, &t.Comments, &bind_info,
			&apname, &title, &icon, &t.APMethodName, &dpname, &t.DpTitle, &t.DpIcon, &t.DpMethodName)
		if err != nil {
			return nil, err
		}
		if dpresult.Valid {
			t.DpResult = dpresult.String
		}

		if result.Valid {
			t.TaskResult = result.String
		}
		if rfile.Valid {
			t.ResultFile = rfile.String
		}
		if bind_info.Valid {
			t.TaskBindInfo = bind_info.String
		}
		if apname.Valid {
			t.ApName = apname.String
		}
		if dpname.Valid {
			t.DpName = dpname.String
		}
		if title.Valid {
			t.ApTitle = title.String
		}
		if icon.Valid {
			t.ApIcon = icon.String
		}

		return t, nil
	}
	return nil, nil
}

func (this *DBCon) SetTaskCredFileAndTxhash(taskid int64, credfile string, txhash string) error {
	strsql := "update task_info set result_file = ?, issue_txhash = ? ,task_status = ?,update_time=sysdate() where task_id = ?"
	_, err := this.Dbconnect.Exec(strsql, credfile, txhash, TASK_STATUS_DONE, taskid)
	return err
}

func (this *DBCon) SetDPResult(taskid int64, dpresult string) error {
	strsql := "update task_info set dp_result = ? , task_status = ?,update_time=sysdate() where task_id = ? and task_status in (? ,?, ?)"
	_, err := this.Dbconnect.Exec(strsql, dpresult, TASK_STATUS_DP_FINISHED, taskid, TASK_STATUS_DP_QUERYING, TASK_STATUS_DP_FAILED, TASK_STATUS_MUTLI_DP_FINISHED)
	return err
}

func (this *DBCon) AppendDPResult(taskid int64, dpresult string, isLast bool, method string) error {
	strsql := "update task_info set dp_result = CONCAT_WS(';;',dp_result,?),task_status=?,update_time=sysdate() where task_id=? and task_status in (?,?,?) "

	s := TASK_STATUS_DP_FINISHED
	if !isLast {
		s = TASK_STATUS_PARTITIAL_FINISHED
	}
	_, err := this.Dbconnect.Exec(strsql, method+"::"+dpresult, s, taskid, TASK_STATUS_DP_QUERYING, TASK_STATUS_DP_FAILED, TASK_STATUS_PARTITIAL_FINISHED)
	return err
}

func (this *DBCon) SetDPResultFailed(taskid int64, dpresult string) error {
	strsql := "update task_info set dp_result = ? , task_status = ?,update_time=sysdate() where task_id = ? and task_status = ? "
	_, err := this.Dbconnect.Exec(strsql, dpresult, TASK_STATUS_DP_FAILED, taskid, TASK_STATUS_DP_QUERYING)
	return err
}

func (this *DBCon) SetTaskComments(taskid int64, comments string) error {
	strsql := "update task_info set comments = ? ,update_time=sysdate() where task_id = ?  "
	_, err := this.Dbconnect.Exec(strsql, comments, taskid)
	return err
}

func (this *DBCon) SetAPResult(taskid int64, apresult string) error {
	strsql := "update task_info set task_result = ? , task_status = ?,update_time=sysdate() where task_id = ? and (task_status = ? or task_status = ?)"
	_, err := this.Dbconnect.Exec(strsql, apresult, TASK_STATUS_AP_FINISHED, taskid, TASK_STATUS_AP_QUERYING, TASK_STATUS_AP_FAILED)
	return err
}

func (this *DBCon) SetAPResultFailed(taskid int64, apresult string) error {
	strsql := "update task_info set task_result = ? , task_status = ?, update_time=sysdate() where task_id = ? and task_status = ? "
	_, err := this.Dbconnect.Exec(strsql, apresult, TASK_STATUS_AP_FAILED, taskid, TASK_STATUS_AP_QUERYING)
	return err
}

func (this *DBCon) SetAddressVisible(userDid, chain, address string, visible bool) error {
	strsql := "update user_address_info set visible = ? where did = ? and chain_name = ? and address = ?"
	_, err := this.Dbconnect.Exec(strsql, visible, userDid, chain, address)
	return err
}

func (this *DBCon) GetUserCredentialsCountByCondition(strWhere string) (int64, error) {
	strsql := `select count(*)
			from (task_info t1
				left join algorithm_provider_method_info t2 on t1.ap_did = t2.did and t1.ap_method = t2.method
				left join algorithm_provider_info t3 on t1.ap_did = t3.did
				left join data_provider_method_info t4 on t1.dp_did = t4.did and t1.dp_method = t4.method
				left join data_provider_info t5 on t1.dp_did = t5.did) ` + strWhere

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

func (this *DBCon) GetUserCredentialsByCondition(strWhere string) ([]*TaskInfo, error) {

	strsql := `select t1.*,t3.ap_name,t3.title,t3.icon,t2.name as ap_method_name,t5.dp_name,t5.title,t5.icon,t4.name as dp_method_name
			from (task_info t1
				left join algorithm_provider_method_info t2 on t1.ap_did = t2.did and t1.ap_method = t2.method
				left join algorithm_provider_info t3 on t1.ap_did = t3.did
				left join data_provider_method_info t4 on t1.dp_did = t4.did and t1.dp_method = t4.method
				left join data_provider_info t5 on t1.dp_did = t5.did) ` + strWhere

	r, err := this.Dbconnect.Query(strsql)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	res := make([]*TaskInfo, 0)
	for r.Next() {
		t := &TaskInfo{}
		var dpresult, result, rfile, bind_info, dmn sql.NullString
		err = r.Scan(&t.TaskId, &t.UserDID, &t.ApDID, &t.ApMethod, &t.DpDID, &t.DpMethod, &dpresult, &t.CreateTime, &t.UpdateTime, &t.TaskStatus, &result, &rfile, &t.IssueTxhash, &t.RevokeTxhash, &t.CallerDid, &t.Comments, &bind_info, &t.ApName, &t.ApTitle, &t.ApIcon, &t.APMethodName, &t.DpName, &t.DpTitle, &t.DpIcon, &dmn)
		if err != nil {
			return nil, err
		}
		if dpresult.Valid {
			t.DpResult = dpresult.String
		}

		if result.Valid {
			t.TaskResult = result.String
		}
		if rfile.Valid {
			t.ResultFile = rfile.String
		}
		if bind_info.Valid {
			t.TaskBindInfo = bind_info.String
		}
		if dmn.Valid {
			t.DpMethodName = dmn.String
		}

		res = append(res, t)

	}
	return res, nil

}

func (this *DBCon) QuerySnapShotAssetsScore(address string) (string, error) {
	userdid := strings.ReplaceAll(address, "0x", "did:etho:")
	strsql := "select task_result from task_info where user_did = ? and ap_did = ? and ap_method = ? and dp_did = ? and dp_method = ?"
	r, err := this.Dbconnect.Query(strsql, userdid, config.GlobalConfig.SnapShotAssetsConfig.ApDID, config.GlobalConfig.SnapShotAssetsConfig.ApMethod,
		config.GlobalConfig.SnapShotAssetsConfig.DpDID, config.GlobalConfig.SnapShotAssetsConfig.DpMethod)
	if err != nil {
		return "", err
	}
	defer r.Close()
	score := "0"
	for r.Next() {
		err = r.Scan(&score)
		return score, err
	}
	return score, nil
}

func (this *DBCon) QueryTaskInfo(userdid, dp_did, dp_method string) (*TaskInfo, error) {
	strsql := "select * from task_info where user_did = ? and dp_did = ? and dp_method = ?"
	r, err := this.Dbconnect.Query(strsql, userdid, dp_did, dp_method)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &TaskInfo{}
		var dpresult, result, rfile, bind_info sql.NullString
		err = r.Scan(&t.TaskId, &t.UserDID, &t.ApDID, &t.ApMethod, &t.DpDID, &t.DpMethod, &dpresult, &t.CreateTime, &t.UpdateTime, &t.TaskStatus, &result, &rfile, &t.IssueTxhash, &t.RevokeTxhash, &t.CallerDid, &t.Comments, &bind_info)
		if err != nil {
			return nil, err
		}
		if dpresult.Valid {
			t.DpResult = dpresult.String
		}

		if result.Valid {
			t.TaskResult = result.String
		}
		if rfile.Valid {
			t.ResultFile = rfile.String
		}
		if bind_info.Valid {
			t.TaskBindInfo = bind_info.String
		}

		return t, nil
	}
	return nil, nil
}
