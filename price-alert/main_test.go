package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/zrayyes/PriceAlert/price-alert/models"
)

func SetUpTestDB() {
	db, _ := gorm.Open("sqlite3", "file::memory:")
	models.DB = db
	models.SetupDatabase()
	models.PopulateDataBase()
}

// GET /alerts
func TestFindAlertsRoute(t *testing.T) {
	SetUpTestDB()
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/alerts", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "{\"id\":1,\"email\":\"Blo558@gmail.com\",\"coin\":\"BTC\",\"currency\":\"USD\",\"price_min\":35650.2,\"price_max\":35651.2,\"active\":true}")
}

// POST /alerts
func TestCreateAlertRoute(t *testing.T) {
	SetUpTestDB()
	router := setupRouter()
	w := httptest.NewRecorder()

	// Prepare post body
	postBody := map[string]interface{}{
		"email":     "LoveETC",
		"coin":      "ETC",
		"currency":  "EUR",
		"price_max": 45.55,
		"price_min": 47,
	}
	body, _ := json.Marshal(postBody)
	req, _ := http.NewRequest("POST", "/alerts", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// Confirm that a new id was created
	assert.Contains(t, w.Body.String(), "{\"data\":{\"id\":")
	// Confirm that the alert was set to active
	assert.Contains(t, w.Body.String(), "\"active\":true")
}

// GET /alerts/1
func TestFindAlertRoute(t *testing.T) {
	SetUpTestDB()
	router := setupRouter()

	// Make request to an existing alert
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/alerts/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Make request to a non existent alert
	w = httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/alerts/5", nil)
	router.ServeHTTP(w, req2)
	assert.Equal(t, 400, w.Code)
}

// PATCH /alerts/1
func TestUpdateAlertRoute(t *testing.T) {
	SetUpTestDB()
	router := setupRouter()

	// Change email and set to disabled
	w := httptest.NewRecorder()
	patchBody := map[string]interface{}{
		"email":  "ThisIsAFakeEMail",
		"active": false,
	}
	body, _ := json.Marshal(patchBody)
	req, _ := http.NewRequest("PATCH", "/alerts/1", bytes.NewReader(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "ThisIsAFakeEMail")
	assert.Contains(t, w.Body.String(), "\"active\":false")

	// Make request to same alert
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/alerts/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.NotContains(t, w.Body.String(), "Blo558@gmail.com")
	assert.Contains(t, w.Body.String(), "ThisIsAFakeEMail")
	assert.Contains(t, w.Body.String(), "\"active\":false")
}

// DELETE /alerts/1
func TestDeleteAlertRoute(t *testing.T) {
	SetUpTestDB()
	router := setupRouter()

	// Delete an existing alert
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/alerts/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Delete same alert
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/alerts/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	// Check that alert does not exist
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/alerts", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.NotContains(t, w.Body.String(), "\"id\":1,")
}
