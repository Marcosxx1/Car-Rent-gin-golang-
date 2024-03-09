package reviewdto

import "github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"

// ReviewInputDTO represents the data required for creating a new review.
type ReviewInputDTO struct {
	Rating  *int   `json:"rating"`
	Content string `json:"content"`
	UserId  string `json:"user_id"`
	CarId   string `json:"car_id"`
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
		// Map other fields as needed
	}
}
