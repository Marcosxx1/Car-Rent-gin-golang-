package carendpoints

import (
	"log"
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/car-controller/car-dtos"
	"github.com/gin-gonic/gin"
)

func UpdateCarController(context *gin.Context, carRepository repositories.CarRepository) {

	var request dtos.CarOutputDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := context.Param("id")

	foundCar, err := usecases.GetCarByIdUseCase(id, carRepository)
	if err != nil {
		log.Println("Error finding car:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if foundCar == nil { 
		context.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	car := &dtos.CarOutputDTO{
		ID:           id,
		Name:         request.Name,
		Description:  request.Description,
		DailyRate:    request.DailyRate,
		Available:    request.Available,
		LicensePlate: request.LicensePlate,
		FineAmount:   request.FineAmount,
		Brand:        request.Brand,
		CategoryID:   request.CategoryID,
	}

	updatedCar, err := usecases.PutCarUseCase(id, *car, carRepository)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, updatedCar)
	}
}
