package carendpoints

import (
	"log"
	"net/http"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	"github.com/gin-gonic/gin"
)

func FindCarByIdController(context *gin.Context) {
	carRepository := database.PGCarRepository{}

	id := context.Param("id")

	car, err := usecases.GetCarByIdUseCase(id, &carRepository )
	if err != nil {
		log.Println("Error finding car:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, car)

	
}