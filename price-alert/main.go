package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zrayyes/PriceAlert/price-alert/controllers"
	"github.com/zrayyes/PriceAlert/price-alert/helpers"
	"github.com/zrayyes/PriceAlert/price-alert/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()
	models.SetupDatabase()
	if os.Getenv("GIN_MODE") != gin.ReleaseMode {
		models.PopulateDataBase()
	}

	r.GET("/alerts", controllers.FindAlerts)
	r.POST("/alerts", controllers.CreateAlert)
	r.GET("/alerts/:id", controllers.FindAlert)
	r.PATCH("/alerts/:id", controllers.UpdateAlert)
	r.DELETE("/alerts/:id", controllers.DeleteAlert)
	r.Run(fmt.Sprintf(":%s", helpers.GetEnv("PORT", "8080")))
}
