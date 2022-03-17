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
