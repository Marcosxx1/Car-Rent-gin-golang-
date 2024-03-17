package maintenanceendpoints

import (
	"net/http"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// GetMaintenancesByCarIDController handles the HTTP GET request to retrieve maintenances by carID.
// @Summary             Get maintenances by carID
// @Description         Get a list of maintenances associated with a specific carID
// @ID                  get-maintenances-by-carID
// @Tags                Maintenance
// @Accept              json
// @Produce             json
// @Param               carID       path    string  true  "CarID"
// @Param               page        query   int     false "Page number (default 1)"
// @Param               pageSize    query   int     false "Number of items per page (default 10)"
// @Success             200         {array} maintenancedtos.MaintenanceOutputDTO "Successfully retrieved maintenances"
// @Failure             422         {array} validation_errors.HTTPError "Validation errors"
// @Router              /api/v1/maintenance/{carID}/maintenances [get]
func GetMaintenancesByCarIDController(context *gin.Context, getMaintenancesByCarIDUseCase *usecases.GetMaintenancesByCarIDUseCase) {
	carID := context.Param("carID")

	maintenances, err := getMaintenancesByCarIDUseCase.Execute(carID)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, maintenances)
}
