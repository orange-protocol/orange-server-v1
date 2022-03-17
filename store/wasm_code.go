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
