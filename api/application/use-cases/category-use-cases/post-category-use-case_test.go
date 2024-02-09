package usecases_test

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/stretchr/testify/mock"
)

type MockCategoryRepository struct {
	mock.Mock
}



func (m *MockCategoryRepository) FindCategoryByName(name string) (*domain.Category, error) {
	args := m.Called(name)
	return args.Get(0).(*domain.Category), args.Error(1)
}
func (m *MockCategoryRepository) GetAll() ([]*domain.Category, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Category), args.Error(1)
}
func (m *MockCategoryRepository) PostCategory(category *domain.Category) error {
	args := m.Called(category)
	return args.Error(0)
}
