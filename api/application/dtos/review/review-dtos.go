package reviewdto

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

// ReviewInputDTO represents the data required for creating a new review.
type ReviewInputDTO struct {
	Rating  *int   `json:"rating" validate:"required,gte=1,lte=5"`
	Content string `json:"content" validate:"required,max=500"`
	UserId  string `json:"user_id" validate:"required"`
	CarId   string `json:"car_id" validate:"required"`
}

type ReviewOutputDTO struct {
	ID      string
	UserId  string
	CarId   string
	Rating  *int
	Content string
}

// ConvertReviewToOutput converts a domain.Reviews object to a ReviewOutputDTO.
func ConvertReviewToOutput(review *domain.Reviews) *ReviewOutputDTO {
	return &ReviewOutputDTO{
		ID:      review.ID,
		UserId:  review.UserId,
		CarId:   review.CarId,
		Rating:  review.Rating,
		Content: review.Content,
	}
}

// Will convert multiple reviews to output
func ConvertMultipleReviewsToOutPut(domainReviews []*domain.Reviews) []*ReviewOutputDTO {
	reviews := make([]*ReviewOutputDTO, len(domainReviews))

	for index, rev := range domainReviews {
		reviews[index] = ConvertReviewToOutput(rev)
	}

	return reviews
}
