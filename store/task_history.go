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
