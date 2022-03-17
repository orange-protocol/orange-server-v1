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
