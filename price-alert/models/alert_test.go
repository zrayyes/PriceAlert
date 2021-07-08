package models

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

// Start a db in memory and populate it with test data
func SetUpTestDB() {
	db, _ := gorm.Open("sqlite3", "file::memory:")
	DB = db
	SetupDatabase()
	PopulateDataBase()
}

func TestFindAlert(t *testing.T) {
	SetUpTestDB()

	// Confirm that an alert exists
	var alert1 Alert
	FindAlert(&alert1, "1")
	assert.Equal(t, alert1.Coin, "BTC")
	assert.Equal(t, alert1.PriceMin, 35650.20)

	// Confirm that an alert does not exist
	var alert2 Alert
	FindAlert(&alert2, "5")
	assert.Equal(t, alert2.Email, "")
	assert.Equal(t, alert2.PriceMax, float64(0))
}

func TestGetAlerts(t *testing.T) {
	SetUpTestDB()

	// Get a list of all alerts
	var alerts []Alert
	GetAlerts(&alerts)
	assert.Len(t, alerts, 3)
	assert.NotEqual(t, alerts[1].Coin, "")
	assert.NotEqual(t, alerts[1].PriceMin, float64(0))
}

func TestCreateAlert(t *testing.T) {
	SetUpTestDB()

	// Confirm the length before creation
	var alertsBefore []Alert
	GetAlerts(&alertsBefore)
	assert.Len(t, alertsBefore, 3)

	// Create a new alert
	alert := Alert{Email: "justForTesting", Coin: "XYZ", Currency: "ABC", PriceMin: 55, PriceMax: 56}
	CreateAlert(&alert)

	// Check that an alert was added
	var alertsAfter []Alert
	GetAlerts(&alertsAfter)
	assert.Len(t, alertsAfter, 4)
	assert.Contains(t, alertsAfter, alert)
}

func TestUpdateAlert(t *testing.T) {
	SetUpTestDB()

	// Find an existing alert
	var alert1 Alert
	fakeEmail := "justForTestingAgain"
	FindAlert(&alert1, "1")
	assert.NotEqual(t, alert1.Email, fakeEmail)

	// Update the email for the existing alert
	input := UpdateAlertInput{Email: fakeEmail}
	UpdateAlert(&alert1, input)

	// Confirm that the email was changed
	var alert2 Alert
	FindAlert(&alert2, "1")
	assert.Equal(t, alert2.Email, fakeEmail)
}

func TestDeleteAlert(t *testing.T) {
	SetUpTestDB()
	// Create a new alert
	alert := Alert{Email: "justForTesting", Coin: "XYZ", Currency: "ABC", PriceMin: 55, PriceMax: 56}
	CreateAlert(&alert)

	// Confirm that the alert was created
	var alertsBefore []Alert
	GetAlerts(&alertsBefore)
	assert.Len(t, alertsBefore, 4)
	assert.Contains(t, alertsBefore, alert)

	// Delete the alert
	DeleteAlert(&alert)

	// Confirm that the alert was deleted
	var alertsAfter []Alert
	GetAlerts(&alertsAfter)
	assert.Len(t, alertsAfter, 3)
	assert.NotContains(t, alertsAfter, alert)
}
