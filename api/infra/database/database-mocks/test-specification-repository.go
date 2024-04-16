package databasemocks

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/stretchr/testify/mock"
)

// MockSpecificationRepository is a mock implementation of SpecificationRepository
type MockSpecificationRepository struct {
	mock.Mock
}

/*
When a method returns multiple values, we can retrieve them from the args object using the corresponding index:

args.Get(0) will retrieve the first return value.
args.Error(1) will retrieve the second return value, which is the error.
So, for a function with the signature *domain.Specification, error:

args.Get(0) will give us the *domain.Specification.
args.Error(1) will give us the error.
And if we have more return values, we can continue in the same pattern:

args.Get(1) will give us the second return value.
args.Get(2) will give us the third return value, and so on.
*/

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

func (m *MockSpecificationRepository) UpdateSpecification(car_id string, specification []*domain.Specification) ([]*domain.Specification, error) {
	args := m.Called(car_id, specification)

	return args.Get(0).([]*domain.Specification), args.Error(1)
}

func (m *MockSpecificationRepository) PostMultipleSpecifications(specifications []*domain.Specification) error {
	args := m.Called(specifications)

	return args.Error(0)
}
