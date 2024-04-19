package databasemocks

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/stretchr/testify/mock"
)

type MockCategoryRepository struct {
	mock.Mock
}

func (c *MockCategoryRepository) FindCategoryByName(name string) (*domain.Category, error) {
	args := c.Called(name)

	return args.Get(0).(*domain.Category), args.Error(1)
}

func (c *MockCategoryRepository) GetAll() ([]*domain.Category, error) {
	args := c.Called()

	return args.Get(0).([]*domain.Category), args.Error(1)
}

func (c *MockCategoryRepository) PostCategory(category *domain.Category) error {
	args := c.Called(category)

	return args.Error(0)
}
