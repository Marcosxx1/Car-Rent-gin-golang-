package maintenancecontroller

import (
	"net/http"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// GetLatestMaintenanceByCar handles the HTTP GET request to retrieve the latest maintenance by carID.
// @Summary             Get latest maintenance by carID
// @Description         Get latest maintenance associated with a specific carID
// @ID                  get-latest-maintenances-by-carID
// @Tags                Maintenance
// @Accept              json
// @Produce             json
// @Param               carID       path    string  true  "CarID"
// @Success             200         {array} maintenancedtos.MaintenanceOutputDTO "Successfully retrieved maintenances"
// @Failure             422         {array} validation_errors.HTTPError "Validation errors"
// @Router              /api/v1/maintenance/latest/{carID} [get]
func GetLatestMaintenanceByCarIDController(context *gin.Context, getLatestMaintenanceByCarUseCase *usecases.GetLatestMaintenanceByCar) {
	carID := context.Param("carID")

	maintenances, err := getLatestMaintenanceByCarUseCase.Execute(carID)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, maintenances)
}
