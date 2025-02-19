package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("mysql", "root:liedsonfsa@/rest_api?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}