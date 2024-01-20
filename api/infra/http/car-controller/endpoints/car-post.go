package endpoints

import (
	"fmt"
	"net/http"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/car-controller/dtos"
	"github.com/gin-gonic/gin"
)

func RegisterCarController(context *gin.Context) {
	carRepository := database.PGCarRepository{}

	var request dtos.CarDto
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%+v\n", request)
	car := dtos.CarDto {
		Name: request.Name,
		Description: request.Description,
		DailyRate: request.DailyRate,
		Available: request.Available,
		LicensePlate: request.LicensePlate,
		FineAmount: request.FineAmount,
		Brand: request.Brand,
		CategoryId: request.CategoryId,

	}
	createdCar, err := usecases.RegisterCarUseCase(car, &carRepository)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, createdCar)
	}
}
