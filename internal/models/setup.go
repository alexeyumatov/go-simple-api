package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("school.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the databse")
	}

	err = database.AutoMigrate(&Student{})
	if err != nil {
		return
	}
	DB = database
}
