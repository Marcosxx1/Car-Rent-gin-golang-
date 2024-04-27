package authautomocks

import (
	"github.com/stretchr/testify/mock"
)

// MockAuthRepository is a mock implementation of AuthRepository
type MockAuthRepository struct {
	mock.Mock
}

// GenerateAuthToken mocks the GenerateAuthToken method
func (m *MockAuthRepository) GenerateAuthToken(user_id string, user_role string) (string, error) {
	args := m.Called(user_id, user_role)
	return args.String(0), args.Error(1)
}

// ValidateAuthToken mocks the ValidateAuthToken method
func (m *MockAuthRepository) ValidateAuthToken(tokenString string) error {
	args := m.Called(tokenString)
	return args.Error(0)
}
