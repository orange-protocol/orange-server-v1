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
