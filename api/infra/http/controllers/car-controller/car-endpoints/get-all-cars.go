package carendpoints

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// @Summary List all cars
// @Description Retrieve a list of cars with pagination support.
// @ID list-car
// @Tags Car
// @Accept json
// @Produce json
// @Param page query int false "Page number (default is 1)"
// @Param pageSize query int false "Number of items per page (default is 10)"
// @Success 200 {array} dtos.CarOutputDTO "List of cars"
// @Failure				422					{array}		validation_errors.HTTPError
// @Router /api/v1/cars [get]
func GetAllCarsController(context *gin.Context, carRepository repositories.CarRepository, specificationRepository repositories.SpecificationRepository) {

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

	findAllCarsUseCase := *usecases.NewGetAllCarsUseCase(carRepository, specificationRepository)

	car, err := findAllCarsUseCase.Execute(page, pageSize)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, car)
}
