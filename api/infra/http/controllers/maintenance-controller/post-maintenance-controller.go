package maintenancecontroller

import (
	"net/http"

	maintenancedtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/maintenance"
	maintenanceusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// PostMaintenanceController handles the HTTP POST request to create a new maintenance.
// @Summary				Create a new maintenance
// @Description			Create a new maintenance with the provided information
// @ID					post-maintenance
// @Tags				Maintenance
// @Accept				json
// @Produce				json
// @Param        		carID   			path   		string  true  "CarID"
// @Param				request				body 		maintenancedtos.MaintenanceInputDTO	true "Maintenance information to be created"
// @Success	    		201   				{object} 	maintenancedtos.MaintenanceOutputDTO "Successfully created maintenance"
// @Failure				422					{array}		validation_errors.HTTPErrorCar
// @Router				/api/v1/maintenance/{carID}/maintenance/create [post]
func PostMaintenanceController(context *gin.Context, postMaintenanceUseCase *maintenanceusecases.PostMaintenanceUseCase) {

	var request maintenancedtos.MaintenanceInputDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	carID := context.Param("carID")

	createdMaintenance, err := postMaintenanceUseCase.Execute(carID, request)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, createdMaintenance)
}
