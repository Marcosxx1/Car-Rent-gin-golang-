// routes/routes.go
package routes

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	endpoints "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller/car-endpoints"
	"github.com/gin-gonic/gin"
)

func SetupCarRoutes(router *gin.Engine) {
	carRepository := database.PGCarRepository{}
	specificationRepository := database.PGSpecification{}

	router.GET("/api/v1/cars", func(context *gin.Context) {
		endpoints.GetAllCarsController(context, &carRepository, &specificationRepository)
	})

	router.GET("/api/v1/cars/:id", func(context *gin.Context) {
			endpoints.FindCarByIdController(context, &carRepository, &specificationRepository) 
	})

	router.POST("/api/v1/cars/create", func(context *gin.Context) {
			endpoints.RegisterCarController(context, &carRepository, &specificationRepository)
	})

	router.DELETE("/api/v1/cars/delete/:id", func(context *gin.Context) {
			endpoints.DeleteCarController(context, &carRepository)
	})

	router.PUT("/api/v1/cars/update/:id", func(context *gin.Context) {
			endpoints.UpdateCarController(context, &carRepository, &specificationRepository)
	})
}
