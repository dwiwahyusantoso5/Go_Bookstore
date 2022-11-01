package routes

import (
	"go-bookstore/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(r *mux.Router) {
	r.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	r.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
