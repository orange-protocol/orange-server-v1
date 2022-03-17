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
