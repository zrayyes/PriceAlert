package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}

func SetupDatabase() {
	// Clear any existing table
	DB.DropTable(&Alert{})
	// Migrate schema
	DB.AutoMigrate(&Alert{})
}

func PopulateDataBase() {
	DB.Create(&Alert{Email: "Blo558@gmail.com", Coin: "BTC", Price: 35650.20})
	DB.Create(&Alert{Email: "xdy123@yahoo.com", Coin: "ETC", Price: 2336.27})
	DB.Create(&Alert{Email: "MioSHA@hotmail.com", Coin: "ADA", Price: 1.440})
}