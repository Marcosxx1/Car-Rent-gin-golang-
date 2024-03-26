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

func TestFindCarByIdUseCase_Success(t *testing.T) {
	// Arrange
	mockCarRepo := new(m.MockCarRepository)
	mockSpecRepo := new(s.MockSpecificationRepository)

	useCase := usecases.NewFindCarByIdUseCase(mockCarRepo, mockSpecRepo)

	carID := "mockedID"

	expectedCar := m.MockCarFromDataBase()

	mockCarRepo.On("FindCarById", carID).Return(expectedCar, nil)
	mockSpecRepo.On("FindAllSpecificationsByCarId", carID).Return(m.MockSpecificationsFromDatabase(), nil)

	// Act
	result, err := useCase.Execute(carID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, carID, result.ID)
	assert.Equal(t, expectedCar.Name, result.Name)

	mockCarRepo.AssertExpectations(t)
	mockSpecRepo.AssertExpectations(t)
}

func TestFindCarByIdUseCase_CarNotFound(t *testing.T) {
	// Arrange
	mockCarRepo := new(m.MockCarRepository)
	mockSpecRepo := new(s.MockSpecificationRepository)

	useCase := usecases.NewFindCarByIdUseCase(mockCarRepo, mockSpecRepo)

	carID := "nonExistentID"

	mockCarRepo.On("FindCarById", carID).Return(&domain.Car{}, errors.New("car not found"))

	// Act
	result, err := useCase.Execute(carID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, errors.New("car not found"), err)

	mockCarRepo.AssertExpectations(t)
	mockSpecRepo.AssertExpectations(t)
}

func TestFindCarByIdUseCase_SpecificationNotFound(t *testing.T) {
	// Arrange
	mockCarRepo := new(m.MockCarRepository)
	mockSpecRepo := new(s.MockSpecificationRepository)

	useCase := usecases.NewFindCarByIdUseCase(mockCarRepo, mockSpecRepo)

	carID := "mockedID"

	expectedCar := m.MockCarFromDataBase()

	mockCarRepo.On("FindCarById", carID).Return(expectedCar, nil)
	mockSpecRepo.On("FindAllSpecificationsByCarId", carID).Return([]*domain.Specification{}, errors.New("specifications not found"))

	// Act
	result, err := useCase.Execute(carID)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "specifications not found")

	mockCarRepo.AssertExpectations(t)
	mockSpecRepo.AssertExpectations(t)
}
*/