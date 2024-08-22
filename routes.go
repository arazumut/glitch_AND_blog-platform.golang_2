package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanı bağlantısı hatalı!")
	}
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}
