package routes

import (
	maintenancefactory "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/factories/maintenance"
	"github.com/gin-gonic/gin"
)

func SetupMaintenanceRoutes(authGroup *gin.RouterGroup) {

	authGroup.POST("/maintenance/:carID/maintenance/create", maintenancefactory.PostMaintenanceFactoryController)
	authGroup.PATCH("/maintenance/:maintenanceID", maintenancefactory.PatchMaintenanceFactoryController)
	authGroup.DELETE("/maintenance/:maintenanceID", maintenancefactory.DeleteMaintenanceFactoryController)
	authGroup.GET("/maintenance/maintenances", maintenancefactory.ListMaintenanceFactoryController)
	authGroup.GET("/maintenance/:carID/maintenances", maintenancefactory.GetMaintenancesByCarIDFactoryController)
	authGroup.GET("/maintenance/latest/:carID", maintenancefactory.GetLatestMaintenanceByCarIDFactoryController)
	authGroup.GET("/maintenance/scheduled", maintenancefactory.GetScheduledMaintenancesFactoryController)
	authGroup.GET("/maintenance/by/:maintenance_status", maintenancefactory.GetMaintenanceByStatusFactoryController)
	authGroup.GET("/maintenance/maintenance/by-date-range", maintenancefactory.GetMaintenancesByDateRangeFactoryController)
}
