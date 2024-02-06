package maintenanceendpoints

import (
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/maintenance-dtos.go"
	"github.com/gin-gonic/gin"
)

func RegisterMaintenanceController(context *gin.Context, carRepository repositories.CarRepository, maintenanceRepository repositories.MaintenanceRepository) {
	var request dtos.MaintenanceInputDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	carID := context.Param("carID")

	postMaintenanceUseCase := usecases.NewPostMaintenanceUseCase(context, carRepository, maintenanceRepository)

	createdMaintenance, err := postMaintenanceUseCase.Execute(carID, request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, createdMaintenance)
}
