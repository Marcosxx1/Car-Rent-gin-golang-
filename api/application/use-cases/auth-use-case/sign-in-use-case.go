package authusecase

import (
	"fmt"

	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/user"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/rs/xid"
)

type PostUserUseCase struct {
	userRepository     repositories.UserRepository
	passwordRepository repositories.PasswordRepository
}

func NewPostUserUseCase(userRepository repositories.UserRepository, passwordRepository repositories.PasswordRepository) *PostUserUseCase {
	return &PostUserUseCase{
		userRepository:     userRepository,
		passwordRepository: passwordRepository,
	}
}

func (useCase *PostUserUseCase) Execute(userInputDto userdtos.UserInputDTO) (*userdtos.UserOutPutDTO, error) {

	if err := validation_errors.ValidateStruct(userInputDto); err != nil {
		return nil, err
	}

	existingUser, err := useCase.userRepository.FindByEmail(userInputDto.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}

	if existingUser != nil {
		return nil, fmt.Errorf("user already exists")
	}

	hashedPassword, err := useCase.passwordRepository.HashPassword(userInputDto.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	newUser := &domain.User{
		ID:       xid.New().String(),
		Name:     userInputDto.Name,
		Email:    userInputDto.Email,
		Password: hashedPassword,
		Role:     userInputDto.Role,
		Status:   userInputDto.Status,
		Avatar:   userInputDto.Avatar,
	}

	if err := useCase.userRepository.PostUser(newUser); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return &userdtos.UserOutPutDTO{
		ID:     newUser.ID,
		Name:   newUser.Name,
		Email:  newUser.Email,
		Status: newUser.Status,
		Avatar: newUser.Avatar,
	}, nil
}
