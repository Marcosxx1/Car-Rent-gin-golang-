package reviewendpoints

import (
	"net/http"

	usecase "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/reviews-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"

	"github.com/gin-gonic/gin"
)

func DeleteReviewController(context *gin.Context, deleteReviewUseCase *usecase.DeleteReviewUseCase) {
	reviewID := context.Param("review_id")

	err := deleteReviewUseCase.Execute(reviewID)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, nil)
}
