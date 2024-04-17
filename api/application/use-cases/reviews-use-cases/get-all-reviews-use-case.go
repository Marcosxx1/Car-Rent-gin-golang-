package usecases

import (
	"errors"

	reviewdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/review"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
)

type GetAllReviewsUseCase struct {
	reviewRepository repositories.ReviewsRepository
}

func NewGetAllReviewsUseCase(reviewRepository repositories.ReviewsRepository) *GetAllReviewsUseCase {
	return &GetAllReviewsUseCase{
		reviewRepository: reviewRepository,
	}
}

func (useCase *GetAllReviewsUseCase) Execute() ([]*reviewdto.ReviewOutputDTO, error) {
	reviews, err := useCase.reviewRepository.GetAllReviews()
	if err != nil {
		return nil, errors.New("error retrieving reviews")
	}

	if len(reviews) == 0 {
		return nil, errors.New("no reviews found")
	}

	outputDTOs := make([]*reviewdto.ReviewOutputDTO, len(reviews))

	for i, review := range reviews {
		outputDTOs[i] = &reviewdto.ReviewOutputDTO{
			ID:     review.ID,
			UserId: review.UserId,
			CarId:  review.CarId,
			Rating: review.Rating,
		} 
	}
	// simple is better!!!
	return outputDTOs, nil
}
