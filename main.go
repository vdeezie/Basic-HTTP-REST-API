package main

import (
	"fmt"
	"net/http"	
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Go HTTP")
	})

	http.ListenAndServe(":8000", nil)
}
