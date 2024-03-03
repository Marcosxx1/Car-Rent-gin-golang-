package carendpoints

import (
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
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
func DeleteCarController(context *gin.Context, carRepository repositories.CarRepository, specificationRepository repositories.SpecificationRepository) {
	id := context.Param("id")

	deleteUseCase := *usecases.NewDeleteCarUseCase(carRepository, specificationRepository)

	deleteUseCase.Execute(id)

	context.JSON(http.StatusOK, "Car deleted")
}
