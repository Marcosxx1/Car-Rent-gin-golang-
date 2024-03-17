package maintenanceendpoints

import (
	"errors"
	"net/http"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// RegisterGetMaintenancesByDateRangeController handles the HTTP GET request to retrieve maintenance records within a date range.
// @Summary				Get maintenance records by date range
// @Description			Get maintenance records within the specified date range
// @ID					get-maintenances-by-date-range
// @Tags				Maintenance
// @Accept				json
// @Produce				json
// @Param				startDate			query		string	true	"Start date of the range (format: '2006-01-02')"
// @Param				endDate				query		string	true	"End date of the range (format: '2006-01-02')"
// @Success	    		200   				{array} 	maintenancedtos.MaintenanceOutputDTO "Successfully retrieved maintenance records"
// @Failure				422					{array}		validation_errors.HTTPErrorCar
// @Router				/api/v1/maintenance/maintenance/by-date-range [get]
func GetMaintenancesByDateRangeController(context *gin.Context, getMaintenancesByDateRangeUseCase *usecases.GetMaintenancesByDateRangeUseCase) {
	startDate := context.Query("startDate")
	endDate := context.Query("endDate")

	if startDate == "" || endDate == "" {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, errors.New("start date and end date are required"))
		return
	}

	maintenances, err := getMaintenancesByDateRangeUseCase.Execute(startDate, endDate)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, maintenances)
}
