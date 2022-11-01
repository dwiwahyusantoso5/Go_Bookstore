package main

import (
	"go-bookstore/database"
	"go-bookstore/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.StartDB()
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
