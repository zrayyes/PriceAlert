package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/zrayyes/PriceAlert/producer/helpers"
)

var DB *gorm.DB

// Connect to the sqlite3 database
func ConnectDataBase() {
	DBPATH := helpers.GetEnv("DBPATH", "/home/test.db")
	database, err := gorm.Open("sqlite3", DBPATH)

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}

// Setup the database table(s)
func SetupDatabase() {
	DB.AutoMigrate(&Alert{})
}
