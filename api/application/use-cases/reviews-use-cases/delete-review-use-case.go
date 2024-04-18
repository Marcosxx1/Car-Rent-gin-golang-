package reviewusecases

import (
	"errors"
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrGetRecord      = errors.New("failed to get record: record not found")
	ErrDeleteRecord   = errors.New("failed to delete record")
)

type DeleteReviewUseCase struct {
	reviewRepository repositories.ReviewsRepository
}

func NewDeleteReviewUseCase(reviewRepository repositories.ReviewsRepository) *DeleteReviewUseCase {
	return &DeleteReviewUseCase{
		reviewRepository: reviewRepository,
	}
}

func (useCase *DeleteReviewUseCase) Execute(reviewID string) error {
	existReview, err := useCase.reviewRepository.GetReviewByID(reviewID)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrGetRecord, err)
	}
	if existReview.ID == "" {
		return errors.New("record not found")
	}

	if err := useCase.reviewRepository.DeleteReview(reviewID); err != nil {
		return fmt.Errorf("%w: %s", ErrDeleteRecord, err)
	}

	return nil
}
