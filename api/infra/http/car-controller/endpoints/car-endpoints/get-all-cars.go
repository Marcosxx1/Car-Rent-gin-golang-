package endpoints

import (
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	"github.com/gin-gonic/gin"
)


func ListCarController (context *gin.Context){
	carRepository := database.PGCarRepository{}

	listOfCars, err := usecases.GetAllCarsUseCase(&carRepository)
	if err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(200, listOfCars)

}