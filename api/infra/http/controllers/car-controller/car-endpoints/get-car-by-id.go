package carendpoints

import (
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// FindCarByIdController handles the HTTP GET request find a car.
// @Summary				Find car
// @Description		Find a car with the provided id
// @ID						get-car
// @Tags					Car
// @Accept				json
// @Produce				json
// @Param        	id   				path   		 string  true  		"Car ID"
// @Success	    	201   			{object} 	dtos.CarOutputDTO "car"
// @Failure				422					{array}		validation_errors.HTTPError
// @Router				/api/v1/cars/{id} [get]
func FindCarByIdController(context *gin.Context,carRepository repositories.CarRepository, specificationRepository repositories.SpecificationRepository) {

	id := context.Param("id")

	findByIdUseCase := *usecases.NewFindCarByIdUseCase(carRepository, specificationRepository)

	car, err := findByIdUseCase.Execute(id)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, car)
}