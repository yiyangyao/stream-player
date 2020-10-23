package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:nathan@tcp(localhost:3306)/stream_media_player?charset=utf8mb4")
	if err != nil {
		panic(err.Error())
	}
}
