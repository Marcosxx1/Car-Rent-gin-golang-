package reviewendpoints

import (
	"net/http"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/reviews-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// GetAllReviewsController handles the HTTP GET request to find all reviews.
// @Summary             find all reviews
// @Description         find all reviews
// @ID                  get-all-review
// @Tags                Reviews
// @Accept              json
// @Produce             json
// @Success             201             {object}   []dtos.ReviewInputDTO "Successfully created review"
// @Failure             422             {array}     validation_errors.HTTPErrorReview
// @Router              /api/v1/reviews/create [post]
func GetAllReviewsController(context *gin.Context, reviewRepository repositories.ReviewsRepository) {

	getAllReviewsUseCase := usecases.NewGetAllReviewsUseCase(reviewRepository)

	allReviews, err := getAllReviewsUseCase.Execute()
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, allReviews)
}
