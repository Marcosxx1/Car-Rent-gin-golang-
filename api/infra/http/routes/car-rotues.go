// routes/routes.go
package routes

import (
	endpoints "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/car-controller/car-endpoints"
	"github.com/gin-gonic/gin"
)

func SetupCarRoutes(router *gin.Engine) {
	router.GET("/api/v1/cars", endpoints.ListCarController)
	router.GET("/api/v1/cars/:id", endpoints.FindCarByIdController)
	router.POST("/api/v1/cars/create", endpoints.RegisterCarController)
	router.DELETE("/api/v1/cars/delete/:id", endpoints.DeleteCarController)
	router.PUT("/api/v1/cars/update/:id", endpoints.UpdateCarController)
}
