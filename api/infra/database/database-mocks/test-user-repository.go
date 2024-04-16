package databasemocks

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) PostUser(userData *domain.User) error {
	args := m.Called(userData)

	return args.Error(0)
}

func (m *MockUserRepository) GetById(id string) (*domain.User, error) {
	args := m.Called()

	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(email string) (*domain.User, error) {
	args := m.Called(email)

	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) Update(id string, data *domain.User) (*domain.User, error) {
	args := m.Called(id, data)

	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) UpdatePassword(id, newPassword string) error {
	args := m.Called(id, newPassword)

	return args.Error(0)
}
