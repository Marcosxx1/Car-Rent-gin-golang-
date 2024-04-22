package authautomocks

import (
	"github.com/stretchr/testify/mock"
)

type MockPasswordRepository struct {
	mock.Mock
}

func (m *MockPasswordRepository) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockPasswordRepository) VerifyPassword(plainPassword, hashedPassword string) bool {
	args := m.Called(plainPassword, hashedPassword)
	return args.Bool(0)
}

func (m *MockPasswordRepository) CompareHashedPassword(hashedPassword, password string) error {
	args := m.Called(hashedPassword, password)
	return args.Error(0)
}
