package maintenancefactory

import (
	maintenanceusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	maintenancecontroller "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller"
	"github.com/gin-gonic/gin"
)

func GetLatestMaintenanceByCarIDFactoryController(context *gin.Context) {
	maintenanceRepository := database.NewPgMaintenanceRepository()

	getLatestMaintenanceByCarUseCase := maintenanceusecases.NewGetLatestMaintenanceByCarIDUseCase(maintenanceRepository)

	maintenancecontroller.GetLatestMaintenanceByCarIDController(context, getLatestMaintenanceByCarUseCase)

}
