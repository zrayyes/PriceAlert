package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "/home/test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}

func SetupDatabase() {
	DB.AutoMigrate(&Alert{})
}

func PopulateDataBase() {
	count := int64(0)
	DB.Model(&Alert{}).Count(&count)
	if count == 0 {
		fmt.Println("DB is empty, populating...")
		DB.Create(&Alert{Email: "Blo558@gmail.com", Coin: "BTC", Currency: "USD", PriceMin: 35650.20, PriceMax: 35651.20})
		DB.Create(&Alert{Email: "xdy123@yahoo.com", Coin: "ETC", Currency: "EUR", PriceMin: 2336.27, PriceMax: 2340.0})
		DB.Create(&Alert{Email: "MioSHA@hotmail.com", Coin: "ADA", Currency: "USD", PriceMin: 1.440, PriceMax: 1.640})
	} else {
		fmt.Println("DB is not empty, skipping population.")
	}
}
