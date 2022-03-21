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

func (db *DBCon) SaveApplicationInfo(did, name, website string) error {
	strsql := "insert into application_info(did,name,website) values (?,?,?) " +
		" ON DUPLICATE KEY UPDATE name=?, website=?"

	_, err := db.Dbconnect.Exec(strsql, did, name, website, name, website)
	return err
}

func (db *DBCon) QueryApplicationInfo(did string) (*ApplicationInfo, error) {
	strsql := "select * from application_info where did = ?"
	r, err := db.Dbconnect.Query(strsql, did)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if r.Next() {
		res := &ApplicationInfo{}
		err = r.Scan(&res.Did, &res.Name, &res.Website)
		return res, err
	}
	return nil, nil
}

func (db *DBCon) EditAppNameAppInfo(did, appName string) error {
	strsql := "insert into application_info(did,name) values (?,?) " +
		" ON DUPLICATE KEY UPDATE name=?"
	_, err := db.Dbconnect.Exec(strsql, did, appName, appName)
	return err
}

func (db *DBCon) EditWebsiteAppInfo(did, website string) error {
	strsql := "insert into application_info(did,website) values (?,?) " +
		" ON DUPLICATE KEY UPDATE website=?"
	_, err := db.Dbconnect.Exec(strsql, did, website, website)
	return err
}
