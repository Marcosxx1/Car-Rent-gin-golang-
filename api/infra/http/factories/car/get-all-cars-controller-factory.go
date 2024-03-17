package carfactory

import (
	carusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	carcontroller "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller"
	"github.com/gin-gonic/gin"
)

func GetAllCarsControllerFactory(context *gin.Context) {
	carRepository := database.NewPGCarRepository()
	specificationRepository := database.NewPGSpecificationRepository()

	findAllCarsUseCase := carusecases.NewGetAllCarsUseCase(carRepository, specificationRepository)

	carcontroller.GetAllCarsController(context, findAllCarsUseCase)
}
