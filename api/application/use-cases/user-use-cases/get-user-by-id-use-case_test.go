package userusecases

import (
	"errors"
	"testing"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	databasemocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/database-mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserByIdSuccess(t *testing.T) {
	mockRepo := new(databasemocks.MockUserRepository)
	useCase := NewGetUserByIdUseCase(mockRepo)

	mockUser := &domain.User{
		ID:          "any_id",
		Name:        "any_name",
		Email:       "any_email@any.com",
		Password:    "any_lengthy_password",
		OldPassword: "any_old_lengthy_password",
		Role:        "admin",
		Status:      true,
		Avatar:      "any_avatar",
		CreatedAt:   "any_creation_date",
	}

	mockRepo.On("GetById", mock.AnythingOfType("string")).Return(mockUser, nil)

	existUser, err := useCase.Execute("any_id")

	assert.Nil(t, err, "Expected error to be nil, it wasn't.")
	assert.NotNil(t, existUser, "Expected return of Execute to not be nil, it was.")

	assert.Equal(t, mockUser.ID, existUser.ID, "Expected user ID to match")
	assert.Equal(t, mockUser.Name, existUser.Name, "Expected user Name to match")
	assert.Equal(t, mockUser.Email, existUser.Email, "Expected user Email to match")
	assert.Equal(t, mockUser.Status, existUser.Status, "Expected user Status to match")
	assert.Equal(t, mockUser.Avatar, existUser.Avatar, "Expected user Avatar to match")
	assert.Equal(t, mockUser.CreatedAt, existUser.CreatedAt, "Expected user CreatedAt to match")
}

func TestGetUserByIdUserNotFound(t *testing.T) {
	mockRepo := new(databasemocks.MockUserRepository)
	useCase := NewGetUserByIdUseCase(mockRepo)

	mockRepo.On("GetById", mock.AnythingOfType("string")).Return((*domain.User)(nil), nil) // User not found

	existUser, err := useCase.Execute("non_existing_id")

	assert.NotNil(t, err, "Expected error, user not found")
	assert.Nil(t, existUser, "Expected user to be nil")
}

func TestGetUserByIdError(t *testing.T) {
	mockRepo := new(databasemocks.MockUserRepository)
	useCase := NewGetUserByIdUseCase(mockRepo)

	mockRepo.On("GetById", mock.AnythingOfType("string")).Return((*domain.User)(nil), errors.New("some error")) // Simulate error

	existUser, err := useCase.Execute("any_id")

	assert.NotNil(t, err, "Expected error, got nil")
	assert.Nil(t, existUser, "Expected user to be nil")
}
