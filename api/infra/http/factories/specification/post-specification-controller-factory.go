package specfactory

import (
	specificationusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/specification-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	specificationcontroller "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/specification-controller"
	"github.com/gin-gonic/gin"
)

func PostSpecificationControllerFactory(context *gin.Context) {
	specificationRepository := database.NewPGSpecificationRepository()

	postSpecificationUseCase := specificationusecases.NewPostSpecificationUseCase(specificationRepository)

	specificationcontroller.PostSpecificationController(context, postSpecificationUseCase)
}
