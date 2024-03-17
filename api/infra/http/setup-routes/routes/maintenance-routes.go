package routes

import (
	maintenancefactory "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/factories/maintenance"
	"github.com/gin-gonic/gin"
)

func SetupMaintenanceRoutes(router *gin.Engine, authGroup *gin.RouterGroup) {

	router.POST("/maintenance/:carID/maintenance/create", maintenancefactory.PostMaintenanceFactoryController)
	router.PATCH("/maintenance/:maintenanceID", maintenancefactory.PatchMaintenanceFactoryController)
	router.DELETE("/maintenance/:maintenanceID", maintenancefactory.DeleteMaintenanceFactoryController)
	router.GET("/maintenance/maintenances", maintenancefactory.ListMaintenanceFactoryController)
	router.GET("/maintenance/:carID/maintenances", maintenancefactory.GetMaintenancesByCarIDFactoryController)
	router.GET("/maintenance/latest/:carID", maintenancefactory.GetLatestMaintenanceByCarIDFactoryController)
	router.GET("/maintenance/scheduled", maintenancefactory.GetScheduledMaintenancesFactoryController)
	router.GET("/maintenance/by/:maintenance_status", maintenancefactory.GetMaintenanceByStatusFactoryController)
	router.GET("/maintenance/maintenance/by-date-range", maintenancefactory.GetMaintenancesByDateRangeFactoryController)
}
