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

func (db *DBCon) GetGenNFTCountByDID(did string) (int64, error) {
	strsql := "select count(*) from gen_nft where user_did = ?"
	r, err := db.Dbconnect.Query(strsql, did)
	if err != nil {
		return 0, nil
	}
	defer r.Close()
	t := int64(0)
	if r.Next() {
		err = r.Scan(&t)
		if err != nil {
			return 0, nil
		}
	}
	return t, nil
}
