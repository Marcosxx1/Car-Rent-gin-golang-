package authusecase

import (
	"fmt"

	hashpassword "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/auth-controller/hash-password"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
)

type ChangePasswordUseCase struct{
	userRepository repositories.UserRepository
}

func NewChangePasswordUseCase(usreRepository repositories.UserRepository)*ChangePasswordUseCase{
	return &ChangePasswordUseCase{
		userRepository: usreRepository,
	}
}

func(userRepo *ChangePasswordUseCase) Execute(id string, currentPassword string, newPassword string) error {
	existingUser, err := userRepo.userRepository.GetById(id)
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

	if err := userRepo.userRepository.UpdatePassword(id, existingUserPassword); err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}
