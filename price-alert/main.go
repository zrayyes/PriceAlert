package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zrayyes/PriceAlert/price-alert/controllers"
	"github.com/zrayyes/PriceAlert/price-alert/models"
)

func getPort() string {
	if port, ok := os.LookupEnv("PORT"); ok {
		return port
	}
	return "8080"
}

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
	r.Run(fmt.Sprintf(":%s", getPort()))
}
