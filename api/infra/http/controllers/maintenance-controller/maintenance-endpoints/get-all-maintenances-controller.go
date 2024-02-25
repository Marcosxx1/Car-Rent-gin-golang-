package maintenanceendpoints

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// ListMaintenanceController handles the HTTP GET request to retrieve a list of maintenance records.
// @Summary             Retrieve a list of maintenance records
// @Description         Retrieve a list of maintenance records for a specific car
// @ID                  get-maintenance-list
// @Tags                Maintenance
// @Accept              json
// @Produce             json
// @Param page query int false "Page number (default is 1)"
// @Param pageSize query int false "Number of items per page (default is 10)"
// @Success             200          {array}    maintenancedtos.MaintenanceOutputDTO "List of maintenance records"
// @Failure             404          {object}   validation_errors.HTTPErrorCar
// @Router              /api/v1/maintenance/maintenances [get]
func ListMaintenanceController(context *gin.Context, maintenanceRepository repositories.MaintenanceRepository) {

	pageStr := context.Query("page")
	pageSizeStr := context.Query("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		validation_errors.NewError(context, http.StatusBadRequest, errors.New("invalid 'page' value"))
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		validation_errors.NewError(context, http.StatusBadRequest, errors.New("invalid 'pageSize' value"))
		return
	}

	listMaintenanceUseCase := usecases.NewListMaintenanceUseCase(maintenanceRepository)

	maintenanceList, err := listMaintenanceUseCase.Execute(page, pageSize)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, maintenanceList)
}
