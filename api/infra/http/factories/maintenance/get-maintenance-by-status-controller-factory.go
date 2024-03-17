package maintenancefactory

import (
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	maintenanceendpoints "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/maintenance-endpoints"
	"github.com/gin-gonic/gin"
)

func GetMaintenanceByStatusFactoryController(context *gin.Context) {
	maintenanceRepository := database.NewPgMaintenanceRepository()

	getMaintenanceByStatus := usecases.NewGetMaintenanceByStatusUseCase(maintenanceRepository)

	maintenanceendpoints.GetMaintenanceByStatusController(context, getMaintenanceByStatus)

}
