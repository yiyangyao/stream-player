package db

import (
	"database/sql"
	"log"
	"strconv"
	"stream-player/src/nathan.com/video-server/api/defs"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

func InsertSession(sessionId string, ttl int64, username string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id, ttl, login_name) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sessionId, ttlstr, username)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sessionId string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("select session_id, ttl, login_name from sessions where sessionId = ?")
	if err != nil {
		return nil, err
	}
	var ttl string
	var login_name string
	err = stmtOut.QueryRow(sessionId).Scan(&ttl, &login_name)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.UserName = login_name
	} else {
		return nil, err
	}
	defer stmtOut.Close()
	return ss, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	stmtOut, err := dbConn.Prepare("select session_id, ttl, login_name from sessions")
	var sessionMap = &sync.Map{}
	rows, err := stmtOut.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id, ttl, loginName string
		if err := rows.Scan(&id, &ttl, &loginName); err != nil {
			break
		}

		if ttl, err := strconv.ParseInt(ttl, 10, 64); err == nil {
			ss := &defs.SimpleSession{
				UserName: loginName, TTL: ttl,
			}
			sessionMap.Store(id, ss)
		}
	}
	defer stmtOut.Close()
	return sessionMap, nil
}

func DeleteSession(sid string) error {
	stmtDel, err := dbConn.Prepare("delete from sessions where session_id = ?")
	if err != nil {
		log.Printf("%s", err)
	}
	_, err = stmtDel.Exec(sid)
	if err != nil {
		return err
	}
	defer stmtDel.Close()

	return nil
}
