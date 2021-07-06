package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zrayyes/PriceAlert/price-alert/src/controllers"
	"github.com/zrayyes/PriceAlert/price-alert/src/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()
	models.SetupDatabase()
	models.PopulateDataBase()

	r.GET("/alerts", controllers.FindAlerts)
	r.POST("/alerts", controllers.CreateAlert)
	r.GET("/alerts/:id", controllers.FindAlert)
	r.PATCH("/alerts/:id", controllers.UpdateAlert)
	r.DELETE("/alerts/:id", controllers.DeleteAlert)
	r.Run()
}
