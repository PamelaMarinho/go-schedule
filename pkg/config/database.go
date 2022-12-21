package config

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Connection() {
	database, err := gorm.Open("mysql", "root:124086@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}
