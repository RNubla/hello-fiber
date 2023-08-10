package database

import (
	"fmt"

	"github.com/RNubla/hello-fiber/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("hello-fiber.sqlite"), &gorm.Config{})
	// db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.User{})
	fmt.Println("Database Migrated")

}
