package models

// Struct that represents a Kafka event carrying an Alert
type AlertEvent struct {
	Email    string  `json:"email"`
	Coin     string  `json:"coin"`
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
}
