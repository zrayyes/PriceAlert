package models

import (
	_ "github.com/jinzhu/gorm"
)

type Alert struct {
	ID     uint    `json:"id" gorm:"primary_key"`
	Email  string  `json:"email"`
	Coin   string  `json:"coin"`
	Price  float64 `json:"price"`
	Active *bool   `json:"active" gorm:"default:true"`
}
