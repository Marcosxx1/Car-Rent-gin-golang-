package authusecase

import (
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
)

type ChangePasswordUseCase struct {
	userRepository     repositories.UserRepository
	passwordRepository repositories.PasswordRepository
}

func NewChangePasswordUseCase(userRepository repositories.UserRepository, passwordRepository repositories.PasswordRepository) *ChangePasswordUseCase {
	return &ChangePasswordUseCase{
		userRepository:     userRepository,
		passwordRepository: passwordRepository,
	}
}

func (userRepo *ChangePasswordUseCase) Execute(id string, currentPassword string, newPassword string) error {
	existingUser, err := userRepo.userRepository.GetById(id)
	if err != nil {
		return fmt.Errorf("failed to retrieve user: %w", err)
	}
	fmt.Println(currentPassword, existingUser.Password)
	if !userRepo.passwordRepository.VerifyPassword(currentPassword, existingUser.Password) {
		return fmt.Errorf("current password does not match")
	}

	newHashedPassword, err := userRepo.passwordRepository.HashPassword(newPassword)
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
