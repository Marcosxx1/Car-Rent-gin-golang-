package usecases_test

import "github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"

func (m *MockCarRepository) UpdateCar(id string, car domain.Car) (*domain.Car, error) {
	args := m.Called(id, car)
	return args.Get(0).(*domain.Car), args.Error(1)
}

//create a mock implementation