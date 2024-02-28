package maintenanceendpoints

import (
	"net/http"
	"strconv"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// GetMaintenanceByStatusController handles the HTTP GET request to retrieve maintenances by its status.
// @Summary             Get maintenances by its status
// @Description         Get a list of maintenances by its status
// @ID                  get-maintenances-by-status
// @Tags                Maintenance
// @Accept              json
// @Produce             json
// @Param               carID       path    boolean  true  "CarID"
// @Success             200         {array} maintenancedtos.MaintenanceOutputDTO "Successfully retrieved maintenances"
// @Failure             422         {array} validation_errors.HTTPError "Validation errors"
// @Router              /api/v1/maintenance/:status/maintenances [get]
func GetMaintenanceByStatusController(context *gin.Context, maintenanceRepository repositories.MaintenanceRepository) {
	maintenanceStatusStr := context.Query("maintenance_status")

	maintenanceStatus, err := strconv.ParseBool(maintenanceStatusStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maintenance_status value"})
		return
	}
	getMaintenanceByStatusUseCase := usecases.NewNewGetMaintenanceByStatusUseCase(maintenanceRepository)

	maintenances, err := getMaintenanceByStatusUseCase.Execute(maintenanceStatus)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, maintenances)
}
