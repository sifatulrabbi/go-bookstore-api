package routes

import (
	"github.com/gorilla/mux"
	"github.com/sifatulrabbi/bookstore/pkg/controllers"
)

func RegisterBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.RemoveBook).Methods("DELETE")
}
