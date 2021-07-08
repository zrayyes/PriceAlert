package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/zrayyes/PriceAlert/price-alert/helpers"
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

// Populate the database with dummy data
func PopulateDataBase() {
	count := int64(0)
	if err := DB.Model(&Alert{}).Count(&count).Error; err != nil {
		fmt.Println("PopulateDataBase: ", err.Error())
		return
	}
	if count == 0 {
		fmt.Println("DB is empty, populating...")
		DB.Create(&Alert{Email: "Blo558@gmail.com", Coin: "BTC", Currency: "USD", PriceMin: 35650.20, PriceMax: 35651.20})
		DB.Create(&Alert{Email: "xdy123@yahoo.com", Coin: "ETC", Currency: "EUR", PriceMin: 2336.27, PriceMax: 2340.0})
		DB.Create(&Alert{Email: "MioSHA@hotmail.com", Coin: "ADA", Currency: "USD", PriceMin: 1.440, PriceMax: 1.640})
	}
}
