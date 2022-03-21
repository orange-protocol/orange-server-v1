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

func (this *DBCon) AddWasmCode(did, address, code, comments string) error {
	strsql := "insert into wasm_code_info(owner_did,address,code,create_time,comments) values(?,?,?,sysdate(),?)"
	_, err := this.Dbconnect.Exec(strsql, did, address, code, comments)
	if err != nil {
		return err
	}
	return nil
}

func (this *DBCon) QueryWasmCodeByDIDAndAddress(did, address string) (*WasmCodeInfo, error) {
	strsql := "select * from wasm_code_info where owner_did = ? and address = ?"
	r, err := this.Dbconnect.Query(strsql, did, address)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	if r.Next() {
		t := &WasmCodeInfo{}
		err = r.Scan(&t.OwnerDID, &t.Address, &t.Code, &t.CreateTime, &t.Comments)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, nil
}
