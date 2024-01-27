package carendpoints

import (
	"log"
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/gin-gonic/gin"
)

func FindCarByIdController(context *gin.Context,carRepository repositories.CarRepository) {

	id := context.Param("id")

	car, err := usecases.GetCarByIdUseCase(id, carRepository )
	if err != nil {
		log.Println("Error finding car:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, car)
}