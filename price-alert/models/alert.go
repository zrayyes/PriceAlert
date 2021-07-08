package models

import (
	_ "github.com/jinzhu/gorm"
)

type CreateAlertInput struct {
	Email    string  `json:"email" binding:"required"`
	Coin     string  `json:"coin" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
	PriceMin float64 `json:"price_min" binding:"required"`
	PriceMax float64 `json:"price_max" binding:"required"`
}

// I really hope there's a better way to do this
type UpdateAlertInput struct {
	Email    string  `json:"email" binding:"required"`
	Coin     string  `json:"coin" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
	PriceMin float64 `json:"price_min" binding:"required"`
	PriceMax float64 `json:"price_max" binding:"required"`
	Active   *bool   `json:"active" gorm:"default:true"`
}

type Alert struct {
	Email    string  `json:"email" binding:"required"`
	Coin     string  `json:"coin" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
	PriceMin float64 `json:"price_min" binding:"required"`
	PriceMax float64 `json:"price_max" binding:"required"`
	Active   *bool   `json:"active" gorm:"default:true"`
	ID       uint    `json:"id" gorm:"primary_key"`
}

func GetAlerts(alerts *[]Alert) error {
	return DB.Find(&alerts).Error
}

func FindAlert(alert *Alert, id string) error {
	return DB.Where("id = ?", id).First(&alert).Error
}

func CreateAlert(alert *Alert) error {
	return DB.Create(alert).Error
}

func UpdateAlert(alert *Alert, update UpdateAlertInput) error {
	return DB.Model(&alert).Updates(update).Error
}

func DeleteAlert(alert *Alert) error {
	return DB.Delete(&alert).Error
}
