package usecases_test

import (
	"errors"
	"testing"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	m "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases/car-use-case-tests/mocks-and-structs"
	s "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/specification-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/stretchr/testify/assert"
)

/*
A quick note since I'm begining to test again with the refactored structure:

In a unit test, the primary goal is to verify that the unit under test (in this case, the DeleteCarUseCase)
behaves as expected and interacts correctly with its dependencies.
We want to ensure that the methods are being called with the correct arguments and in the correct order.

The actual logic of the methods in the dependencies (like FindCarById and DeleteCar) should be tested separately
in their own unit tests. The purpose of the DeleteCarUseCase unit test is to check if it correctly orchestrates
the calls to these methods and handles their results appropriately.
*/

func TestDeleteCarUseCase_Success(t *testing.T) {
	// Arrange
	mockCarRepo := new(m.MockCarRepository)
	mockSpecRepo := new(s.MockSpecificationRepository)

	useCase := usecases.NewDeleteCarUseCase(mockCarRepo, mockSpecRepo)

	existingCarID := "123"

	// Mock the behavior of the repositories
	mockCarRepo.On("FindCarById", existingCarID).Return(&domain.Car{}, nil)
	mockCarRepo.On("DeleteCar", existingCarID).Return(nil) // Success case

	// Act
	err := useCase.Execute(existingCarID)

	// Assert
	assert.NoError(t, err)
	mockCarRepo.AssertExpectations(t)
	mockSpecRepo.AssertExpectations(t)
}

func TestDeleteCarUseCase_CarNotFound(t *testing.T) {
	// Arrange
	mockCarRepo := new(m.MockCarRepository)
	mockSpecRepo := new(s.MockSpecificationRepository)

	useCase := usecases.NewDeleteCarUseCase(mockCarRepo, mockSpecRepo)

	notFoundCarID := "456"

	// Act
	mockCarRepo.On("FindCarById", notFoundCarID).Return(&domain.Car{}, errors.New("car not found"))
	err := useCase.Execute(notFoundCarID)

	// Assert
	assert.Equal(t, errors.New("car not found"), err)
	mockCarRepo.AssertExpectations(t)
	mockSpecRepo.AssertExpectations(t)
}

func TestDeleteCarUseCase_DeleteCarFailure(t *testing.T) {
	// Arrange
	mockCarRepo := new(m.MockCarRepository)
	mockSpecRepo := new(s.MockSpecificationRepository)

	useCase := usecases.NewDeleteCarUseCase(mockCarRepo, mockSpecRepo)

	existingCarID := "789"
	expectedError := errors.New("some error")

	mockCarRepo.On("FindCarById", existingCarID).Return(&domain.Car{}, nil)
	mockCarRepo.On("DeleteCar", existingCarID).Return(expectedError) // Deletion failure

	// Act
	err := useCase.Execute(existingCarID)

	// Assert
	assert.Equal(t, expectedError, err)
	mockCarRepo.AssertExpectations(t)
	mockSpecRepo.AssertExpectations(t)
}
