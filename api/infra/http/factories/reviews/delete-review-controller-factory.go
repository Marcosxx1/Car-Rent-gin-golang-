package factory

import (
	usecase "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/reviews-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	reviewendpoints "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/review-controller/review-endpoints"
	"github.com/gin-gonic/gin"
)

func DeleteReviewFactoryController(context *gin.Context) {
	repository := database.NewPGReviewRepository()
	deleteReviewUseCase := usecase.NewDeleteReviewUseCase(repository)
	reviewendpoints.DeleteReviewController(context, deleteReviewUseCase)
}
