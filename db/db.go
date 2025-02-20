package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() (*sql.DB, error){
	db, err := sql.Open("mysql", "root:liedsonfsa@/rest_api?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}