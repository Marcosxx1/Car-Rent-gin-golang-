package usecases_test

import (
	"errors"
	"testing"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	spec "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/specification-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

type MockCarRepository struct {
	mock.Mock
}

func TestDeleteCarUseCase_Execute(t *testing.T) {
	// Arrange
	mockCarRepo := new(MockCarRepository)
	mockSpecRepo := new(spec.MockSpecificationRepository)

	useCase := usecases.NewDeleteCarUseCase(mockCarRepo, mockSpecRepo)

	existingCarID := "123"
	expectedError := errors.New("some error")

	// Mock the behavior of the repositories
	//mockCarRepo.On("FindCarById", existingCarID).Return(nil, expectedError)
	mockCarRepo.On("FindCarById", existingCarID).Return(&domain.Car{}, expectedError)

	// Add more mocks as needed

	// Act
	err := useCase.Execute(existingCarID)

	// Assert
	assert.Equal(t, expectedError, err)

	// Assert that the expected methods were called
	mockCarRepo.AssertExpectations(t)
	mockSpecRepo.AssertExpectations(t)
}

func (m *MockCarRepository) RegisterCar(car *domain.Car) error {
	args := m.Called(car)
	return args.Error(0)
}

func (m *MockCarRepository) FindCarByLicensePlate(licensePlate string) (*domain.Car, error) {
	args := m.Called(licensePlate)
	return args.Get(0).(*domain.Car), args.Error(1)
}

func (m *MockCarRepository) FindAllCars(page, pageSize int) ([]*domain.Car, error) {
	args := m.Called(page, pageSize)
	return args.Get(0).([]*domain.Car), args.Error(1)
}

func (m *MockCarRepository) DeleteCar(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockCarRepository) FindCarById(id string) (*domain.Car, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Car), args.Error(1)
}

func (m *MockCarRepository) UpdateCar(id string, car *domain.Car) (*domain.Car, error) {
	args := m.Called(id, car)
	return args.Get(0).(*domain.Car), args.Error(1)
}
