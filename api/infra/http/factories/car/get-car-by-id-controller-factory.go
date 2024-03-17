package carfactory

import (
	carusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	carcontroller "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller"
	"github.com/gin-gonic/gin"
)

func FindCarByIdControllerFacotry(context *gin.Context) {
	carRepository := database.NewPGCarRepository()
	specificationRepository := database.NewPGSpecificationRepository()

	findByIdUseCase := carusecases.NewFindCarByIdUseCase(carRepository, specificationRepository)

	carcontroller.FindCarByIdController(context, findByIdUseCase)
}
