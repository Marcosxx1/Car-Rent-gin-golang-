package usecases

import (
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
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

	err := useCase.reviewRepository.DeleteReview(reviewID)
	if err != nil {
		return fmt.Errorf("failed delete record: %w", err) // lembrar de adicionar melhor tratamento de erros
	}

	return nil
}
