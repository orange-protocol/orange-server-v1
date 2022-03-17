package store

import (
	"github.com/orange-protocol/orange-server-v1/graph/model"
)

func (db *DBCon) SaveThirdPartyVc(did, mediaType, credential string) error {
	strsql := "insert into third_party_vc(did,media_type,credential) values (?,?,?)" +
		" ON DUPLICATE KEY UPDATE credential=?"
	_, err := db.Dbconnect.Exec(strsql, did, mediaType, credential, credential)
	return err
}

func (db *DBCon) QueryThirdPartyVc(did, mediaType string) (string, error) {
	strsql := "select credential from third_party_vc where did=? and media_type=?"
	r, err := db.Dbconnect.Query(strsql, did, mediaType)
	if err != nil {
		return "", err
	}
	defer r.Close()
	credential := ""
	if r.Next() {
		err = r.Scan(&credential)
		return credential, err
	}
	return "", nil
}

//MediaType: BrightID  Twitter ShuftiPro  Github  Linkedin  Facebook  Line  Amazon  Kakao
type ThirdPartyCredential struct {
	Did        string `json:"did"`
	MediaType  string `json:"media_type"`
	Credential string `json:"credential"`
}

func (db *DBCon) QueryAllThirdPartyVcStatus(did string) ([]*model.ThirdPartyVcStatus, error) {
	strsql := "select *from third_party_vc where did=?"
	r, err := db.Dbconnect.Query(strsql, did)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	res := make([]*model.ThirdPartyVcStatus, 0)
	for r.Next() {
		t := &ThirdPartyCredential{}
		err = r.Scan(&t.Did, &t.MediaType, &t.Credential)
		if err != nil {
			return nil, err
		}
		if t.Credential != "" {
			res = append(res, &model.ThirdPartyVcStatus{
				MediaType: t.MediaType,
				Status:    true,
			})
		} else {
			res = append(res, &model.ThirdPartyVcStatus{
				MediaType: t.MediaType,
				Status:    false,
			})
		}
	}
	return res, nil
}

func (db *DBCon) GetThirdPartyVCCounts(did string) (int64, error) {
	strsql := "select count(*) from third_party_vc where did=?"
	r, err := db.Dbconnect.Query(strsql, did)
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

func (db *DBCon) UnBindThirdParty(did, mediaType string) error {
	strsql := "delete from third_party_vc where did=? and media_type=?"
	_, err := db.Dbconnect.Exec(strsql, did, mediaType)
	if err != nil {
		return err
	}
	if mediaType == "ShuftiPro" {
		strSql := "delete from user_kyc where did=?"
		_, err = db.Dbconnect.Exec(strSql, did)
	}
	return err
}

func (db *DBCon) SaveUserKycInfo(did, kyc string) error {
	strsql := "insert into user_kyc(did,kyc) values (?,?)" +
		" ON DUPLICATE KEY UPDATE kyc=?"
	_, err := db.Dbconnect.Exec(strsql, did, kyc, kyc)
	return err
}

func (db *DBCon) QueryUserKycInfo(did string) (string, error) {
	strsql := "select kyc from user_kyc where did=?"
	r, err := db.Dbconnect.Query(strsql, did)
	if err != nil {
		return "", err
	}
	defer r.Close()
	kyc := ""
	for r.Next() {
		err = r.Scan(&kyc)
		return kyc, err
	}
	return kyc, nil
}
