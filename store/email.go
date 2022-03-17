package store

import "fmt"

func (db *DBCon) AddNewEmailVerificationCode(did, email, vcode string) error {
	strsql := "insert into email_verification_code(user_did,email,verification_code,update_time) values(?,?,?,sysdate())"
	_, err := db.Dbconnect.Exec(strsql, did, email, vcode)
	return err
}

func (db *DBCon) UpdateEmailVerificationCode(did, email, newcode string) error {
	strsql := "update email_verification_code set verification_code = ?,update_time = sysdate() where user_did = ? and email = ? and sysdate() - update_time > 30"
	l, err := db.Dbconnect.Exec(strsql, newcode, did, email)
	if err != nil {
		return err
	}
	r, err := l.RowsAffected()
	if err != nil {
		return err
	}
	if r != 1 {
		return fmt.Errorf("update verification code failed: should after 30 sec and retry")
	}
	return nil
}

func (db *DBCon) GetEmailVerificationCode(did, email string) (string, error) {
	strsql := "select verification_code from email_verification_code where user_did = ? and email = ?"
	r, err := db.Dbconnect.Query(strsql, did, email)
	if err != nil {
		return "", err
	}
	defer r.Close()
	t := ""
	if r.Next() {
		err = r.Scan(&t)
	}
	return t, err
}
