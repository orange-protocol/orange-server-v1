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
	"strings"
)

func (db *DBCon) QueryAuthInfoByDid(did string) (*AuthInfo, error) {
	strsql := "select * from auth_info where did = ?"
	r, err := db.Dbconnect.Query(strsql, did)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		t := &AuthInfo{}
		err = r.Scan(&t.Did, &t.AppName, &t.DataAuth, &t.AlgorithmAuth, &t.State)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, nil
}

func (db *DBCon) CheckProviderMethodMatch(apdid, apmethod, dpdid, dpmethod string) (bool, error) {
	dm, err := db.QueryDPMethodByDIDAndMethod(dpdid, dpmethod)
	if err != nil {
		return false, err
	}

	am, err := db.QueryAPMethodByDIDAndMethod(apdid, apmethod)
	if err != nil {
		return false, err
	}

	return strings.EqualFold(dm.ResultSchema, am.ParamSchema), nil

}
