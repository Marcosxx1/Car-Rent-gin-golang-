package endpoints

import (
	"net/http"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	"github.com/gin-gonic/gin"
)

func RegisterCarController(context *gin.Context) {
	carRepository := database.PGCarRepository{}

	var request usecases.RegisterCarRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	car := usecases.RegisterCarRequest {
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
