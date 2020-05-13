package config

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
)

var connection *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:595983Fc@tcp(127.0.0.1:3306)bookservicedb")
	if err != nil {
		log.Fatal(err)
	} 

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Connected to MySQL database")
	connection = db
	return db
}