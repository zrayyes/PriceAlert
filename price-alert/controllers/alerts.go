package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zrayyes/PriceAlert/price-alert/models"
)

type CreateAlertInput struct {
	Email    string  `json:"email" binding:"required"`
	Coin     string  `json:"coin" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
	PriceMin float64 `json:"price_min" binding:"required"`
	PriceMax float64 `json:"price_max" binding:"required"`
}

type UpdateAlertInput struct {
	Email    string  `json:"email"`
	Coin     string  `json:"coin"`
	Currency string  `json:"currency"`
	PriceMin float64 `json:"price_min"`
	PriceMax float64 `json:"price_max"`
	Active   *bool   `json:"active"`
}

// GET /alerts
// Get all alerts
func FindAlerts(c *gin.Context) {
	var alerts []models.Alert

	if err := models.DB.Find(&alerts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": alerts})
}

// GET /alerts/:id
// Find an alert
func FindAlert(c *gin.Context) {
	var alert models.Alert

	if err := models.DB.Where("id = ?", c.Param("id")).First(&alert).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": alert})
}

// POST /alerts
// Create a new alert
func CreateAlert(c *gin.Context) {
	var input CreateAlertInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create alert
	alert := models.Alert{Email: input.Email, Coin: input.Coin, Currency: input.Currency, PriceMin: input.PriceMin, PriceMax: input.PriceMax}
	if err := models.DB.Create(&alert).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": alert})
}

// PATCH /alerts/:id
// Update an alert
func UpdateAlert(c *gin.Context) {
	var alert models.Alert
	if err := models.DB.Where("id = ?", c.Param("id")).First(&alert).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not found!"})
		return
	}

	// Validate input
	var input UpdateAlertInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Model(&alert).Updates(input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Update!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": alert})
}

// DELETE /alerts/:id
// Delete an alert
func DeleteAlert(c *gin.Context) {
	// Get model if exist
	var alert models.Alert
	if err := models.DB.Where("id = ?", c.Param("id")).First(&alert).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not found!"})
		return
	}

	if err := models.DB.Delete(&alert).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Delete!"})
		return
	}
	models.DB.Delete(&alert)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
