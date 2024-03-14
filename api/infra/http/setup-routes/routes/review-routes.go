package routes

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/middlewares"
	factory "github.com/Marcosxx1/Car-Rent-gin-golang-/api/main/factories/controllers"
	"github.com/gin-gonic/gin"
)

func SetupReviewRoutes(router *gin.Engine) {
	authGroup := router.Group("/api/v1").Use(middlewares.JWTMiddleware())
	{
		authGroup.POST("/review/create", factory.PostReviewsFactoryController)
		authGroup.GET("/review/list", factory.GetAllReviewsFactoryController)
		authGroup.PUT("/review/update", factory.UpdateReviewFactoryController)
	}
}
