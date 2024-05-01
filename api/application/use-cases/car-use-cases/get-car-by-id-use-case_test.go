package carusecases

import (
	"errors"
	"testing"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	databasemocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/database-mocks"
	"github.com/stretchr/testify/assert"
)

func TestErrorFindingCarById(t *testing.T) {
	mockCarRepo := new(databasemocks.MockCarRepository)
	mockSpecRepo := new(databasemocks.MockSpecificationRepository)

	mockCarRepo.On("FindCarById", "nonexistentID").Return((*domain.Car)(nil), errors.New("car not found"))

	usecase := NewFindCarByIdUseCase(mockCarRepo, mockSpecRepo)
	car, err := usecase.Execute("nonexistentID")

	assert.Error(t, err)
	assert.Empty(t, car)
	assert.Equal(t, "car not found", err.Error())

	mockCarRepo.AssertExpectations(t)
	mockSpecRepo.AssertNotCalled(t, "FindAllSpecificationsByCarId")
}

func TestErrorFindingCarByIdEmptyDomainCar(t *testing.T) {
	mockCarRepo := new(databasemocks.MockCarRepository)
	mockSpecRepo := new(databasemocks.MockSpecificationRepository)

	mockCarRepo.On("FindCarById", "carID").Return(&domain.Car{}, nil)

	usecase := NewFindCarByIdUseCase(mockCarRepo, mockSpecRepo)
	car, err := usecase.Execute("carID")

	assert.Error(t, err)
	assert.Nil(t, car)
	assert.Equal(t, "car not found", err.Error())
	mockCarRepo.AssertExpectations(t)
	mockSpecRepo.AssertExpectations(t)
}

func TestErrorFindingCarsSpecification(t *testing.T) {
	mockCarRepo := new(databasemocks.MockCarRepository)
	mockSpecRepo := new(databasemocks.MockSpecificationRepository)

	mockCarRepo.On("FindCarById", "any_id").Return(&domain.Car{ID: "any_id"}, nil)
	mockSpecRepo.On("FindAllSpecificationsByCarId", "any_id").Return(([]*domain.Specification)(nil), errors.New("specifications not found"))

	usecase := NewFindCarByIdUseCase(mockCarRepo, mockSpecRepo)
	car, err := usecase.Execute("any_id")

	assert.Error(t, err)
	assert.Nil(t, car)
	assert.Equal(t, "specifications not found", err.Error())
	mockCarRepo.AssertExpectations(t)
	mockSpecRepo.AssertExpectations(t)
}

func TestGetCarByIdSuccess(t *testing.T) {
	mockCarRepo := new(databasemocks.MockCarRepository)
	mockCar := &domain.Car{
		ID:           "carID",
		Name:         "Test Car",
		Description:  "Test Description",
		DailyRate:    50.0,
		Available:    true,
		LicensePlate: "ABC123",
		FineAmount:   20.0,
		Brand:        "Test Brand",
	}
	mockCarRepo.On("FindCarById", "carID").Return(mockCar, nil)

	mockSpecRepo := new(databasemocks.MockSpecificationRepository)
	mockSpec := []*domain.Specification{
		{
			ID:          "specID",
			Name:        "Test Specification",
			Description: "Test Specification Description",
		},
	}
	mockSpecRepo.On("FindAllSpecificationsByCarId", "carID").Return(mockSpec, nil)

	usecase := NewFindCarByIdUseCase(mockCarRepo, mockSpecRepo)

	car, err := usecase.Execute("carID")

	assert.NoError(t, err)
	assert.NotNil(t, car)
	assert.Equal(t, "carID", car.ID)
	assert.Equal(t, "Test Car", car.Name)

	mockCarRepo.AssertExpectations(t)
	mockSpecRepo.AssertExpectations(t)
}
