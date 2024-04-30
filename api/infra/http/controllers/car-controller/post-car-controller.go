package carcontroller

import (
	"net/http"

	cardtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/car"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// RegisterCarController handles the HTTP POST request to create a new car.
// @Summary				Create a new car
// @Description			Create a new car with the provided information
// @ID					post-car
// @Tags				Car
// @Accept				json
// @Produce				json
// @Security ApiKeyAuth
// @Param				request				body 		cardtos.CarInputDTO	true "Car information to be created"
// @Success	    		201   				{object} 	cardtos.CarOutputDTO "Successfully created car"
// @Failure				422					{array}		validation_errors.HTTPErrorCar
// @Router				/api/v1/cars/create [post]
func RegisterCarController(context *gin.Context, postCarUseCase *usecases.PostCarUseCase) {

	var request *cardtos.CarInputDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	createdCar, err := postCarUseCase.ExecuteConcurrently(request)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, createdCar)

}
