package models

import (
	"BookServiceRESTAPI/pkg/config"
	"database/sql"
)

var db *sql.DB

type Book struct {
	sql.Model
	//Id          string `json:"id"`
	Name        string `sql:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func  GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}


func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(book)
	return book
}