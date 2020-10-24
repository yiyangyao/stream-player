package util

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DBConn *sql.DB
	err    error
)

func init() {
	DBConn, err = sql.Open("mysql", "root:nathan@tcp(localhost:3306)/stream_media_player?charset=utf8mb4")
	if err != nil {
		panic(err.Error())
	}
}
