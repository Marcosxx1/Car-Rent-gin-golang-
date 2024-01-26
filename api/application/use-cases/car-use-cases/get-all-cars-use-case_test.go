package usecases_test

import (
	"testing"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/stretchr/testify/assert"
)

func (m *MockCarRepository) FindAllCars(int, int) ([]*domain.Car, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Car), args.Error(1)
}
func (m *MockCarRepository) UpdateCar(id string, car domain.Car) (*domain.Car, error) {
	args := m.Called(id, car)
	return args.Get(0).(*domain.Car), args.Error(1)
}

func TestFindAllUseCase_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockCarRepository)
	expectedCars := []*domain.Car{
		{ID: "1", Name: "Car1"},
		{ID: "2", Name: "Car2"},
	}

	mockRepo.On("FindAllCars").Return(expectedCars, nil)

	// Act
	resultCars, err := usecases.GetAllCarsUseCase(mockRepo, 1, 1)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resultCars)
	assert.Equal(t, expectedCars, resultCars)

	mockRepo.AssertExpectations(t)
}
