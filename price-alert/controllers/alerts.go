package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zrayyes/PriceAlert/price-alert/models"
)

// GET /alerts
// Get all alerts
func FindAlerts(c *gin.Context) {
	var alerts []models.Alert

	if err := models.GetAlerts(&alerts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": alerts})
}

// GET /alerts/:id
// Find an alert
func FindAlert(c *gin.Context) {
	var alert models.Alert

	if err := models.FindAlert(&alert, c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": alert})
}

// POST /alerts
// Create a new alert
func CreateAlert(c *gin.Context) {
	var input models.CreateAlertInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create alert
	alert := models.Alert{Email: input.Email, Coin: input.Coin, Currency: input.Currency, PriceMin: input.PriceMin, PriceMax: input.PriceMax}
	if err := models.CreateAlert(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": alert})
}

// PATCH /alerts/:id
// Update an alert
func UpdateAlert(c *gin.Context) {
	var alert models.Alert
	if err := models.FindAlert(&alert, c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not found!"})
		return
	}

	// Validate input
	var input models.UpdateAlertInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateAlert(&alert, input); err != nil {
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
	if err := models.FindAlert(&alert, c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not found!"})
		return
	}

	if err := models.DeleteAlert(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Delete!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
