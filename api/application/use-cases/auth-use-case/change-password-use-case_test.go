package authusecase

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	authautomocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/auth-autho/auth-auto-mocks"
	databasemocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/database-mocks"
	"github.com/stretchr/testify/assert"
)

func TestChangePasswordFailedToRetrieveUser(t *testing.T) {
	userRepository := new(databasemocks.MockUserRepository)
	passwordRepository := new(authautomocks.MockPasswordRepository)
	mockUseCase := NewChangePasswordUseCase(userRepository, passwordRepository)

	userRepository.On("GetById", "1").Return((*domain.User)(nil), fmt.Errorf("failed to retrieve user")).Once()

	id := "1"
	currentPassword := "password"
	newPassword := "newpassword"

	err := mockUseCase.Execute(id, currentPassword, newPassword)

	assert.NotNil(t, err)
	assert.Equal(t, "failed to retrieve user: failed to retrieve user", err.Error())
}

func TestChangePasswordUserNotFound(t *testing.T) {
	userRepository := new(databasemocks.MockUserRepository)
	passwordRepository := new(authautomocks.MockPasswordRepository)
	mockUseCase := NewChangePasswordUseCase(userRepository, passwordRepository)

	mockUser := &domain.User{}

	userRepository.On("GetById", "any_id").Return(mockUser, nil).Once()

	err := mockUseCase.Execute("any_id", "any_lengthy_password", "new_any_lengthy_password")

	assert.NotNil(t, err)
	assert.Equal(t, "user not found", err.Error())
	assert.True(t, errors.Is(err, err), "user not found")

}

	func TestChangePasswordCurrentPasswordDoesNotMatch(t *testing.T) {
		userRepository := new(databasemocks.MockUserRepository)
		passwordRepository := new(authautomocks.MockPasswordRepository)
		mockUseCase := NewChangePasswordUseCase(userRepository, passwordRepository)

		passwordRepository.On("VerifyPassword", "any_incorrect_password", "any_lengthy_password").Return(false).Once()

		mockUser := &domain.User{
			ID:       "any_id",
			Password: "any_lengthy_password",
		}

		userRepository.On("GetById", "any_id").Return(mockUser, nil).Once()

		err := mockUseCase.Execute("any_id", "any_incorrect_password", "new_any_lengthy_password")

		assert.NotNil(t, err)
		assert.Equal(t, "current password does not match", err.Error())
	}


	func TestChangePasswordFailedToHashNewPassword(t *testing.T) {
		userRepository := new(databasemocks.MockUserRepository)
		passwordRepository := new(authautomocks.MockPasswordRepository)
		mockUseCase := NewChangePasswordUseCase(userRepository, passwordRepository)

		passwordRepository.On("VerifyPassword", "any_incorrect_password", "any_lengthy_password").Return(true).Once()
		passwordRepository.On("HashPassword", "new_any_lengthy_password").Return("", fmt.Errorf("failed to hash new password")).Once()
		mockUser := &domain.User{
			ID:       "any_id",
			Password: "any_lengthy_password",
		}

		userRepository.On("GetById", "any_id").Return(mockUser, nil).Once()

		passwordRepository.On("HashPassword", "new_any_lengthy_password").Return("", fmt.Errorf("failed to hash new password")).Once()

		err := mockUseCase.Execute("any_id", "any_incorrect_password", "new_any_lengthy_password")

		assert.NotNil(t, err)
		assert.Equal(t, "failed to hash new password: failed to hash new password", err.Error())
	}


	func TestChangePasswordFailedToUpdatePassword(t *testing.T) {
		userRepository := new(databasemocks.MockUserRepository)
		passwordRepository := new(authautomocks.MockPasswordRepository)
		mockUseCase := NewChangePasswordUseCase(userRepository, passwordRepository)

		passwordRepository.On("VerifyPassword", "any_incorrect_password", "any_lengthy_password").Return(true).Once()
		passwordRepository.On("HashPassword", "new_any_lengthy_password").Return("new_any_lengthy_password", nil).Once()
		mockUser := &domain.User{
			ID:       "any_id",
			Password: "any_lengthy_password",
		}

		userRepository.On("GetById", "any_id").Return(mockUser, nil).Once()
		userRepository.On("UpdatePassword", "any_id", "new_any_lengthy_password").Return(fmt.Errorf("failed to update password")).Once()

		err := mockUseCase.Execute("any_id", "any_incorrect_password", "new_any_lengthy_password")

		assert.NotNil(t, err)
		assert.Equal(t, "failed to update password: failed to update password", err.Error())
		assert.True(t, errors.Is(err, err), "failed to update password")
	}

	/* 
	return nil
}
assert that the password is updated successfully
*/
func TestChangePasswordSuccess(t *testing.T) {
	userRepository := new(databasemocks.MockUserRepository)
	passwordRepository := new(authautomocks.MockPasswordRepository)
	mockUseCase := NewChangePasswordUseCase(userRepository, passwordRepository)

	passwordRepository.On("VerifyPassword", "any_incorrect_password", "any_lengthy_password").Return(true).Once()
	passwordRepository.On("HashPassword", "new_any_lengthy_password").Return("new_any_lengthy_password", nil).Once()
	mockUser := &domain.User{
		ID:       "any_id",
		Password: "any_lengthy_password",
	}

	userRepository.On("GetById", "any_id").Return(mockUser, nil).Once()
	userRepository.On("UpdatePassword", "any_id", "new_any_lengthy_password").Return(nil).Once()

	err := mockUseCase.Execute("any_id", "any_incorrect_password", "new_any_lengthy_password")
	assert.Nil(t, err)
}