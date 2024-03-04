package maintenanceendpoints

import (
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	maintenanceusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases"
	"github.com/gin-gonic/gin"
)

// @Summary Delete a maintenance
// @Description Delete a maintenance with the provided ID.
// @ID delete-maintenance
// @Tags Maintenance
// @Accept json
// @Produce json
// @Param id path string true "Maintenance ID to be deleted"
// @Success 200 "Maintenance deleted successfully"
// @Failure 400 {object} validation_errors.HTTPErrorCar "Error details"
// @Router				/api/v1/maintenance/{maintenanceID} [delete]
func DeleteMaintenanceController(context *gin.Context, carRepository repositories.CarRepository, maintenanceRepository repositories.MaintenanceRepository) {
	maintenanceID := context.Param("maintenance_id")

	if maintenanceID == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maintenance ID"})
		return
	}

	deleteMaintenanceUseCase := maintenanceusecases.NewDeleteMaintenanceUseCase(carRepository, maintenanceRepository)

	err := deleteMaintenanceUseCase.Execute(maintenanceID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, nil)
}
