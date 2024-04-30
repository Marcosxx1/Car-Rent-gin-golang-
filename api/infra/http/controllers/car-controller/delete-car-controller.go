package carcontroller

import (
	"net/http"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/gin-gonic/gin"
)

// @Summary Delete a car
// @Description Delete a car with the provided ID.
// @ID delete-car
// @Tags Car
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Car ID to be deleted"
// @Success 200 "Car deleted successfully"
// @Failure 400 {object} validation_errors.HTTPErrorCar "Error details"
// @Router /api/v1/cars/delete/{id} [delete]
func DeleteCarController(context *gin.Context, deleteUseCase *usecases.DeleteCarUseCase) {
	id := context.Param("id")

	err := deleteUseCase.Execute(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, nil)
}
