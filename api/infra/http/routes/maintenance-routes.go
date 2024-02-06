package routes

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	maintenanceendpoints "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/maintenance-endpoints"
	"github.com/gin-gonic/gin"
)

func SetupMaintenanceRoutes(router *gin.Engine) {
	carRepository := database.PGCarRepository{}
	maintenanceRepository := database.PGMaintenanceRepository{}

	router.POST("/api/v1/cars/:carID/maintenance/create", func(context *gin.Context) {
		maintenanceendpoints.RegisterMaintenanceController(context, &carRepository, &maintenanceRepository)
	})
}
