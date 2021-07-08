package main

import (
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

func TestAlertsRoute(t *testing.T) {
	SetUpTestDB()
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/alerts", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
