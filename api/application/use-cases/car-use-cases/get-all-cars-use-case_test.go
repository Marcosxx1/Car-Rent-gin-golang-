package carusecases

import (
	"fmt"
	"testing"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	databasemocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/database-mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllCarsErrorFindingCarForPageAndPageSize(t *testing.T) {

	mockCarRepo := new(databasemocks.MockCarRepository)
	mockSpecifcRepo := new(databasemocks.MockSpecificationRepository)

	mockCarRepo.On("FindAllCars", 1, 10).Return(([]*domain.Car)(nil), fmt.Errorf("error finding cars for page %d, pageSize %d", 1, 10))

	usecase := NewGetAllCarsUseCase(mockCarRepo, mockSpecifcRepo)
	_, err := usecase.Execute(1, 10)

	assert.Error(t, err)
	assert.Equal(t, "error finding cars for page 1, pageSize 10", err.Error())
	mockCarRepo.AssertExpectations(t)
	mockSpecifcRepo.AssertExpectations(t)
	mockCarRepo.AssertNumberOfCalls(t, "FindAllCars", 1)

}

func TestGetAllCarsNoCarsFoundForPageAndPageSize(t *testing.T) {

	mockCarRepo := new(databasemocks.MockCarRepository)
	mockSpecifcRepo := new(databasemocks.MockSpecificationRepository)

	mockCarRepo.On("FindAllCars", 1, 10).Return([]*domain.Car{}, nil)

	usecase := NewGetAllCarsUseCase(mockCarRepo, mockSpecifcRepo)
	allCars, err := usecase.Execute(1, 10)

	assert.Error(t, err)
	assert.Nil(t, allCars)
	assert.Equal(t, "no cars found for page 1, pageSize 10", err.Error())
	mockCarRepo.AssertExpectations(t)
	mockSpecifcRepo.AssertExpectations(t)

}

func TestGetAllCarsErrorFetchingSpecifications(t *testing.T) {
	mockCarRepo := new(databasemocks.MockCarRepository)
	mockSpecifcRepo := new(databasemocks.MockSpecificationRepository)

	mockCarRepo.On("FindAllCars", 1, 10).Return([]*domain.Car{{ID: "carID1"}}, nil)

	mockSpecifcRepo.On("FindAllSpecificationsByCarId", "carID1").Return(([]*domain.Specification)(nil), fmt.Errorf("error fetching specifications"))

	usecase := NewGetAllCarsUseCase(mockCarRepo, mockSpecifcRepo)
	allCars, err := usecase.Execute(1, 10)

	assert.Error(t, err)
	assert.Nil(t, allCars)
	assert.Equal(t, "failed to retrieve specifications for car carID1: error fetching specifications", err.Error())

	mockCarRepo.AssertExpectations(t)
	mockSpecifcRepo.AssertExpectations(t)
}

func TestGetAllCarsSuccess(t *testing.T) {
	mockCarRepo := new(databasemocks.MockCarRepository)
	mockCarRepo.On("FindAllCars", 1, 10).Return([]*domain.Car{
		{
			ID:           "carID1",
			Name:         "Car 1",
			Description:  "Description 1",
			DailyRate:    50,
			Available:    true,
			LicensePlate: "ABC123",
			FineAmount:   20,
			Brand:        "Brand 1",
			CategoryID:   "categoryID1",
		},
		{
			ID:           "carID2",
			Name:         "Car 2",
			Description:  "Description 2",
			DailyRate:    60,
			Available:    true,
			LicensePlate: "XYZ789",
			FineAmount:   30,
			Brand:        "Brand 2",
			CategoryID:   "categoryID2",
		},
	}, nil)

	mockSpecRepo := new(databasemocks.MockSpecificationRepository)
	mockSpecRepo.On("FindAllSpecificationsByCarId", mock.Anything).Return([]*domain.Specification{
		{
			ID:          "specID1",
			Name:        "Spec 1",
			Description: "Spec Description 1",
		},
		{
			ID:          "specID2",
			Name:        "Spec 2",
			Description: "Spec Description 2",
		},
	}, nil)

	usecase := NewGetAllCarsUseCase(mockCarRepo, mockSpecRepo)

	allCars, err := usecase.Execute(1, 10)

	assert.NoError(t, err)
	assert.NotNil(t, allCars)
	assert.Equal(t, 2, len(allCars))

	mockCarRepo.AssertExpectations(t)
	mockSpecRepo.AssertExpectations(t)
}
