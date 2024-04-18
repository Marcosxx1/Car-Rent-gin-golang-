package reviewusecases

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	reviewdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/review"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	databasemocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/database-mocks"
)


func TestUpdateReviewUseCase_Execute(t *testing.T) {
	// Mock review data
	reviewID := "1"
	rating := 5
	inputDTO := &reviewdto.ReviewInputDTO{
		UserId:  "user1",
		CarId:   "car1",
		Rating:  &rating,
		Content: "Great car!",
	}

	existingRating := 4
	existingReview := &domain.Reviews{
		ID:      "1",
		UserId:  "user1",
		CarId:   "car1",
		Rating:  &existingRating,
		Content: "Good car",
	}

	// Mock repository
	mockRepo := new(databasemocks.MockReviewRepository)
	mockRepo.On("GetReviewByID", reviewID).Return(existingReview, nil)
	mockRepo.On("UpdateReview", reviewID, mock.Anything).Return(nil)

	// Create use case
	useCase := NewUpdateReviewUseCase(mockRepo)

	// Execute use case
	outputDTO, err := useCase.Execute(reviewID, inputDTO)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, outputDTO)
	assert.Equal(t, inputDTO.Rating, outputDTO.Rating)
	assert.Equal(t, inputDTO.Content, outputDTO.Content)
}
 
func TestUpdateReviewUseCase_Execute_Error(t *testing.T) {
	// Mock review data
	reviewID := "1"
	rating := 5

	inputDTO := &reviewdto.ReviewInputDTO{
		UserId:  "user1",
		CarId:   "car1",
		Rating:  &rating,
		Content: "Great car!",
	}

	// Mock repository
	mockRepo := new(databasemocks.MockReviewRepository)
	mockRepo.On("GetReviewByID", reviewID).Return((*domain.Reviews)(nil), errors.New("failed to get review"))

	// Create use case
	useCase := NewUpdateReviewUseCase(mockRepo)

	// Execute use case
	_, err := useCase.Execute(reviewID, inputDTO)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get review")
}

func TestUpdateReviewUseCase_Execute_ReviewNotFound(t *testing.T) {
	// Mock review data
	reviewID := "1"
	rating := 5
	inputDTO := &reviewdto.ReviewInputDTO{
		UserId:  "user1",
		CarId:   "car1",
		Rating:  &rating,
		Content: "Great car!",
	}

	// Mock repository
	mockRepo := new(databasemocks.MockReviewRepository)
	mockRepo.On("GetReviewByID", reviewID).Return(&domain.Reviews{}, nil)

	// Create use case
	useCase := NewUpdateReviewUseCase(mockRepo)

	// Execute use case
	_, err := useCase.Execute(reviewID, inputDTO)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "review not found")
}

func TestUpdateReviewUseCase_performReviewUpdate_ValidationError(t *testing.T) {
	// Mock data
	rating := 6
	inputDTO := &reviewdto.ReviewInputDTO{
		UserId:  "user1",
		CarId:   "",
		Rating:  &rating,
		Content: "",
	}

	// Mock repository
	mockRepo := new(databasemocks.MockReviewRepository)
	
	// Create use case
	useCase := NewUpdateReviewUseCase(mockRepo)

	// Channels
	resultChan := make(chan *reviewdto.ReviewOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool, 1)
	validationErrorSignal <- true

	// Execute performReviewUpdate
	var wg sync.WaitGroup
	wg.Add(1)
	go useCase.performReviewUpdate(&wg, resultChan, errorChan, validationErrorSignal, "reviewID", inputDTO)

	// Wait for the goroutine to complete or timeout after 5 seconds
	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	// Assertions
	select {
	case errs :=  <-resultChan:
		t.Errorf("Expected no result from resultChan, err :%v", errs)
	case err := <-errorChan:
		assert.Error(t, err) // Assert that an error is sent to errorChan
		assert.Contains(t, err.Error(), "not possible to perform review update") // Assert specific error message
	}
}

func TestUpdateReviewUseCase_performReviewUpdate_FailedToUpdateReview(t *testing.T) {
	// Mock data
	rating := 6
	inputDTO := &reviewdto.ReviewInputDTO{
		UserId:  "user1",
		CarId:   "user1",
		Rating:  &rating,
		Content: "user1",
	}

	// Mock repository
	mockRepo := new(databasemocks.MockReviewRepository)
	mockRepo.On("UpdateReview", "reviewID", mock.AnythingOfType("*domain.Reviews")).Return(errors.New("failed to update review"))
	mockRepo.On("GetReviewByID", "reviewID").Return(&domain.Reviews{ID: "any_id"}, nil)

	// Create use case
	useCase := NewUpdateReviewUseCase(mockRepo)

	// Channels
	resultChan := make(chan *reviewdto.ReviewOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool, 1)
	validationErrorSignal <- false

	// Execute performReviewUpdate
	var wg sync.WaitGroup
	wg.Add(1)
	go useCase.performReviewUpdate(&wg, resultChan, errorChan, validationErrorSignal, "reviewID", inputDTO)

	// Wait for the goroutine to complete or timeout after 5 seconds
	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	// Assertions
	select {
	case errs :=  <-resultChan:
		t.Errorf("Expected no result from resultChan, err :%v", errs)
	case err := <-errorChan:
		assert.Error(t, err) // Assert that an error is sent to errorChan
		assert.Contains(t, err.Error(), "failed to update review: failed to update review") // Assert specific error message
	}
}

// %v value

/* func TestUpdateReviewUseCase_performReviewUpdate_Success(t *testing.T) {
	// Mock data
	rating := 4
	inputDTO := &reviewdto.ReviewInputDTO{
		UserId:  "user1",
		CarId:   "car1",
		Rating:  &rating,
		Content: "Great car!",
	}

	// Mock repository
	mockRepo := new(databasemocks.MockReviewRepository)
	mockRepo.On("UpdateReview", "reviewID", mock.AnythingOfType("*domain.Reviews")).Return(nil)

	// Create use case
	useCase := NewUpdateReviewUseCase(mockRepo)

	// Channels
	resultChan := make(chan *reviewdto.ReviewOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool)

	// Execute performReviewUpdate
	var wg sync.WaitGroup
	wg.Add(1)
	go useCase.performReviewUpdate(&wg, resultChan, errorChan, validationErrorSignal, "reviewID", inputDTO)

	// Wait for the goroutine to complete or timeout after 5 seconds
	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	// Assertions
	select {
	case errs := <-resultChan:
		assert.NotNil(t, errs) // Assert that a result is sent to resultChan
	case <-errorChan:
		t.Errorf("Expected no error from errorChan")
	}
}
 */