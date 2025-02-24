package db

import (
	"database/sql"
	"fmt"
	"os"
	"rest-api/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB stores a pointer to the database
var DB *sql.DB

// InitDB establishes the connection to the database and loads the data that is in the environment variables
func InitDB() (*sql.DB, error){
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	stringConnection := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	utils.SecretKey = os.Getenv("SECRET_KEY")
	
	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}