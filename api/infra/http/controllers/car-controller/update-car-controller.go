package carcontroller

import (
	"net/http"

	cardtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/car"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
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
// @Security ApiKeyAuth
// @Param       id         		    path   		string  true  "id"
// @Param		request				body 		cardtos.CarInputDTO	true "Car information to be updated"
// @Success	    201   				{object} 	cardtos.CarOutputDTO "Successfully updated car"
// @Failure		400       		    {object} 	validation_errors.HTTPErrorCar
// @Router		/api/v1/cars/update/{id} [put]
func UpdateCarController(context *gin.Context, updateCarUseCase *usecases.PutCarUseCase) {

	var request *cardtos.CarInputDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	id := context.Param("id")

	updatedCar, err := updateCarUseCase.Execute(id, request)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, updatedCar)
	}
}
