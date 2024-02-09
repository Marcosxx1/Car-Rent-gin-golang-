package usecases_test

import (
	"errors"
	"testing"

	usecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/car-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/car-controller/car-dtos"
	"github.com/stretchr/testify/assert"
)

// FindAllCars é um método do MockCarRepository que simula o comportamento
// do método correspondente na interface CarRepository.
// Ele recebe os parâmetros 'page' e 'pageSize' e retorna uma lista de carros
// e um possível erro.
//  func (m *MockCarRepository) FindAllCars(page, pageSize int) ([]*domain.Car, error)
// Chama o método 'Called' do mock para registrar a chamada com os argumentos fornecidos.
//	args := m.Called(page, pageSize)
// Verifica se o método foi chamado pelo menos uma vez (o primeiro argumento não é nulo).
//	if args.Get(0) != nil {
// Converte o primeiro argumento retornado (lista simulada de carros) para []*domain.Car.
// Retorna também o erro simulado (segundo argumento).
//		return args.Get(0).([]*domain.Car), args.Error(1)
//	}
// Se o método não foi chamado, retorna nil para a lista de carros simulada e o erro simulado.
//	return nil, args.Error(1)
//}
func (m *MockCarRepository) FindAllCars(page, pageSize int) ([]*domain.Car, error) {
	args := m.Called(page, pageSize)

	if args.Get(0) != nil {
		return args.Get(0).([]*domain.Car), args.Error(1)
	}

	return nil, args.Error(1)
}

func TestGetAllCarsUseCase(t *testing.T) {
	mockRepo := new(MockCarRepository)

	expectedCars := []*domain.Car{
		{
			ID:           "asfa4fa4a4f",
			Name:         "Car1",
			Description:  "Description1",
			DailyRate:    50.0,
			Available:    true,
			LicensePlate: "ABC123",
			FineAmount:   10.0,
			Brand:        "Brand1",
		},
	}

	mockRepo.On("FindAllCars", 1, 10).Return(expectedCars, nil)

	result, err := usecases.GetAllCarsUseCase(mockRepo, 1, 10)

	assert.NoError(t, err)

	expectedDTO := []*dtos.CarOutputDTO{
		{
			ID:           "asfa4fa4a4f",
			Name:         "Car1",
			Description:  "Description1",
			DailyRate:    50.0,
			Available:    true,
			LicensePlate: "ABC123",
			FineAmount:   10.0,
			Brand:        "Brand1",
		},
	}

	assert.Equal(t, expectedDTO, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAllCarsUseCase_Error(t *testing.T) {
	mockRepo := new(MockCarRepository)

	mockRepo.On("FindAllCars", 1, 10).Return(nil, errors.New("erro simulado"))

	result, err := usecases.GetAllCarsUseCase(mockRepo, 1, 10)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}