package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
	Router *mux.Router
)

//Connect to MySQL
func Connect() {

	d, err := sql.Open("mysql", "root:595983Fc@tcp(127.0.0.1:3306)/bookdb")
	if err != nil{
		panic(err.Error())
	}

	defer db.Close()
	db = d
}

func GetDB() *sql.DB {
	return db
}
