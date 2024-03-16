package reviewendpoints

import (
	"net/http"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/reviews-use-cases"
	reviewdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/review-controller/review-dto"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// PostReviewsController handles the HTTP POST request to create a new review.
// @Summary             Create a new review
// @Description         Create a new review with the provided information
// @ID                  post-review
// @Tags                Reviews
// @Accept              json
// @Produce             json
// @Param               request         body        dtos.ReviewInputDTO  true "Review information to be created"
// @Success             201             {object}    dtos.ReviewOutputDTO "Successfully created review"
// @Failure             422             {array}     validation_errors.HTTPErrorReview
// @Router              /api/v1/reviews/create [post]
func PostReviewsController(context *gin.Context, postReviewUseCase *usecases.PostReviewUseCase) {
	userID := context.GetString("user_id")

	var request *reviewdto.ReviewInputDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	carID := request.CarId

	createdReview, err := postReviewUseCase.Execute(userID, carID, request)
	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, createdReview)
}
