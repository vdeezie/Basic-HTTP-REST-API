package main

import (
	"log"
	"net/http"

	"BookServiceRESTAPI/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)
func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
