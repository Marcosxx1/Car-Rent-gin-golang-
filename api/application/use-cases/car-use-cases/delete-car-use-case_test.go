package carusecases

import (
	"errors"
	"testing"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	databasemocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/database-mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteCarNoCarWithProvidedId(t *testing.T) {

	mockCarRepo := new(databasemocks.MockCarRepository)
	mockSpecficRepo := new(databasemocks.MockSpecificationRepository)

	mockCarRepo.On("FindCarById", mock.AnythingOfType("string")).Return((*domain.Car)(nil), errors.New("no car with provided id")) // na verdade é "no record found..."

	useCase := NewDeleteCarUseCase(mockCarRepo, mockSpecficRepo)
	err := useCase.Execute("1")

	assert.Error(t, err)
	assert.Equal(t, "no car with provided id", err.Error())
	mockCarRepo.AssertExpectations(t)

}

func TestDeleteCarNotFound(t *testing.T) {

	mockCarRepo := new(databasemocks.MockCarRepository)
	mockSpecficRepo := new(databasemocks.MockSpecificationRepository)

	mockCarRepo.On("FindCarById", mock.AnythingOfType("string")).Return((*domain.Car)(nil), nil)

	useCase := NewDeleteCarUseCase(mockCarRepo, mockSpecficRepo)
	err := useCase.Execute("1")

	assert.Error(t, err)
	assert.Equal(t, "car not found", err.Error())
	mockCarRepo.AssertExpectations(t)

}
func TestDeleteCarSuccess(t *testing.T) {

	mockCarRepo := new(databasemocks.MockCarRepository)
	mockSpecficRepo := new(databasemocks.MockSpecificationRepository)

	mockCarRepo.On("FindCarById", mock.AnythingOfType("string")).Return((&domain.Car{ID: "any_id"}), nil)
	mockCarRepo.On("DeleteCar", mock.AnythingOfType("string")).Return(nil)

	useCase := NewDeleteCarUseCase(mockCarRepo, mockSpecficRepo)
	err := useCase.Execute("1")

	assert.Nil(t, err)
	mockCarRepo.AssertExpectations(t)

}
