package store

func (db *DBCon) GetNFTSettingByCondition(where string) ([]*NFTSetting, error) {
	strsql := "select * from nft_setting " + where
	r, err := db.Dbconnect.Query(strsql)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	res := make([]*NFTSetting, 0)
	for r.Next() {
		t := &NFTSetting{}

		err = r.Scan(&t.Id, &t.Name, &t.Description, &t.Image, &t.DpDID, &t.DpMethod, &t.ApDID, &t.ApMethod, &t.LowestScore, &t.ValidDays, &t.Restriction, &t.IssueBy, &t.AltImage)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (db *DBCon) GetNFTSettingCountByCondition(where string) (int64, error) {
	strsql := "select count(*) from nft_setting " + where

	r, err := db.Dbconnect.Query(strsql)
	if err != nil {
		return 0, err
	}
	defer r.Close()
	res := int64(0)
	if r.Next() {
		err = r.Scan(&res)
	}
	return res, err
}
