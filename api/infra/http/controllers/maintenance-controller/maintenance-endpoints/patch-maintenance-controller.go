package maintenanceendpoints

import (
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	maintenanceusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases"
	maintenancedtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/maintenance-controller/maintenance-dtos.go"
	"github.com/gin-gonic/gin"
)

// PatchMaintenanceController handles the HTTP PATCH request to update an existing maintenance.
// @Summary				Update an existing maintenance
// @Description			Update an existing maintenance with the provided information
// @ID					patch-maintenance
// @Tags				Maintenance
// @Accept				json
// @Produce				json
// @Param				maintenanceID		path		string	true	"maintenanceID"
// @Param				request				body 		maintenancedtos.MaintenanceInputDTO	true "Maintenance information to be updated"
// @Success	    		200   				{object} 	maintenancedtos.MaintenanceOutputDTO "Successfully updated maintenance"
// @Router				/api/v1/maintenance/:maintenanceID [patch]
func PatchMaintenanceController(context *gin.Context, carRepository repositories.CarRepository, maintenanceRepository repositories.MaintenanceRepository) {
	var request maintenancedtos.MaintenanceInputDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	maintenanceID := context.Param("maintenanceID")

	patchMaintenanceUseCase := maintenanceusecases.NewPatchMaintenanceUseCase(carRepository, maintenanceRepository)

	updatedMaintenance, err := patchMaintenanceUseCase.Execute(maintenanceID, request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, updatedMaintenance)
}
