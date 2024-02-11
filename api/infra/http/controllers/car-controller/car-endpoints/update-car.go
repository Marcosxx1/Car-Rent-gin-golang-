package carendpoints

import (
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller/car-dtos"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// UpdateCarController handles the HTTP PUT request to update an existing car.
// @Summary     Update a car
// @Description Update a car with the provided ID.
// @ID          put-car
// @Tags        Car
// @Accept      json
// @Produce     json
// @Param       id         		path   		string  true  "Car ID"
// @Param				request				body 		  dtos.CarInputDTO	true "Car information to be updated"
// @Success	    201   				{object} 	dtos.CarOutputDTO "Successfully updated car"
// @Failure			400       		{object} 	validation_errors.HTTPErrorCar
// @Router			/api/v1/cars/update/:id [put]
func UpdateCarController(context *gin.Context, carRepository repositories.CarRepository, specificationRepository repositories.SpecificationRepository) {

	var request dtos.CarOutputDTO
	if err := context.ShouldBindJSON(&request); err != nil { 
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	id := context.Param("id")
	if err := context.ShouldBindJSON(&request); err != nil { 
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	findByIdUseCase := *usecases.NewFindCarByIdUseCase(carRepository, specificationRepository)
	updateCarUseCase := *usecases.NewUpdateCarUseCase(carRepository, specificationRepository)

	foundCar, err := findByIdUseCase.Execute(id)
	if foundCar != nil { 
		validation_errors.NewError(context, http.StatusNotFound, err)
		return
	}

	updatedCar, err := updateCarUseCase.Execute(id, &request)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, updatedCar)
	}
}
