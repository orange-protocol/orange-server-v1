package store

import "database/sql"

func (db *DBCon) AddTaskHistory(th *TaskHistory) error {
	strsql := "insert into task_history (task_id,user_did,ap_did,ap_method,dp_did,dp_method,create_time,update_time,task_status,task_result,issue_txhash,revoke_txhash,caller_did) " +
		"values (?,?,?,?,?,?,sysdate(),sysdate(),?,?,?,?)"
	_, err := db.Dbconnect.Exec(strsql, th.TaskId, th.UserDID, th.ApDID, th.ApMethod, th.DpDID, th.DpMethod, th.TaskStatus, th.TaskResult, th.IssueTxhash, th.RevokeTxhash, th.callerDid)
	return err
}

func (db *DBCon) AddTaskHistoryByTaskInfo(taskinfo *TaskInfo) error {
	strsql := "insert into task_history (task_id,user_did,ap_did,ap_method,dp_did,dp_method,create_time,update_time,task_status,task_result,issue_txhash,revoke_txhash,caller_did) " +
		"values (?,?,?,?,?,?,sysdate(),sysdate(),?,?,?,?,?)"
	_, err := db.Dbconnect.Exec(strsql, taskinfo.TaskId, taskinfo.UserDID, taskinfo.ApDID, taskinfo.ApMethod, taskinfo.DpDID, taskinfo.DpMethod, taskinfo.TaskStatus, taskinfo.TaskResult, taskinfo.IssueTxhash, taskinfo.RevokeTxhash, taskinfo.CallerDid)
	return err
}

func (db *DBCon) UpdateTaskHistoryByTaskInfo(taskinfo *TaskInfo) error {
	strsql := "update task_history set task_result = ? ,task_status = ? where task_id = ?"
	_, err := db.Dbconnect.Exec(strsql, taskinfo.TaskResult, taskinfo.TaskStatus, taskinfo.TaskId)
	return err
}

func (db *DBCon) UpdateTaskHistoryTxhash(taskid int64, txhash string) error {
	strsql := "update task_history set  issue_txhash = ? ,task_status = ? where task_id = ?"
	_, err := db.Dbconnect.Exec(strsql, txhash, TASK_STATUS_DONE, taskid)
	return err
}

func (db *DBCon) QueryTaskHistoryByUserDID(userdid string) ([]*TaskHistory, error) {
	strsql := "select * from task_history where user_did = ?"
	r, err := db.Dbconnect.Query(strsql, userdid)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	res := make([]*TaskHistory, 0)
	for r.Next() {
		t := &TaskHistory{}
		tmp := sql.NullString{}
		err := r.Scan(&t.TaskId, &t.UserDID, &t.ApDID, &t.ApMethod, &t.DpDID, &t.DpMethod, &t.CreateTime, &t.UpdateTime, &t.TaskStatus, tmp, &t.IssueTxhash, &t.RevokeTxhash, &t.callerDid)
		if err != nil {
			return nil, err
		}
		if tmp.Valid {
			t.TaskResult = tmp.String
		}
		res = append(res, t)
	}

	return res, nil
}

func (db *DBCon) QueryTaskHistoryCountByUserDID(did string) (int64, error) {
	strsql := "select count(*) from task_history where user_did = ?"
	r, err := db.Dbconnect.Query(strsql, did)
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
