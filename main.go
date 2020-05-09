package main

import (
	"fmt"
	"net/http"	
)

type books map [string]string

func (b books) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.String() {
	case "/home":
		fmt.Fprintf(w, "Welcome to book server. You came via %s\n", r.URL.String())
	case "/about":
		fmt.Fprintf(w, "All about our library of books. You came via %s\n", r.URL.String())
	case "/book":
		item := r.URL.Query().Get("item")
		book, ok := b[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "book not found: %s\n", item)
		}
		fmt.Fprintf(w, "The book is %s", book)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "url not found! %s\n", r.URL.String())
	}
	
}

func main() {
	//http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "Hello, Go HTTP\n")
	//})

	//http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "About, HTTP in Go\n")
	//})


	bk := books{}
	bk ["jrrtolkein"] = "The Silmarilion"
	fmt.Println("server listening on port 8000")
	http.ListenAndServe(":8000", bk)
	
}
