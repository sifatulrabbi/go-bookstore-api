package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sifatulrabbi/bookstore/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/api/v1", router)
	fmt.Printf("Starting the server on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}
