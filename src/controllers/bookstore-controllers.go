package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sifatulrabbi/bookstore/src/models"
	"github.com/sifatulrabbi/bookstore/src/utils"
)

func CreateBook(wr http.ResponseWriter, req *http.Request) {
}

func GetBooks(wr http.ResponseWriter, req *http.Request) {
	books := models.GetAllBooks()
	jsonData, _ := json.Marshal(books)
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write(jsonData)
}

func GetBookById(wr http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 0, 0)
	book, _ := models.GetBookById(id)
	utils.SendHTTPResp(wr, http.StatusOK, book)
}

func UpdateBook(wr http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	// id, _ := strconv.ParseInt(params["id"], 0, 0)
	fmt.Printf(params["id"], req.Body)
	utils.SendHTTPResp(wr, http.StatusOK, nil)
}

func RemoveBook(wr http.ResponseWriter, req *http.Request) {}
