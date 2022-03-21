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

func (db *DBCon) AddUUIDNonce(uuid string, action int) error {
	strsql := "insert into uuid_nonce(uuid,action,create_time) values(?,?,sysdate())"
	_, err := db.Dbconnect.Exec(strsql, uuid, action)
	return err
}

func (db *DBCon) QueryUUIDAction(uuid string) (int, error) {
	strsql := "select action from uuid_nonce where uuid = ?"
	r, err := db.Dbconnect.Query(strsql, uuid)
	if err != nil {
		return -1, err
	}
	defer r.Close()
	if r.Next() {
		t := 0
		err := r.Scan(&t)
		if err != nil {
			return -1, err
		}
		return t, nil
	}
	return -1, nil
}
