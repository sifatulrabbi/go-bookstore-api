package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sifatulrabbi/bookstore/src/configs"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	configs.Connect()
	db = configs.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	bookDb := db.Where("ID=?", id).Find(&book)
	return &book, bookDb
}

func RemoveBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return book
}
