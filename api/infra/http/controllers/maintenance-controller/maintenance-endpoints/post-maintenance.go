package maintenanceendpoints

import (
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases"
	maintenancedtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/maintenance-dtos.go"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// RegisterMaintenanceController handles the HTTP POST request to create a new maintenance.
// @Summary				Create a new maintenance
// @Description		Create a new maintenance with the provided information
// @ID						post-car
// @Tags					Maintenance
// @Accept				json
// @Produce				json
// @Param        	carID   		path   		 string  true  "CarID"
// @Param					request			body 		  maintenancedtos.MaintenanceInputDTO	true "Maintenance information to be created"
// @Success	    	201   			{object} 	maintenancedtos.MaintenanceOutputDTO "Successfully created maintenance"
// @Failure				422					{array}		validation_errors.HTTPErrorCar
// @Router				/api/v1/cars/{carID}/maintenance/create [post]
func RegisterMaintenanceController(context *gin.Context, carRepository repositories.CarRepository, maintenanceRepository repositories.MaintenanceRepository) {
	var request maintenancedtos.MaintenanceInputDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	carID := context.Param("carID")

	postMaintenanceUseCase := usecases.NewPostMaintenanceUseCase(carRepository, maintenanceRepository)

	createdMaintenance, err := postMaintenanceUseCase.ExecuteConcurrently(carID, request)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, createdMaintenance)
}
