package routes

import (
	"github.com/gorilla/mux"
	"github.com/sifatulrabbi/bookstore/src/controllers"
)

func RegisterBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.RemoveBook).Methods("DELETE")
}
