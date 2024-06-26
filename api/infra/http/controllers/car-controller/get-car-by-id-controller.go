package carcontroller

import (
	"net/http"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// FindCarByIdController handles the HTTP GET request find a car.
// @Summary				Find car
// @Description		Find a car with the provided id
// @ID				get-car
// @Tags			Car
// @Accept			json
// @Produce			json
// @Security 		BearerAuth
// @Param        	id   			path   		string  true  	  "Car ID"
// @Success	    	201   			{object} 	cardtos.CarOutputDTO "car"
// @Failure			422				{array}		validation_errors.HTTPError
// @Router			/api/v1/cars/{id} [get]
func FindCarByIdController(context *gin.Context, findByIdUseCase *usecases.GetCarByIdUseCase) {

	id := context.Param("id")

	car, err := findByIdUseCase.Execute(id)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, car)
}
