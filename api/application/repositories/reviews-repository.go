package repositories

import "github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"

type ReviewsRepository interface {
	// CreateReview adds a new review to the repository and returns *domain.Reviews
	CreateReview(review *domain.Reviews) (*domain.Reviews, error)
	// GetReviewByID retrieves a review from the repository based on its ID
	GetReviewByID(reviewID string) (*domain.Reviews, error)
	// GetAllReviews retrieves all reviews from the repository
	GetAllReviews() ([]*domain.Reviews, error)
	// DeleteReview removes a review from the repository based on its ID
	DeleteReview(reviewID string) error
	// UpdateReview updates an existing review in the repository
	UpdateReview(id string, review *domain.Reviews) error

}
