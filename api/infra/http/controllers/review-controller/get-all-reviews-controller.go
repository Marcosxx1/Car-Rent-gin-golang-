package reviewcontroller

import (
	"net/http"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/reviews-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// GetAllReviewsController handles the HTTP GET request to create a new review.
// @Summary             Get all reviews
// @Description         Get all reviewsn
// @ID                  get-review
// @Tags                Reviews
// @Accept              json
// @Produce             json
// @Security 			BearerAuth
// @Success             201             {object}    []reviewdto.ReviewOutputDTO "Successfully created review"
// @Router              /api/v1/review/list [get]
func GetAllReviewsController(context *gin.Context, getAllReviewsUseCase *usecases.GetAllReviewsUseCase) {

	allReviews, err := getAllReviewsUseCase.Execute()
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, allReviews)
}
