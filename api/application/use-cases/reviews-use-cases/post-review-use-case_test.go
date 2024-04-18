package reviewusecases

/* 	a := true
if a == true {
	t.Errorf("Error: %s", err)
}

TODO - help debug test*/

import (
	"errors"
	"testing"

	reviewdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/review"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	databasemocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/database-mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostReviewUseCase_Success(t *testing.T) {
	mockRepo := new(databasemocks.MockReviewRepository)
	useCase := NewPostReviewUseCase(mockRepo)

	id := 5
	userID := "user123"
	carID := "car123"
	content := "Some content"

	inputDTO := &reviewdto.ReviewInputDTO{
		Rating:  &id,
		Content: content,
		UserId:  userID,
		CarId:   carID,
	}

	mockRepo.On("CreateReview", mock.AnythingOfType("*domain.Reviews")).Return(&domain.Reviews{ID: "review123"}, nil)

	outputDTO, err := useCase.Execute(userID, carID, inputDTO)

	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, outputDTO, "Expected a non-nil outputDTO")
	assert.Equal(t, "review123", outputDTO.ID, "Expected review ID to match")

	mockRepo.AssertExpectations(t)
}

func TestPostReviewUseCase_ValidationFailed(t *testing.T) {
	mockRepo := new(databasemocks.MockReviewRepository)
	useCase := NewPostReviewUseCase(mockRepo)

	id := 0
	content := ""
	userID := ""
	carID := ""

	inputDTO := &reviewdto.ReviewInputDTO{
		Rating:  &id,
		Content: content,
		UserId:  userID,
		CarId:   carID,
	}

	// We're not expecting any repository calls since validation should fail before that
	mockRepo.AssertNotCalled(t, "CreateReview", mock.AnythingOfType("*domain.Reviews"))

	outputDTO, err := useCase.Execute(userID, carID, inputDTO)

	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	assert.Contains(t, err.Error(), "content is required", "Expected error message for empty content")
/* 	assert.Contains(t, err.Error(), "userid is required", "Expected error message for empty UserID") // since we're passing the id via token on the context, this may not be necessary, TODO
 */	assert.Contains(t, err.Error(), "carid is required", "Expected error message for empty CarID")

	assert.Nil(t, outputDTO, "Expected a nil outputDTO due to validation failure")

	mockRepo.AssertExpectations(t)
}

func TestPostReviewUseCase_RepositoryError(t *testing.T) {
	mockRepo := new(databasemocks.MockReviewRepository)
	useCase := NewPostReviewUseCase(mockRepo)

	rating := 5

	content := "content is required"
	userID := "123wdew"
	carID := "123wdew"

	inputDTO := &reviewdto.ReviewInputDTO{
		Rating:  &rating,
		Content: content,
		UserId:  userID,
		CarId:   carID,
	}

	mockRepo.On("CreateReview", mock.AnythingOfType("*domain.Reviews")).Return(&domain.Reviews{}, errors.New("repository error"))

	_, err := useCase.Execute(userID, carID, inputDTO)

	assert.Error(t, err, "Expected an error")
	assert.Contains(t, err.Error(), "repository error", "Expected error message to contain 'repository error'")

	mockRepo.AssertExpectations(t)
}

/* func TestValidateReviewInput_InvalidDTO(t *testing.T) {
	// Create buffered channels
	errorChan := make(chan error, 1)
	validationErrorSignal := make(chan bool, 1)

	// Create an invalid inputDTO with missing required fields
	inputDTO := &reviewdto.ReviewInputDTO{}

	// Create a WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)

	// Execute validateReviewInput function
	go validateReviewInput(&wg, errorChan, validationErrorSignal, inputDTO)

	// Wait for validation to complete
	wg.Wait()

	// Close channels after sending
	close(errorChan)
	close(validationErrorSignal)

	// Validate the interactions
	err, ok := <-errorChan

	if ok {
		t.Errorf("#########%s ", err)
		t.Fatal("Expected an error")
	}
	assert.Error(t, err, "Expected an error")
	assert.Contains(t, err.Error(), "invalid", "Expected error message to indicate invalid input")

	validSignal, ok := <-validationErrorSignal
	if !ok {
		t.Fatal("Expected a validation error signal")
	}
	assert.True(t, validSignal, "Expected validation error signal to be true")
	assert.Contains(t, err.Error(), "invalid rating", "Expected error message to contain 'invalid rating'")

} */

/* func TestValidateReviewInput_InvalidDTO(t *testing.T) {
	// Create channels
	errorChan := make(chan error, 1)
	validationErrorSignal := make(chan bool, 1)

	// Create an invalid inputDTO with missing required fields
	inputDTO := &reviewdto.ReviewInputDTO{}

	// Create a WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)

	// Execute validateReviewInput function
	go validateReviewInput(&wg, errorChan, validationErrorSignal, inputDTO)

	// Wait for validation to complete
	wg.Wait()

	// Validate the interactions
	err, ok := <-errorChan

	if ok {
		t.Errorf("error: %s, ok: %t", err, ok)
	}
	validSignal, ok := <-validationErrorSignal
	if ok {
		t.Fatal("Expected a validation error signal")
	}

	// Assertions
	assert.Error(t, err, "Expected an error")
	assert.Contains(t, err.Error(), "invalid", "Expected error message to indicate invalid input")
	assert.True(t, validSignal, "Expected validation error signal to be true")
}

func TestValidateReviewInput_InvalidRating(t *testing.T) {
 	mockErrorChan := make(chan error, 1)
	mockValidationErrorSignal := make(chan bool, 1)

 	id := 15
	inputDTO := &reviewdto.ReviewInputDTO{
		Rating: &id,
	}

 	var wg sync.WaitGroup
	wg.Add(1)
	go validateReviewInput(&wg, mockErrorChan, mockValidationErrorSignal, inputDTO)
	wg.Wait()

 	err := <-mockErrorChan
	validSignal := <-mockValidationErrorSignal

 	assert.Error(t, err, "Expected an error")
	assert.Contains(t, err.Error(), "invalid rating", "Expected error message to contain 'invalid rating'")
	assert.True(t, validSignal, "Expected validation error signal to be true")
}

func TestPostReviewUseCase_ValidationErrors(t *testing.T) { we'll need to test  api\application\use-cases\reviews-use-cases\post-review-use-case.go lines 60-64
	mockRepo := new(databasemocks.MockReviewRepository)
	mockValidator := new(validation_errors.MockValidator)
	useCase := NewPostReviewUseCase(mockRepo, mockValidator)

	id := 5

	inputDTO := &reviewdto.ReviewInputDTO{
		Rating: &id,
	}
	userID := "user123"
	carID := "car123"

	mockValidationError := errors.New("validation error")
	mockValidator.ValidateStructFunc = func(obj interface{}) error {
		return mockValidationError
	}

	_, err := useCase.Execute(userID, carID, inputDTO)

	assert.Error(t, err, "Expected an error")
	assert.Equal(t, mockValidationError, err, "Expected validation error to match")

	mockRepo.AssertExpectations(t)
} */
