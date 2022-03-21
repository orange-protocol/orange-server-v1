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
	"fmt"
	"github.com/orange-protocol/orange-server-v1/graph/model"
	"strings"
)

func (db *DBCon) AddNewBasicInfo(did, nickName, email, avatar string) error {
	strsql := "insert into user_basic_info(did,nick_name,avatar,email,create_time,update_time) " +
		"values(?,?,?,?,sysdate(),sysdate())"

	_, err := db.Dbconnect.Exec(strsql, did, nickName, avatar, email)
	return err
}

func (db *DBCon) UpdateBasicInfo(did, nickName, email string) error {
	strsql := "update user_basic_info set "
	s := ""
	if len(nickName) > 0 {
		s = "nick_name = '" + nickName + "'"
	}
	if len(email) > 0 {
		and := ""
		if len(s) > 0 {
			and = " , "
		}
		s = s + and + " email = '" + email + "' "
	}
	strsql = strsql + s + " where did = '" + did + "'"

	//log.Debugf("UpdateBasicInfo strsql:%s\n", strsql)
	_, err := db.Dbconnect.Exec(strsql)
	return err

}

func (db *DBCon) GetUserBasicInfo(did string) (*UserBasicInfo, error) {
	strsql := "select * from user_basic_info where did = ?"
	r, err := db.Dbconnect.Query(strsql, did)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &UserBasicInfo{}
		var avar sql.NullString
		err = r.Scan(&t.Did, &t.NickName, &avar, &t.Email, &t.CreateTime, &t.UpdateTime)
		if err != nil {
			return nil, err
		}
		if avar.Valid {
			t.Avatar = avar.String
		}
		return t, nil
	}
	return nil, nil
}

func (db *DBCon) EditNickNameBasicInfo(did, nickName string) error {
	strsql := "insert into user_basic_info(did,nick_name) values (?,?) " +
		" ON DUPLICATE KEY UPDATE nick_name=?"
	_, err := db.Dbconnect.Exec(strsql, did, nickName, nickName)
	return err
}

func (db *DBCon) EditEmailAddrBasicInfo(did, email, verifyCode string) error {
	vcode, err := db.GetEmailVerificationCode(did, email)
	if err != nil {
		return err
	}
	if vcode != verifyCode {
		return fmt.Errorf("Email Verify Code errror")
	}
	strsql := "insert into user_basic_info(did,email) values (?,?) " +
		" ON DUPLICATE KEY UPDATE email=?"
	_, err = db.Dbconnect.Exec(strsql, did, email, email)
	return err
}

func (db *DBCon) GetDataSetLabels(labelType string) (*model.LabelsInfos, error) {
	strsql := "select block_chain,category,scenario from data_set_labels where label_type=?"
	r, err := db.Dbconnect.Query(strsql, labelType)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	labelsInfo := &LabelsInfo{}
	if r.Next() {
		err = r.Scan(&labelsInfo.BlockChain, &labelsInfo.Category, &labelsInfo.Scenario)
		if err != nil {
			return nil, err
		}
	}
	return &model.LabelsInfos{
		BlockChain: strings.Split(labelsInfo.BlockChain, ","),
		Category:   strings.Split(labelsInfo.Category, ","),
		Scenario:   strings.Split(labelsInfo.Scenario, ","),
	}, nil
}
