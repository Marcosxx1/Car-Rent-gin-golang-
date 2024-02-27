package maintenanceendpoints

import (
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/maintenance-use-cases"
	"github.com/gin-gonic/gin"
)

// GetScheduledMaintenancesController handles the HTTP GET request to retrieve scheduled maintenances.
// @Summary Get scheduled maintenances
// @Description Get a list of scheduled maintenances based on the next maintenance due date
// @ID get-scheduled-maintenances
// @Tags Maintenance
// @Accept json
// @Produce json
// @Param next_maintenance_due_date body string true "Next Maintenance Due Date"
// @Success 200 {array} maintenancedtos.MaintenanceOutputDTO "Successfully retrieved scheduled maintenances"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/maintenance/scheduled [get]
func GetScheduledMaintenancesController(context *gin.Context, maintenanceRepository repositories.MaintenanceRepository) {
	var inputDTO struct {
		NextMaintenanceDueDate string `json:"next_maintenance_due_date" binding:"required"`
	}

	if err := context.ShouldBindJSON(&inputDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := usecases.NewGetScheduledMaintenancesUseCase(maintenanceRepository)
	scheduledMaintenances, err := useCase.Execute()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, scheduledMaintenances)
}
