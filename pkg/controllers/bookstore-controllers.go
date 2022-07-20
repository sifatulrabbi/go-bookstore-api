package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sifatulrabbi/bookstore/pkg/models"
	"github.com/sifatulrabbi/bookstore/pkg/utils"
)

func CreateBook(wr http.ResponseWriter, req *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(req, CreateBook)
	book := CreateBook.CreateBook()
	utils.SendHTTPResp(wr, http.StatusCreated, book)
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
	UpdateBook := &models.Book{}
	utils.ParseBody(req, UpdateBook)
	params := mux.Vars(req)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		utils.SendHTTPResp(wr, http.StatusBadRequest, "Invalid book id")
		return
	}
	book, db := models.GetBookById(id)
	if UpdateBook.Name != "" {
		book.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		book.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		book.Publication = UpdateBook.Publication
	}
	db.Save(&book)
	utils.SendHTTPResp(wr, http.StatusOK, book)
}

func RemoveBook(wr http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		utils.SendHTTPResp(wr, http.StatusBadRequest, "Book ID not found.")
		return
	}
	book := models.RemoveBook(id)
	utils.SendHTTPResp(wr, http.StatusOK, book)
}
