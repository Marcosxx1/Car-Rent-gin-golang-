package reviewusecases

import (
	"errors"
	"testing"

	testutils "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/utils/test-utils"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	databasemocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/database-mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAllReviews(t *testing.T) {
	// Setup mock data
	mockResultDomain := testutils.GenerateDomainReviews(4)
	//mockResultDto := testutils.GenerateReviews(4)

	// Setup mock repository
	mockRepo := new(databasemocks.MockReviewRepository)

	// Setup mock repository behavior
	mockRepo.On("GetAllReviews").Return(mockResultDomain, nil)

	// Create use case with mock repository
	useCase := NewGetAllReviewsUseCase(mockRepo)

	// Execute use case
	reviewsOutput, err := useCase.Execute()

	// Assertions
	assert.NoError(t, err, "Expected no error, but got an error")
	assert.NotNil(t, reviewsOutput, "Expected reviewsOutput to be not nil")

	mockRepo.AssertExpectations(t)
}

func TestGetAllReviewsNotFound(t *testing.T) {
	// Setup mock data
	mockResultDomain := testutils.GenerateDomainReviews(0)
	//mockResultDto := testutils.GenerateReviews(4)

	// Setup mock repository
	mockRepo := new(databasemocks.MockReviewRepository)

	// Setup mock repository behavior
	mockRepo.On("GetAllReviews").Return(mockResultDomain, nil)

	// Create use case with mock repository
	useCase := NewGetAllReviewsUseCase(mockRepo)

	// Execute use case
	_, err := useCase.Execute()

	// Assertions
	assert.Error(t, err, "Expected error, but no error got returned")
	assert.Contains(t, err.Error(), "no reviews found", "Expected error message to contain 'no reviews found'")

	mockRepo.AssertExpectations(t)
}

func TestGetAllReviewsErrorRetrievingReviews(t *testing.T) {
	mockRepo := new(databasemocks.MockReviewRepository)
	useCase := NewGetAllReviewsUseCase(mockRepo)

	mockRepo.On("GetAllReviews").Return(([]*domain.Reviews)(nil), errors.New("error retrieving reviews"))

	_, err := useCase.Execute()

	assert.Error(t, err, "Expected error, but none returned")
	assert.Contains(t, err.Error(), "error retrieving reviews", "Expected error message to contain 'error retrieving reviews'")

	mockRepo.AssertExpectations(t)

}

/*
// errado:
	mockRepo.On("GetAllReviews").Return((nil), errors.New("error retrieving reviews"))

	vai causar:
// --- FAIL: TestGetAllReviewsErrorRetrievingReviews (0.00s)
// panic: interface conversion: interface {} is nil, not []*domain.Reviews [recovered]
// 	panic: interface conversion: interface {} is nil, not []*domain.Reviews

// certo:
	mockRepo.On("GetAllReviews").Return(([]*domain.Reviews)(nil), errors.New("error retrieving reviews"))


func TestGetAllReviewsErrorRetrievingReviews(t *testing.T){
	//mockRepo.On("FindSpecificationByName", mock.AnythingOfType("string")).Return((*domain.Specification)(nil), errors.New("error querying specification"))

	mockRepo := new(databasemocks.MockReviewRepository)
	useCase := NewGetAllReviewsUseCase(mockRepo)

	mockRepo.On("GetAllReviews").Return((nil), errors.New("error retrieving reviews"))

	_, err := useCase.Execute()
	assert.Error(t, err, "Expected error, but no got returned")
	assert.Contains(t, err.Error(), "error retrieving reviews", "Expected error message to contain 'error retrieving reviews'")
} */
