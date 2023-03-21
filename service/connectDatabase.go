package service

import (
	"tab/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Error connecting to DB")
	}

	err = db.AutoMigrate(&models.Patrimony{})
	if err != nil {
		panic("Error connecting to DB")
	}
	DB = db
}
