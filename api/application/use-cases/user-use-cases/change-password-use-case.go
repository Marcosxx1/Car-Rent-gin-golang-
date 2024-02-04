// change-password-use-case
package userusecases

import (
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/error_handling"
	hashpassword "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/user-controller/hash-password"
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

	if err := error_handling.ValidateStruct(existingUser); err != nil {
		return err
	}

	// Update the user in the repository
	if err := userRepository.UpdatePassword(id, existingUserPassword); err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}
