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
// @Sumary				Create a new car
// @Description		Create a new car with the provided information
// @ID						post-car
// @Tags					Car
// @Accept				json
// @Produce				json
// @Param					request			body 		 dtos.CarInputDTO	true "Car information to be created"
// @Success	    	201   			{object} 	dtos.CarOutputDTO "Successfully created car"
//	@Failure			400					{array}		validation_errors.HTTPError
// @Router				/api/v1/cars/create [post]
func RegisterCarController(context *gin.Context, carRepository repositories.CarRepository) {

	var request dtos.CarInputDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		validation_errors.NewError(context, http.StatusBadRequest, err)
		return
	}
	//fmt.Printf("%+v\n", request)
	car := dtos.CarInputDTO{
		Name:         request.Name,
		Description:  request.Description,
		DailyRate:    request.DailyRate,
		Available:    request.Available,
		LicensePlate: request.LicensePlate,
		FineAmount:   request.FineAmount,
		Brand:        request.Brand,
		CategoryID:   request.CategoryID,
	}

	createdCar, err := usecases.PostCarUseCase(car, carRepository)
	if err != nil {
		validation_errors.NewError(context, http.StatusBadRequest, err)
		return
	} else {
		context.JSON(http.StatusOK, createdCar)
	}
}
