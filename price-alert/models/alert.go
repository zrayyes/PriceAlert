package models

import (
	_ "github.com/jinzhu/gorm"
)

type Alert struct {
	ID       uint    `json:"id" gorm:"primary_key"`
	Email    string  `json:"email"`
	Coin     string  `json:"coin"`
	Currency string  `json:"currency"`
	PriceMax float64 `json:"price_max"`
	PriceMin float64 `json:"price_min"`
	Active   *bool   `json:"active" gorm:"default:true"`
}
