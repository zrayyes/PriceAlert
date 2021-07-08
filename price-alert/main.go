package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zrayyes/PriceAlert/price-alert/controllers"
	"github.com/zrayyes/PriceAlert/price-alert/helpers"
	"github.com/zrayyes/PriceAlert/price-alert/models"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/alerts", controllers.FindAlerts)
	r.POST("/alerts", controllers.CreateAlert)
	r.GET("/alerts/:id", controllers.FindAlert)
	r.PATCH("/alerts/:id", controllers.UpdateAlert)
	r.DELETE("/alerts/:id", controllers.DeleteAlert)
	return r
}

func main() {
	models.ConnectDataBase()
	models.SetupDatabase()

	r := setupRouter()
	r.Run(fmt.Sprintf(":%s", helpers.GetEnv("PORT", "8080")))
}
