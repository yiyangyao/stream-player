package db

import (
	"database/sql"
	"log"
	"stream-player/src/nathan.com/video-server/display/user/consts"
	"stream-player/src/nathan.com/video-server/display/util"
)

func AddUserCredential(username string, password string) error {
	stmtIns, err := util.DBConn.Prepare("INSERT INTO users (username, passwd) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(username, password)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(username string) (*consts.UserCredential, error) {
	stmtOut, err := util.DBConn.Prepare("select id, passwd from users where username = ?")
	if err != nil {
		log.Printf("%s", err)
	}
	var userId int
	var password string
	err = stmtOut.QueryRow(username).Scan(&userId, &password)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer stmtOut.Close()

	userInfo := &consts.UserCredential{
		UserID:   userId,
		UserName: username,
		PassWord: password,
	}

	return userInfo, nil
}

func DeleteUser(username string, password string) error {
	stmtDel, err := util.DBConn.Prepare("delete from users where username = ? and passwd = ?")
	if err != nil {
		log.Printf("%s", err)
	}
	_, err = stmtDel.Exec(username, password)
	if err != nil {
		return err
	}
	defer stmtDel.Close()

	return nil
}
