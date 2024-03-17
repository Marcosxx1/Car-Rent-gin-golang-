package carfactory

import (
	carusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	carcontroller "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller"
	"github.com/gin-gonic/gin"
)

func UpdateCarControllerFacotry(context *gin.Context) {
	carRepository := database.NewPGCarRepository()
	specificationRepository := database.NewPGSpecificationRepository()

	updateCarUseCase := carusecases.NewUpdateCarUseCase(carRepository, specificationRepository)

	carcontroller.UpdateCarController(context, updateCarUseCase)
}
