package factory

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	reviewendpoints "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/review-controller/review-endpoints"
	"github.com/gin-gonic/gin"
)

func UpdateReviewFactoryController(context *gin.Context) {
	repository := database.NewPGReviewRepository()
	reviewendpoints.PutReviewController(context, repository)
}
