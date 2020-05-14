package models

import (
	"BookService/BookService/pkg/config"
	"database/sql"
)

var db *sql.DB
var err error
//Book struct data
type Book struct { 
	Id          string `json:"id"`
	Name        string `sql:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

//CreateBook method for database 
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

//GetAllBooks method for database 
func  GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

//DeleteBook method deletes books on the databse
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(book)
	return book
}