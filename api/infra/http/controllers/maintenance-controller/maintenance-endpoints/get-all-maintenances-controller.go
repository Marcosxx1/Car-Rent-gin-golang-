package maintenanceendpoints

import (
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases"
	"github.com/gin-gonic/gin"
)

// ListMaintenanceController handles the HTTP GET request to retrieve a list of maintenance records.
// @Summary             Retrieve a list of maintenance records
// @Description         Retrieve a list of maintenance records for a specific car
// @ID                  get-maintenance-list
// @Tags                Maintenance
// @Accept              json
// @Produce             json
// @Success             200          {array}    maintenancedtos.MaintenanceOutputDTO "List of maintenance records"
// @Failure             404          {object}   validation_errors.HTTPErrorCar
// @Router              /api/v1/maintenance/maintenances [get]
func ListMaintenanceController(context *gin.Context, maintenanceRepository repositories.MaintenanceRepository) {

	listMaintenanceUseCase := usecases.NewListMaintenanceUseCase(maintenanceRepository)

	maintenanceList, err := listMaintenanceUseCase.Execute()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, maintenanceList)
}
