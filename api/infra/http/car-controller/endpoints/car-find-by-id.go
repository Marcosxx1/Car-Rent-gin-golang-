package endpoints

import (
	"log"
	"net/http"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	"github.com/gin-gonic/gin"
)

/* package endpoints

import (
	"log"
	"net/http"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	"github.com/gin-gonic/gin"
)

func DeleteCarController(context *gin.Context) {
	carRepository := database.PGCarRepository{}

	id := context.Param("id")

	err := usecases.DeleteCarUseCase(&carRepository, id)
	if err != nil {
		log.Println("Error deleting car:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}
*/
func FindCarByIdController(context *gin.Context) {
	carRepository := database.PGCarRepository{}

	id := context.Param("id")

	car, err := usecases.FindCarByIdUseCase(id, &carRepository )
	if err != nil {
		log.Println("Error finding car:", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, car)

	
}