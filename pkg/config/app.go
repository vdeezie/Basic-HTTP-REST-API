package config

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

var (
	db * sql.DB
)

func Connect() {
	db, err := sql.Open("mysql", "root:595983Fc@tcp(127.0.0.1:3306)bookservicedb")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *sql.DB {
	return db
}
