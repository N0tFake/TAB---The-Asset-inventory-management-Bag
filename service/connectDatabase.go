package service

import (
	"fmt"
	"tab/initializers"
	models "tab/models/Patrimony"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(config *initializers.Config) {

	fmt.Println(config.DBHost)
	fmt.Println(config.DBName)
	fmt.Println(config.DBPort)
	fmt.Println(config.DBUserPassword)
	fmt.Println(config.DBUser)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName)

	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error connecting to DB")
	}

	err = db.AutoMigrate(&models.Patrimony{})
	if err != nil {
		panic("Error connecting to DB")
	}
	DB = db
}
