package main

import (
	"fmt"
	"net/http"	
)

type Book struct
	Name		string	`json: "name"`
	Author		string	`json: "author"`
	PublishedAt	string	`json: "published_at"`
	
}

func loadBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{"The Silmarilion", "JRR Tolkein", time.Now().Local().String()}

	bk, err :=json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bk)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/books", loadBooks)

	fs := http.FileServer(http.Dir("assets/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	
	http.ListenAndServe(":8000", mux)
}
/*type books map [string]string

func (b books) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/home":
		fmt.Fprintf(w, "Welcome to book server. You came via %s\n", r.URL.String())
	case "/about":
		fmt.Fprintf(w, "All about our library of books. You came via %s\n", r.URL.String())
	case "/book":
		item := r.URL.Query().Get("item")
		book, ok := b[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "book not found for: %q\n", item)
			return
		}
		fmt.Fprintf(w, "The book is %s\n", book)
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

	// Routing
	bk := books{}
	bk ["jrrtolkein"] = "The Silmarilion"
	fmt.Println("server listening on port 8000")
	//http.ListenAndServe(":8000", bk)
}*/
