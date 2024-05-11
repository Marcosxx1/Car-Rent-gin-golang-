package carusecases

import (
	"testing"

	cardtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/car"
	databasemocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/database-mocks"
	"github.com/stretchr/testify/assert"
)

func TestValidationErrorForPut(t *testing.T) {

	// Verificar, está passando mas não está pegando os erros de validação
	mockCarRepo := new(databasemocks.MockCarRepository)

	mockSpecRepo := new(databasemocks.MockSpecificationRepository)
	useCase := NewPostCarUseCase(mockCarRepo, mockSpecRepo)

	inputDTO := &cardtos.CarInputDTO{}

	result, err := useCase.ExecuteConcurrently(inputDTO)

	assert.Error(t, err)
	assert.Nil(t, result)
}
