package usecases

import (
	"fmt"
	"sync"

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
	resultChan := make(chan []*reviewdto.ReviewOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(2)
	go useCase.performaGetAllReviews(&wg, resultChan, errorChan, validationErrorSignal)

	wg.Add(1)
	go func() {
		defer wg.Done()
		wg.Wait()
		close(resultChan)
		close(errorChan)
		close(validationErrorSignal)
	}()

	select {
	case reviewsOutput := <-resultChan:
		return reviewsOutput, nil
	case err := <-errorChan:
		if <-validationErrorSignal {
			fmt.Println("Some error to be defined TODO")
		}
		return nil, err
	}
}

func (useCase *GetAllReviewsUseCase) performaGetAllReviews(wg *sync.WaitGroup, resultChan chan<- []*reviewdto.ReviewOutputDTO, errorChan chan<- error, validationErrorSignal chan<- bool) {
	defer wg.Done()

	reviews, err := useCase.reviewRepository.GetAllReviews()

	if err != nil {
		errorChan <- fmt.Errorf("failed to retrieve reviews: %w", err)
		validationErrorSignal <- true
		return
	}

	if len(reviews) == 0 {
		errorChan <- fmt.Errorf("no reviews found")
		validationErrorSignal <- true
		return
	}

	go func() {
		outPutDto := reviewdto.ConvertMultipleReviewsToOutPut(reviews)
		validationErrorSignal <- false
		resultChan <- outPutDto
	}()
}
