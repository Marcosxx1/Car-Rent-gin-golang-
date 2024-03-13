package factory

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	reviewendpoints "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/review-controller/review-endpoints"
	"github.com/gin-gonic/gin"
)

func GetAllReviewsFactoryController(context *gin.Context) {
    repository := database.PGReviewRepository{}
    reviewendpoints.GetAllReviewsController(context, &repository)
}
