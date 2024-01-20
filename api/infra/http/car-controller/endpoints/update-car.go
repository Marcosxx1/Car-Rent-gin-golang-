package endpoints

import (
	"fmt"
	"log"
	"net/http"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/car-controller/dtos"
	"github.com/gin-gonic/gin"
)

func UpdateCarController(context *gin.Context) {
	carRepository := database.PGCarRepository{}

	var request dtos.CarDto
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%+v\n", request)

	id := context.Param("id")

	foundCar, err := usecases.GetCarByIdUseCase(id, &carRepository)
	if err != nil {
		log.Println("Error finding car:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if foundCar == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	car := domain.Car{
		Name:         request.Name,
		Description:  request.Description,
		DailyRate:    request.DailyRate,
		Available:    request.Available,
		LicensePlate: request.LicensePlate,
		FineAmount:   request.FineAmount,
		Brand:        request.Brand,
		CategoryId:   request.CategoryId,
	}

	updatedCar, err := usecases.PutCarUseCase(id, &car, &carRepository)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, updatedCar)
	}
}
