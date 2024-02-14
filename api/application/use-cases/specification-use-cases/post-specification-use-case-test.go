package specificationusecases

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/stretchr/testify/mock"
)

type MockSpecificationRepository struct {
	mock.Mock
}

func (m *MockSpecificationRepository) FindSpecificationByName(name string) (*domain.Specification, error) {
	args := m.Called(name)
	return args.Get(0).(*domain.Specification), args.Error(1)
}

func (m *MockSpecificationRepository) GetAll() ([]*domain.Specification, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Specification), args.Error(1)
}

func (m *MockSpecificationRepository) PostSpecification(specification *domain.Specification) error {
	args := m.Called(specification)
	return args.Error(0)
}

func (m *MockSpecificationRepository) FindAllSpecificationsByCarId(carID string) ([]*domain.Specification, error) {
	args := m.Called(carID)
	return args.Get(0).([]*domain.Specification), args.Error(1)
}

func (m *MockSpecificationRepository) UpdateSpecification(carID string, specifications []*domain.Specification) ([]*domain.Specification, error) {
	args := m.Called(carID, specifications)
	return args.Get(0).([]*domain.Specification), args.Error(1)
}

func (m *MockSpecificationRepository) PostMultipleSpecifications(specifications []*domain.Specification) error {
	args := m.Called(specifications)
	return args.Error(0)
}
