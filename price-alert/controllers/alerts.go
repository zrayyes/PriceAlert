package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zrayyes/PriceAlert/price-alert/models"
)

type CreateAlertInput struct {
	Email string  `json:"email" binding:"required"`
	Coin  string  `json:"coin" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type UpdateAlertInput struct {
	Email  string  `json:"email"`
	Coin   string  `json:"coin"`
	Price  float64 `json:"price"`
	Active *bool   `json:"active"`
}

// GET /alerts
// Get all alerts
func FindAlerts(c *gin.Context) {
	var alerts []models.Alert
	models.DB.Find(&alerts)

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
	alert := models.Alert{Email: input.Email, Coin: input.Coin, Price: input.Price}
	models.DB.Create(&alert)

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

	models.DB.Model(&alert).Updates(input)

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

	models.DB.Delete(&alert)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
