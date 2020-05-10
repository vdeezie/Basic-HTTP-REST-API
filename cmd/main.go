package main

import (
	"html/template"
	"encoding/json"
	//"fmt"
	"net/http"
	"time"	
)

var books []Book

//Book model
type Book struct {
	Name		string	`json: "name"`
	Author		string	`json: "author"`
	PublishedAt	string	`json: "published_at"`
	
}

func loadBooks(w http.ResponseWriter, r *http.Request) {
	bk, err :=json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bk)
}

func create(w http.ResponseWriter, r*http.Request) {
	tmpl, _ := template.ParseFile("./assets/form.html")
	tmpl.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/books", loadBooks)

	fs := http.FileServer(http.Dir("assets/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	
	mux.HandleFunc("/new", create)
	http.ListenAndServe(":8000", mux)
}