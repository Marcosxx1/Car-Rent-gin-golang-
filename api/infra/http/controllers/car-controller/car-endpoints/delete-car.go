package carendpoints

import (
	"log"
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/gin-gonic/gin"
)

func DeleteCarController(context *gin.Context, carRepository repositories.CarRepository) {
 
	id := context.Param("id")

	err := usecases.DeleteCarUseCase(carRepository, id)
	if err != nil {
		log.Println("Error deleting car:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}
