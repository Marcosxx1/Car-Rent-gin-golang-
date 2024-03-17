package carfactory

import (
	carusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	carcontroller "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller"
	"github.com/gin-gonic/gin"
)

func RegisterCarControllerFacotry(context *gin.Context) {
	carRepository := database.NewPGCarRepository()
	specificationRepository := database.NewPGSpecificationRepository()

	registerCarUseCase := carusecases.NewPostCarUseCase(carRepository, specificationRepository)

	carcontroller.RegisterCarController(context, registerCarUseCase)
}
