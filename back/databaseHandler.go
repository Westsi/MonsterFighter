package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	_db, err := gorm.Open(sqlite.Open("./monsters.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = _db

	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
}
