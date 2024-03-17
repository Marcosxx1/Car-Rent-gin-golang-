package maintenancefactory

import (
	maintenanceusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	maintenancecontroller "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller"
	"github.com/gin-gonic/gin"
)

func DeleteMaintenanceFactoryController(context *gin.Context) {
	maintenanceRepository := database.NewPgMaintenanceRepository()
	carRepository := database.NewPGCarRepository()

	deleteMaintenanceUseCase := maintenanceusecases.NewDeleteMaintenanceUseCase(carRepository, maintenanceRepository)

	maintenancecontroller.DeleteMaintenanceController(context, deleteMaintenanceUseCase)
}
