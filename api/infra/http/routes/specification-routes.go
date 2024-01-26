package routes

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	specificationendpoints "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/specification-controller/specification-endpoints"
	"github.com/gin-gonic/gin"
)

func SetupSpecificationRoutes(router *gin.Engine) {
	specificationRepository := database.PGSpecification{}

	router.GET("/api/v1/specification/create", func(context *gin.Context) {
		specificationendpoints.PostSpecificationController(context, &specificationRepository)
	})
}
