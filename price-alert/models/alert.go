package models

// I really hope there's a better way to do this

// JSON input for alert creation
type CreateAlertInput struct {
	Email    string  `json:"email" binding:"required"`
	Coin     string  `json:"coin" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
	PriceMin float64 `json:"price_min" binding:"required"`
	PriceMax float64 `json:"price_max" binding:"required"`
}

// JSON input for alert update
type UpdateAlertInput struct {
	Email    string  `json:"email"`
	Coin     string  `json:"coin"`
	Currency string  `json:"currency"`
	PriceMin float64 `json:"price_min"`
	PriceMax float64 `json:"price_max"`
	Active   *bool   `json:"active"`
}

// Price alert structure Database/JSON mapping
type Alert struct {
	ID       uint    `json:"id" gorm:"primary_key"`
	Email    string  `json:"email" binding:"required"`
	Coin     string  `json:"coin" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
	PriceMin float64 `json:"price_min" binding:"required"`
	PriceMax float64 `json:"price_max" binding:"required"`
	Active   *bool   `json:"active" gorm:"default:true"`
}

// Return a single alert given an ID
func FindAlert(alert *Alert, id string) error {
	return DB.Where("id = ?", id).First(&alert).Error
}

// Return a list of all alerts
func GetAlerts(alerts *[]Alert) error {
	return DB.Find(&alerts).Error
}

// Create a new alert
func CreateAlert(alert *Alert) error {
	return DB.Create(alert).Error
}

// Update an existing alert
func UpdateAlert(alert *Alert, update UpdateAlertInput) error {
	return DB.Model(&alert).Updates(update).Error
}

// Delete an existing alert
func DeleteAlert(alert *Alert) error {
	return DB.Delete(&alert).Error
}
