package usecases

import (
	"fmt"
	"sync"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	reviewdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/review-controller/review-dto"
)

type PostReviewUseCase struct {
	reviewRepository repositories.ReviewsRepository
}

func NewPostReviewUseCase(reviewRepository repositories.ReviewsRepository) *PostReviewUseCase {
	return &PostReviewUseCase{
		reviewRepository: reviewRepository,
	}
}

func (useCase *PostReviewUseCase) Execute(userID string, carID string, inputDTO *reviewdto.ReviewInputDTO) (*reviewdto.ReviewOutputDTO, error) {
	resultChan := make(chan *reviewdto.ReviewOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(2)
	go validateReviewInput(&wg, errorChan, validationErrorSignal, inputDTO)
	go useCase.performReviewCreation(&wg, resultChan, errorChan, validationErrorSignal, userID, carID, inputDTO)

	wg.Add(1)
	go func() {
		defer wg.Done()
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	select {
	case createdReview := <-resultChan:
		return createdReview, nil
	case err := <-errorChan:
		return nil, err
	}
}

func validateReviewInput(wg *sync.WaitGroup, errorChan chan<- error, validationErrorSignal chan<- bool, inputDTO *reviewdto.ReviewInputDTO) {
	defer wg.Done()

	if inputDTO.Rating == nil || *inputDTO.Rating < 1 || *inputDTO.Rating > 5 {
		errorChan <- fmt.Errorf("invalid rating")
		validationErrorSignal <- true
		return
	}

	validationErrorSignal <- false
}

func (useCase *PostReviewUseCase) performReviewCreation(wg *sync.WaitGroup, resultChan chan<- *reviewdto.ReviewOutputDTO, errorChan chan<- error, validationErrorSignal <-chan bool, userID string, carID string, inputDTO *reviewdto.ReviewInputDTO) {
	defer wg.Done()

	if <-validationErrorSignal {
		return
	}

	newReview := &domain.Reviews{
		UserId:  userID,
		CarId:   carID,
		Rating:  inputDTO.Rating,
		Content: inputDTO.Content,
	}

	createdReviewID, err := useCase.reviewRepository.CreateReview(newReview)
	if err != nil {
		errorChan <- fmt.Errorf("failed to create review: %w", err)
		return
	}

	outputDTO := reviewdto.ConvertReviewToOutput(createdReviewID)

	resultChan <- outputDTO
}
