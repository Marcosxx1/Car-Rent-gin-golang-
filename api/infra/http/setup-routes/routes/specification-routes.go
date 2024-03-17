package routes

import (
	specfactory "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/factories/specification"
	"github.com/gin-gonic/gin"
)

func SetupSpecificationRoutes(router *gin.Engine) {
	router.POST("/specification/create", specfactory.PostSpecificationControllerFactory)
}
