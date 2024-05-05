package carusecases

import (
	"fmt"
	"sync"
	"testing"

	cardtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/car"
	databasemocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/database-mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestValidationError(t *testing.T) {
	mockCarRepo := new(databasemocks.MockCarRepository)

	mockSpecRepo := new(databasemocks.MockSpecificationRepository)

	usecase := NewPostCarUseCase(mockCarRepo, mockSpecRepo)

	inputDTO := &cardtos.CarInputDTO{}

	car, err := usecase.ExecuteConcurrently(inputDTO)

	assert.Error(t, err)
	assert.Nil(t, car)

	mockCarRepo.AssertNotCalled(t, "RegisterCar")
	mockSpecRepo.AssertNotCalled(t, "PostMultipleSpecifications")
}

func TestErrorRegisteringCar(t *testing.T) {
	mockCarRepo := new(databasemocks.MockCarRepository)
	mockSpecRepo := new(databasemocks.MockSpecificationRepository)

	mockCarRepo.On("RegisterCar", mock.AnythingOfType("*domain.Car")).Return(fmt.Errorf("failed to create car record"))

	usecase := NewPostCarUseCase(mockCarRepo, mockSpecRepo)

	inputDTO := &cardtos.CarInputDTO{
		Name:         "Test Car",
		Description:  "Test Description",
		DailyRate:    50.0,
		Available:    true,
		LicensePlate: "ABC123",
		FineAmount:   20.0,
		Brand:        "Test Brand",
	}
	resultChan := make(chan *cardtos.CarOutputDTO)

	errorChan := make(chan error)
	validationErrorSignal := make(chan bool, 1)
	validationErrorSignal <- false

	var wg sync.WaitGroup
	wg.Add(1)
	go usecase.performCarCreation(&wg, resultChan, errorChan, validationErrorSignal, "id", inputDTO)

	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()
	select {
	case <-resultChan:
		t.Error("failed to create car record: failed to create car record")
	case err := <-errorChan:
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create car record: failed to create car record") // Assert specific error message
	}
}

/*func TestCarCreationError(t *testing.T) {
	mockCarRepo := new(databasemocks.MockCarRepository)
	mockCarRepo.On("RegisterCar", mock.AnythingOfType("*domain.Car")).Return(errors.New("failed to create car"))

	mockSpecRepo := new(databasemocks.MockSpecificationRepository)

	usecase := NewPostCarUseCase(mockCarRepo, mockSpecRepo)

	inputDTO := &cardtos.CarInputDTO{
		Name:         "Test Car",
		Description:  "Test Description",
		DailyRate:    50.0,
		Available:    true,
		LicensePlate: "ABC123",
		FineAmount:   20.0,
		Brand:        "Test Brand",
		CategoryID:   "categoryID",
	}

	car, err := usecase.ExecuteConcurrently(inputDTO)

	assert.Error(t, err)
	assert.Nil(t, car)

	mockCarRepo.AssertExpectations(t)
	mockSpecRepo.AssertNotCalled(t, "PostMultipleSpecifications")
}

 func TestSpecificationCreationError(t *testing.T) {
	// Mock Car Repository
	mockCarRepo := new(databasemocks.MockCarRepository)
	mockCarRepo.On("RegisterCar", mock.AnythingOfType("*domain.Car")).Return(nil) // Mock successful car creation

	// Mock Specification Repository to simulate specification creation error
	mockSpecRepo := new(databasemocks.MockSpecificationRepository)
	mockSpecRepo.On("PostMultipleSpecifications", mock.AnythingOfType("[]*domain.Specification")).Return(errors.New("failed to create specifications"))

	// Create the use case instance
	usecase := NewPostCarUseCase(mockCarRepo, mockSpecRepo)

	// Create the input DTO
	inputDTO := &cardtos.CarInputDTO{
		Name:         "Test Car",
		Description:  "Test Description",
		DailyRate:    50.0,
		Available:    true,
		LicensePlate: "ABC123",
		FineAmount:   20.0,
		Brand:        "Test Brand",
		CategoryID:   "categoryID",
		// Add specifications if needed
	}

	// Execute the use case
	car, err := usecase.ExecuteConcurrently(inputDTO)

	// Assertions
	assert.Error(t, err)   // An error should occur due to specification creation failure
	assert.Nil(t, car)     // No car should be returned
	// Add more assertions as needed

	// Verify expectations for the mock interactions
	mockCarRepo.AssertExpectations(t)   // Car creation should be attempted
	mockSpecRepo.AssertExpectations(t)  // Specification creation should be attempted
}
*/
