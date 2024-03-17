// reviewcontroller/factory.go
package factory

import (
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/reviews-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	reviewcontroller "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/review-controller"
	"github.com/gin-gonic/gin"
)

func PostReviewsFactoryController(context *gin.Context) {
	repository := database.NewPGReviewRepository()
	postReviewUseCase := usecases.NewPostReviewUseCase(repository)
	reviewcontroller.PostReviewsController(context, postReviewUseCase)
}
