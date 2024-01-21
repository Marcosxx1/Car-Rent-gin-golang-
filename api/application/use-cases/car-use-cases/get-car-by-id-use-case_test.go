package usecases_test

import (
	"testing"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/stretchr/testify/assert"
)

func TestFindCarById(t *testing.T) {

	assert := assert.New(t)
	mockRepo := new(MockCarRepository)

	expectedCar := &domain.Car{
		Name:         "Test Car",
		Description:  "Test Description",
		DailyRate:    50.0,
		Available:    true,
		LicensePlate: "ABC123",
		FineAmount:   10.0,
		Brand:        "Test Brand",
		CategoryId:   "123",
	}

	mockRepo.On("FindCarById", "1").Return(expectedCar, nil)

	carToBeFound, err := usecases.GetCarByIdUseCase("1", mockRepo)
	
	assert.Nil(err)
	assert.Equal(expectedCar, carToBeFound)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNumberOfCalls(t, "FindCarById", 1)

}
