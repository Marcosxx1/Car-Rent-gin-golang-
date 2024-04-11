package testingmocks

import (
	"errors"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

type MockUserRepository struct{}

func (m *MockUserRepository) PostUser(userData *domain.User) error {
	return nil
}

func (m *MockUserRepository) GetById(id string) (*domain.User, error) {
	if id == "123" {
		return &domain.User{ID: "123", Email: "mock@example.com", Password: "mock_password", Role: "user"}, nil
	}
	return nil, errors.New("user not found")
}

func (m *MockUserRepository) FindByEmail(email string) (*domain.User, error) {
	if email == "mock@example.com" {
		return &domain.User{ID: "123", Email: "mock@example.com", Password: "mock_password", Role: "user"}, nil
	}
	return nil, errors.New("user not found")
}

func (m *MockUserRepository) Update(id string, data *domain.User) (*domain.User, error) {
	if id == "123" {
		return &domain.User{ID: "123", Email: data.Email, Password: data.Password, Role: data.Role}, nil
	}
	return nil, errors.New("user not found")
}

func (m *MockUserRepository) UpdatePassword(id, newPassword string) error {
	return nil
}
