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
