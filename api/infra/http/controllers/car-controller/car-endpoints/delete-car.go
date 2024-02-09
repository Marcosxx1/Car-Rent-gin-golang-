package carendpoints

import (
	"log"
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// @Summary Delete a car
// @Description Delete a car with the provided ID.
// @ID delete-car
// @Tags Car
// @Accept json
// @Produce json
// @Param id path string true "Car ID to be deleted"
// @Success 200 "Car deleted successfully"
// @Failure 400 {object} validation_errors.HTTPErrorCar "Error details"
// @Router /api/v1/cars/delete/{id} [delete]
func DeleteCarController(context *gin.Context, carRepository repositories.CarRepository) {
 
	id := context.Param("id")

	err := usecases.DeleteCarUseCase(carRepository, id)
	if err != nil {
		log.Println("Error deleting car:", err)
		validation_errors.NewError(context, http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}
