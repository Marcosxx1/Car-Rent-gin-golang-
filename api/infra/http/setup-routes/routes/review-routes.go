package routes

import (
	factory "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/factories/reviews"
	"github.com/gin-gonic/gin"
)

func SetupReviewRoutes(router *gin.Engine, authGroup *gin.RouterGroup) {
	authGroup.POST("/review/create", factory.PostReviewsFactoryController)
	router.GET("/review/list", factory.GetAllReviewsFactoryController)
	authGroup.PUT("/review/:id", factory.UpdateReviewFactoryController)
	authGroup.DELETE("/review/:id", factory.DeleteReviewFactoryController)
}
