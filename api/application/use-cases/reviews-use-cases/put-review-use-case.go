package reviewusecases

import (
	"errors"
	"fmt"
	"sync"

	reviewdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/review"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

type UpdateReviewUseCase struct {
	reviewRepository repositories.ReviewsRepository
}

func NewUpdateReviewUseCase(reviewRepository repositories.ReviewsRepository) *UpdateReviewUseCase {
	return &UpdateReviewUseCase{
		reviewRepository: reviewRepository,
	}
}

func (useCase *UpdateReviewUseCase) Execute(reviewID string, inputDTO *reviewdto.ReviewInputDTO) (*reviewdto.ReviewOutputDTO, error) {
	resultChan := make(chan *reviewdto.ReviewOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(2)
	go validateReviewInput(&wg, errorChan, validationErrorSignal, inputDTO)
	go useCase.performReviewUpdate(&wg, resultChan, errorChan, validationErrorSignal, reviewID, inputDTO)

	wg.Add(1)
	go func() {
		defer wg.Done()
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	select {
	case updatedReview := <-resultChan:
		return updatedReview, nil
	case err := <-errorChan:
		return nil, err
	}
}

func (useCase *UpdateReviewUseCase) performReviewUpdate(wg *sync.WaitGroup, resultChan chan<- *reviewdto.ReviewOutputDTO, errorChan chan<- error, validationErrorSignal chan bool, reviewID string, inputDTO *reviewdto.ReviewInputDTO) {
	defer wg.Done()

	if <-validationErrorSignal {
		errorChan <- errors.New("not possible to perform review update")
		return
	}

	existsReview, err := useCase.reviewRepository.GetReviewByID(reviewID)
	if err != nil {
		errorChan <- fmt.Errorf("failed to get review: %w", err)
		validationErrorSignal <- true
		return
	}

	if existsReview.ID == "" {
		errorChan <- fmt.Errorf("review not found")
		validationErrorSignal <- true
		return
	}

	updatedReview := &domain.Reviews{
		ID:      reviewID,
		UserId:  existsReview.UserId,
		CarId:   existsReview.CarId,
		Rating:  inputDTO.Rating,
		Content: inputDTO.Content,
	}

	err = useCase.reviewRepository.UpdateReview(reviewID, updatedReview)
	if err != nil {
		errorChan <- fmt.Errorf("failed to update review: %w", err)
		return
	}

	outputDTO := reviewdto.ConvertReviewToOutput(updatedReview)

	resultChan <- outputDTO
}
