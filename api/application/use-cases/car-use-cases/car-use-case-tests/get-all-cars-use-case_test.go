package usecases_test

/*
import (
	"errors"
	"testing"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	m "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases/car-use-case-tests/mocks-and-structs"
	s "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/specification-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCarsUseCase_Success(t *testing.T) {
	// Arrange
	mockCarRepo := new(m.MockCarRepository)
	mockSpecRepo := new(s.MockSpecificationRepository)

	useCase := usecases.NewGetAllCarsUseCase(mockCarRepo, mockSpecRepo)

	page := 1
	pageSize := 10

	expectedCars := m.MockListOfCarsFromDatabase(5)
	mockCarRepo.On("FindAllCars", page, pageSize).Return(expectedCars, nil)

	for _, car := range expectedCars {
		mockSpecRepo.On("FindAllSpecificationsByCarId", car.ID).Return(m.MockSpecificationsFromDatabase(), nil)
	}

	// Act
	result, err := useCase.Execute(page, pageSize)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, len(expectedCars))

	for i, car := range expectedCars {
		assert.Equal(t, car.ID, result[i].ID)
		assert.Equal(t, car.Name, result[i].Name)
		mockSpecRepo.AssertExpectations(t)
	}

	mockCarRepo.AssertExpectations(t)
}

func TestGetAllCarsUseCase_ErrorFindingCars(t *testing.T) {
	// Arrange
	mockCarRepo := new(m.MockCarRepository)
	mockSpecRepo := new(s.MockSpecificationRepository)

	useCase := usecases.NewGetAllCarsUseCase(mockCarRepo, mockSpecRepo)

	page := 1
	pageSize := 10

	mockCarRepo.On("FindAllCars", page, pageSize).Return([]*domain.Car{}, errors.New("error finding cars"))

	// Act
	result, err := useCase.Execute(page, pageSize)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "error finding cars")

	mockCarRepo.AssertExpectations(t)
}

func TestGetAllCarsUseCase_ErrorFindingSpecifications(t *testing.T) {
	// Arrange
	mockCarRepo := new(m.MockCarRepository)
	mockSpecRepo := new(s.MockSpecificationRepository)

	useCase := usecases.NewGetAllCarsUseCase(mockCarRepo, mockSpecRepo)

	page := 1
	pageSize := 10

	expectedCars := m.MockListOfCarsFromDatabase(5)
	mockCarRepo.On("FindAllCars", page, pageSize).Return(expectedCars, nil)

	mockSpecRepo.On("FindAllSpecificationsByCarId", expectedCars[0].ID).Return([]*domain.Specification{}, errors.New("error finding specifications"))

	// Act
	result, err := useCase.Execute(page, pageSize)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "error finding specifications")

	mockCarRepo.AssertExpectations(t)
	mockSpecRepo.AssertExpectations(t)
}
*/