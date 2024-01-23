package usecases_test

import (
	"testing"
	"time"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/car-controller/car-dtos"
	"github.com/stretchr/testify/assert"
)

func TestFindCarById(t *testing.T) {
	assert := assert.New(t)
	mockRepo := new(MockCarRepository)

	inputCar := &domain.Car{
			Name:         "Test Car",
			Description:  "Test Description",
			DailyRate:    50.0,
			Available:    true,
			LicensePlate: "ABC123",
			FineAmount:   10.0,
			Brand:        "Test Brand",
			CategoryId:   "123",
	}

	carToBeReturned := &dtos.CarOutputDTO{
			Id:           "1",
			Name:         inputCar.Name,
			Description:  inputCar.Description,
			DailyRate:    inputCar.DailyRate,
			Available:    inputCar.Available,
			LicensePlate: inputCar.LicensePlate,
			FineAmount:   inputCar.FineAmount,
			Brand:        inputCar.Brand,
			CategoryId:   inputCar.CategoryId,
			CreatedAt:    time.Now(),
	}

	mockRepo.On("FindCarById", "1").Return(carToBeReturned, nil)

	carToBeFound, err := usecases.GetCarByIdUseCase("1", mockRepo)

	assert.Nil(err)
	assert.EqualValues(carToBeReturned, carToBeFound)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNumberOfCalls(t, "FindCarById", 1)
}

