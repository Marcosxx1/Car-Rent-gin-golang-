package reviewendpoints

import (
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecase "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/reviews-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"

	"github.com/gin-gonic/gin"
)

func DeleteReviewController(context *gin.Context, reviewRepository repositories.ReviewsRepository) {
	reviewID := context.Param("review_id")

	deleteReviewUseCase := usecase.NewDeleteReviewUseCase(reviewRepository)

	err := deleteReviewUseCase.Execute(reviewID)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, nil)
}
