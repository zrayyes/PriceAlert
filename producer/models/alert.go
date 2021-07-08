package models

// Price alert structure Database/JSON mapping
type Alert struct {
	ID       uint    `json:"id" gorm:"primary_key"`
	Email    string  `json:"email"`
	Coin     string  `json:"coin"`
	Currency string  `json:"currency"`
	PriceMax float64 `json:"price_max"`
	PriceMin float64 `json:"price_min"`
	Active   *bool   `json:"active" gorm:"default:true"`
}

// Struct that represents a Kafka event carrying an Alert
type AlertEvent struct {
	Email    string  `json:"email"`
	Coin     string  `json:"coin"`
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
}
