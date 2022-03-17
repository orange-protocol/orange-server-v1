package store

func (db *DBCon) AddFeedBack(fb *FeedBack) error {
	strsql := "insert into feedback_info(user_did,title,email,content,create_time) values(?,?,?,?,sysdate())"
	_, err := db.Dbconnect.Exec(strsql, fb.UserDID, fb.Title, fb.Email, fb.Content)
	return err
}
