package routes

import (
	factory "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/factories/orders"
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(authGroup *gin.RouterGroup, router *gin.Engine) {
	router.GET("order", factory.GetOrderByQueryFactoryController)
	authGroup.POST("/order", factory.PostOrderFactoryController)
	authGroup.DELETE("order/id", factory.DeleteOrderFactoryController)
}
