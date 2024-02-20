package carendpoints

import (
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller/car-dtos"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// RegisterCarController handles the HTTP POST request to create a new car.
// @Summary				Create a new car
// @Description		Create a new car with the provided information
// @ID						post-car
// @Tags					Car
// @Accept				json
// @Produce				json
// @Param					request			body 		  dtos.CarInputDTO	true "Car information to be created"
// @Success	    	201   			{object} 	dtos.CarOutputDTO "Successfully created car"
// @Failure				422					{array}		validation_errors.HTTPErrorCar
// @Router				/api/v1/cars/create [post]
func RegisterCarController(context *gin.Context, carRepository repositories.CarRepository, specificationRepository repositories.SpecificationRepository) {

	var request *dtos.CarInputDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	} 

	postCarUseCase := usecases.NewPostCarUseCase(carRepository, specificationRepository)

	createdCar, err := postCarUseCase.ExecuteConcurrently(request)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, createdCar)

}
