package maintenancefactory

import (
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	maintenanceendpoints "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/maintenance-endpoints"
	"github.com/gin-gonic/gin"
)

func DeleteMaintenanceFactoryController(context *gin.Context) {
	maintenanceRepository := database.NewPgMaintenanceRepository()
	carRepository := database.NewPGCarRepository()

	deleteMaintenanceUseCase := usecases.NewDeleteMaintenanceUseCase(carRepository, maintenanceRepository)

	maintenanceendpoints.DeleteMaintenanceController(context, deleteMaintenanceUseCase)
}
