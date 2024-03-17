package authusecase

import (
	"fmt"

	hashpassword "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/auth-controller/hash-password"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
)

func ChangePasswordUseCase(id string, currentPassword string, newPassword string, userRepository repositories.UserRepository) error {
	existingUser, err := userRepository.GetById(id)
	if err != nil {
		return fmt.Errorf("failed to retrieve user: %w", err)
	}

	if !hashpassword.VerifyPassword(currentPassword, existingUser.Password) {
		return fmt.Errorf("current password does not match")
	}

	newHashedPassword, err := hashpassword.HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("failed to hash new password: %w", err)
	}

	existingUserPassword := newHashedPassword

	if err := validation_errors.ValidateStruct(existingUser); err != nil {
		return err
	}

	if err := userRepository.UpdatePassword(id, existingUserPassword); err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}
