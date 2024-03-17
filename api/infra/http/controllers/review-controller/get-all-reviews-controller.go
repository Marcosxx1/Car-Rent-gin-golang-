package reviewcontroller

import (
	"net/http"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/reviews-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

func GetAllReviewsController(context *gin.Context, getAllReviewsUseCase *usecases.GetAllReviewsUseCase) {

	allReviews, err := getAllReviewsUseCase.Execute()
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, allReviews)
}
