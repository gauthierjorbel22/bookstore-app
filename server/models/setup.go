package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB 

func ConnectDatabase() {

	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{}) //automigrate the book table 
	database.AutoMigrate(&User{}) //automigrate the user table 

	DB = database

}
