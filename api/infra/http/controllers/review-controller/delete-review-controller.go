package reviewcontroller

import (
	"net/http"

	usecase "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/reviews-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"

	"github.com/gin-gonic/gin"
)

// @Summary Delete a review
// @Description Delete a review with the provided ID
// @ID delete-review
// @Tags Reviews
// @Produces json
// @Security 			BearerAuth
// @Param id path string true "Review ID to be delete"
// @Success 200
// @Failure 400 {object} validation_errors.HTTPErrorCar "Error details"
// @Router				/api/v1/review/{review_id} [delete]
func DeleteReviewController(context *gin.Context, deleteReviewUseCase *usecase.DeleteReviewUseCase) {
	reviewID := context.Param("review_id")

	err := deleteReviewUseCase.Execute(reviewID)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, nil)
}
