package routes

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	reviewendpoints "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/review-controller/review-endpoints"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupReviewRoutes(router *gin.Engine) {
	reviewRepository := database.PGReviewRepository{}

	authGroup := router.Group("/api/v1").Use(middlewares.JWTMiddleware())
	{
		authGroup.POST("/review/create", func(context *gin.Context) {
			reviewendpoints.PostReviewsController(context, &reviewRepository)
		})
		authGroup.GET("/review/list", func(context *gin.Context) {
			reviewendpoints.GetAllReviewsController(context, &reviewRepository)
		})
	}
}
