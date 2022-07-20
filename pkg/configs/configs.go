package configs

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:Savage0.5@tcp(0.0.0.0:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Failed to connect to the database\n")
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
