package databasemocks

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/stretchr/testify/mock"
)

type MockCarRepository struct {
	mock.Mock
}

func (m *MockCarRepository) RegisterCar(car *domain.Car) error {
	args := m.Called(car)
	return args.Error(0)
}

func (m *MockCarRepository) UpdateAvailableCar(id string, available bool) error {
	args := m.Called(id, available)
	return args.Error(0)
}

func (m *MockCarRepository) FindAvailableCarById(id string) *domain.Car {
	args := m.Called(id)
	return args.Get(0).(*domain.Car)
}

func (m *MockCarRepository) FindAvailableCars(brand string, categoryID string, name string) (string, string, string, []*domain.Car) {
	args := m.Called(brand, categoryID, name)
	return args.String(0), args.String(1), args.String(2), args.Get(3).([]*domain.Car)
}

func (m *MockCarRepository) FindCarByLicensePlate(licensePlate string) (*domain.Car, error) {
	args := m.Called(licensePlate)
	return args.Get(0).(*domain.Car), args.Error(1)
}

func (m *MockCarRepository) FindAllCars(page, pageSize int) ([]*domain.Car, error) {
	args := m.Called(page, pageSize)
	return args.Get(0).([]*domain.Car), args.Error(1)
}

func (m *MockCarRepository) DeleteCar(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockCarRepository) FindCarById(id string) (*domain.Car, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Car), args.Error(1)
}

func (m *MockCarRepository) UpdateCar(id string, car *domain.Car) (*domain.Car, error) {
	args := m.Called(id, car)
	return args.Get(0).(*domain.Car), args.Error(1)
}

func (m *MockCarRepository) AlterCarStatus(id string, available bool) error {
	args := m.Called(id, available)
	return args.Error(0)
}
