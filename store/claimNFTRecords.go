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

func (db *DBCon) AddNewClaimRecord(record *ClaimNFTRecord) error {
	strsql := "insert into claim_nft_records(tx_hash,chain,contract_address,nft_type,user_did,user_address,create_time,update_time,score) values " +
		"(?,?,?,?,?,?,sysdate(),sysdate(),?)"
	_, err := db.Dbconnect.Exec(strsql, record.TxHash, record.Chain, record.ContractAddress, record.NftType, record.UserDID, record.UserAddress, record.Score)
	return err
}

func (db *DBCon) QueryClaimRecordsByStatus(status int) ([]*ClaimNFTRecord, error) {
	strsql := "select * from claim_nft_records where status = ?"
	r, err := db.Dbconnect.Query(strsql, status)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	res := make([]*ClaimNFTRecord, 0)
	for r.Next() {
		t := &ClaimNFTRecord{}
		err = r.Scan(&t.TxHash, &t.Chain, &t.ContractAddress, &t.NftType, &t.UserDID, &t.UserAddress, &t.CreateTime, &t.UpdateTime, &t.Status, &t.Result, &t.TokenId, &t.Score)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (db *DBCon) QueryClaimRecordCountByStatus(status int) (int64, error) {
	strsql := "select count(*) from claim_nft_records where status = ?"
	r, err := db.Dbconnect.Query(strsql, status)
	if err != nil {
		return 0, err
	}
	defer r.Close()
	t := int64(0)

	if r.Next() {
		err = r.Scan(&t)
	}
	return t, err
}

func (db *DBCon) UpdateClaimRecordStatusByLimit(fromStatus, toStatus, limit int) (int64, error) {
	strsql := "update claim_nft_records set status = ?, update_time= sysdate() where status = ? limit ?  "
	l, err := db.Dbconnect.Exec(strsql, toStatus, fromStatus, limit)
	if err != nil {
		return 0, err
	}
	return l.RowsAffected()
}
func (db *DBCon) UpdateClaimRecordStatusByPK(fromStatus, toStatus int, txhash string, chain string) (int64, error) {
	strsql := "update claim_nft_records set status = ?, update_time= sysdate() where tx_hash = ? and chain= ? and status = ?  "
	l, err := db.Dbconnect.Exec(strsql, toStatus, txhash, chain, fromStatus)
	if err != nil {
		return 0, err
	}
	return l.RowsAffected()
}
func (db *DBCon) SetClaimRecordResultByPK(fromStatus, toStatus int, txhash string, chain string, result string, tokenId int64) (int64, error) {
	strsql := "update claim_nft_records set status = ?, update_time= sysdate(),result=? ,token_id = ? where tx_hash = ? and chain= ? and status = ?  "
	l, err := db.Dbconnect.Exec(strsql, toStatus, result, tokenId, txhash, chain, fromStatus)
	if err != nil {
		return 0, err
	}
	return l.RowsAffected()
}

func (db *DBCon) QueryClaimRecordsByCondition(where string) ([]*ClaimNFTRecord, error) {
	strsql := "select * from claim_nft_records " + where
	r, err := db.Dbconnect.Query(strsql)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	res := make([]*ClaimNFTRecord, 0)
	for r.Next() {
		t := &ClaimNFTRecord{}
		err = r.Scan(&t.TxHash, &t.Chain, &t.ContractAddress, &t.NftType, &t.UserDID, &t.UserAddress, &t.CreateTime, &t.UpdateTime, &t.Status, &t.Result, &t.TokenId, &t.Score)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (db *DBCon) QueryClaimRecordCountByCondition(where string) (int64, error) {
	strsql := "select count(*) from claim_nft_records " + where
	r, err := db.Dbconnect.Query(strsql)
	if err != nil {
		return 0, err
	}
	defer r.Close()
	t := int64(0)

	if r.Next() {
		err = r.Scan(&t)
	}
	return t, err
}

func (db *DBCon) UpdateClaimRecordResult(fromStatus, toStatus int, txhash string, chain string, result string, nftType, tokenId int64, oldTxHash string) (int64, error) {
	strsql := "update claim_nft_records set status = ?, update_time= sysdate(),result=? ,token_id = ?,tx_hash = ? where  chain= ? and nft_type=? and status = ? and tx_hash = ?"
	l, err := db.Dbconnect.Exec(strsql, toStatus, result, tokenId, txhash, chain, nftType, fromStatus, oldTxHash)
	if err != nil {
		return 0, err
	}
	return l.RowsAffected()
}
