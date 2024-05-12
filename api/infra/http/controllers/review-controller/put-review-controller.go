package reviewcontroller

import (
	"net/http"

	reviewdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/review"
	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/reviews-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// PutReviewController handles the HTTP PUT request to update an existing review.
// @Summary             Update an existing review
// @Description         Update an existing review with the provided information
// @ID                  put-review
// @Tags                Reviews
// @Accept              json
// @Produce             json
// @Security 			BearerAuth
// @Param               id              path        string    true  "Review ID to be updated"
// @Param               request         body        reviewdto.ReviewInputDTO  true "Review information to be updated"
// @Success             200             {object}    reviewdto.ReviewOutputDTO "Successfully updated review"
// @Failure             404             {object}    error                "Review not found"
// @Failure             422             {array}     validation_errors.HTTPError
// @Router              /api/v1/reviews/:id [put]
func PutReviewController(context *gin.Context, putReviewUseCase *usecases.UpdateReviewUseCase) {
	reviewID := context.Param("id")

	var request *reviewdto.ReviewInputDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	updatedReview, err := putReviewUseCase.Execute(reviewID, request)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, updatedReview)
}
