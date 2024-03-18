// routes/routes.go
package routes

import (
	carfactory "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/factories/car"
	"github.com/gin-gonic/gin"
)

func SetupCarRoutes(router *gin.Engine, authGroup *gin.RouterGroup) {

	authGroup.POST("/cars/create", carfactory.RegisterCarControllerFacotry)
	authGroup.PUT("/cars/update/:id", carfactory.UpdateCarControllerFacotry)
	authGroup.DELETE("/cars/delete/:id", carfactory.DeleteCarControllerFacotry)
	router.GET("/cars", carfactory.GetAllCarsControllerFactory)
	authGroup.GET("/cars/:id", carfactory.FindCarByIdControllerFacotry)

}
