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

func (this *DBCon) AddUserOScoreInfo(ai *UserOscoreInfo) error {
	strsql := "insert into user_oscore_info (did,oscore,ap_did,dp_did,create_time) values (?,?,?,?,sysdate())"
	_, err := this.Dbconnect.Exec(strsql, ai.Did, ai.Oscore, ai.ApDid, ai.DpDid)
	return err
}

func (this *DBCon) UpdateUserLatestOscore(ai *UserOscoreInfo) error {
	strsql := "update user_oscore_info set oscore=? , ap_did=? , dp_did = ? ,create_time = sysdate() where did = ?"
	_, err := this.Dbconnect.Exec(strsql, ai.Oscore, ai.ApDid, ai.DpDid, ai.Did)
	return err
}

func (this *DBCon) QueryUserLatestOScoreInfo(did string) (*UserOscoreInfo, error) {
	strsql := "select * from user_oscore_info where did= ?"
	r, err := this.Dbconnect.Query(strsql, did)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	if r.Next() {
		t := &UserOscoreInfo{}
		err = r.Scan(&t.Did, &t.Oscore, &t.ApDid, &t.DpDid, &t.CreateTime)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, nil
}
