package carendpoints

import (
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/car-controller/car-dtos"
	"github.com/gin-gonic/gin"
)

func RegisterCarController(context *gin.Context, carRepository repositories.CarRepository) {

	var request dtos.CarInputDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//fmt.Printf("%+v\n", request)
	car := dtos.CarInputDTO {
		Name: request.Name,
		Description: request.Description,
		DailyRate: request.DailyRate,
		Available: request.Available,
		LicensePlate: request.LicensePlate,
		FineAmount: request.FineAmount,
		Brand: request.Brand,
		CategoryID: request.CategoryID,
	}
	
	createdCar, err := usecases.PostCarUseCase(car, carRepository)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, createdCar)
	}
}
