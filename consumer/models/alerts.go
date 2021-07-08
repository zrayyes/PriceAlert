package models

import (
	_ "github.com/jinzhu/gorm"
)

type AlertEvent struct {
	Email    string  `json:"email"`
	Coin     string  `json:"coin"`
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
}
