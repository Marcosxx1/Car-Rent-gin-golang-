package carendpoints

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/gin-gonic/gin"
)


func ListCarController (context *gin.Context, carRepository repositories.CarRepository){
 
	listOfCars, err := usecases.GetAllCarsUseCase(carRepository)
	if err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(200, listOfCars)

}