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

func TestErrorPostingSpecifications(t *testing.T) {
	mockCarRepo := new(databasemocks.MockCarRepository)
	mockSpecRepo := new(databasemocks.MockSpecificationRepository)

	expectedErr := fmt.Errorf("failed to create specification record")
	mockCarRepo.On("RegisterCar", mock.AnythingOfType("*domain.Car")).Return(nil)
	mockSpecRepo.On("PostMultipleSpecifications", mock.AnythingOfType("[]*domain.Specification")).Return(expectedErr)

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
	carID := "test_car_id"

	resultChan := make(chan *cardtos.CarOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool, 1)
	validationErrorSignal <- false

	var wg sync.WaitGroup
	wg.Add(1)

	go usecase.performCarCreation(&wg, resultChan, errorChan, validationErrorSignal, carID, inputDTO)

	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	select {
	case <-resultChan:
		t.Error("Expected an error but got a result")
	case err := <-errorChan:
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	}

	mockSpecRepo.AssertCalled(t, "PostMultipleSpecifications", mock.AnythingOfType("[]*domain.Specification"))
}

func TestSuccessfulCarCreation(t *testing.T) {
	mockCarRepo := new(databasemocks.MockCarRepository)
	mockSpecRepo := new(databasemocks.MockSpecificationRepository)

	mockCarRepo.On("RegisterCar", mock.AnythingOfType("*domain.Car")).Return(nil)

	mockSpecRepo.On("PostMultipleSpecifications", mock.AnythingOfType("[]*domain.Specification")).Return(nil)

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
	carID := "test_car_id"

	resultChan := make(chan *cardtos.CarOutputDTO)
	errorChan := make(chan error)
	validationErrorSignal := make(chan bool, 1)
	validationErrorSignal <- false

	var wg sync.WaitGroup
	wg.Add(1)

	go usecase.performCarCreation(&wg, resultChan, errorChan, validationErrorSignal, carID, inputDTO)

	go func() {
		wg.Wait()
		close(resultChan)
		close(errorChan)
	}()

	select {
	case createdCar := <-resultChan:
		assert.NoError(t, nil)

		assert.NotNil(t, createdCar)
		assert.Equal(t, inputDTO.Name, createdCar.Name)

	case err := <-errorChan:
		t.Errorf("Unexpected error: %v", err)
	}

	mockCarRepo.AssertCalled(t, "RegisterCar", mock.AnythingOfType("*domain.Car"))
	mockSpecRepo.AssertCalled(t, "PostMultipleSpecifications", mock.AnythingOfType("[]*domain.Specification"))
}
