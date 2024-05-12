package maintenancecontroller

import (
	"fmt"
	"net/http"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/utils"

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
// @Security 			BearerAuth
// @Param               maintenance_status       path    string  true  "maintenance_status" Enums(Scheduled,InProgress,Completed,PendingApproval,Canceled,AwaitingParts,AwaitingPayment,Rescheduled,MaintenanceFailed,AwaitingInspection) "maintenance status"
// @Success             200         {array} maintenancedtos.MaintenanceOutputDTO "Successfully retrieved maintenances"
// @Failure             422         {array} validation_errors.HTTPError "Validation errors"
// @Router              /api/v1/maintenance/by/{maintenance_status} [get]
func GetMaintenanceByStatusController(context *gin.Context, getMaintenanceByStatusUseCase *usecases.GetMaintenanceByStatusUseCase) {
	fmt.Println("maintenance_status coming from Param", context.Param("maintenance_status"))

	// Get the maintenance_status from the path parameters
	rawMaintenanceStatus := context.Param("maintenance_status")

	// Convert the string to the corresponding enum value
	maintenanceStatus, err := utils.ParseMaintenanceStatus(rawMaintenanceStatus)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	maintenances, err := getMaintenanceByStatusUseCase.Execute(maintenanceStatus)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, maintenances)
}
