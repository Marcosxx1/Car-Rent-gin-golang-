package usecases_test

import (
	"errors"
	"testing"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller/car-dtos"
	"github.com/stretchr/testify/assert"
)



func (m *MockCarRepository) FindCarById(id string) (*domain.Car, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Car), args.Error(1)
}

func TestFindCarById(t *testing.T) {
	assert := assert.New(t)
	mockRepo := new(MockCarRepository)

	inputCar := &domain.Car{
		ID:           "1",
		Name:         "Test Car",
		Description:  "Test Description",
		DailyRate:    50.0,
		Available:    true,
		LicensePlate: "ABC123",
		FineAmount:   10.0,
		Brand:        "Test Brand",
	}
	mockRepo.On("FindCarById", "1").Return(inputCar, nil)

	carToBeReturned := &dtos.CarOutputDTO{
		ID:           "1",
		Name:         inputCar.Name,
		Description:  inputCar.Description,
		DailyRate:    inputCar.DailyRate,
		Available:    inputCar.Available,
		LicensePlate: inputCar.LicensePlate,
		FineAmount:   inputCar.FineAmount,
		Brand:        inputCar.Brand,
		CreatedAt:    inputCar.CreatedAt,
	}


	carToBeFound, err := usecases.GetCarByIdUseCase("1", mockRepo)

	assert.Nil(err)
	assert.EqualValues(carToBeReturned, carToBeFound)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNumberOfCalls(t, "FindCarById", 1)
}

func TestFindCarByIdWithError(t *testing.T) {
	assert := assert.New(t)
	mockRepo := new(MockCarRepository)

	mockRepo.On("FindCarById", "1").Return(&domain.Car{}, errors.New("erro simulado"))

	_, err := usecases.GetCarByIdUseCase("1", mockRepo)

	assert.Equal(errors.New("erro simulado"), err)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNumberOfCalls(t, "FindCarById", 1)
}