package userusecases

import (
	"errors"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	hashpassword "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/user-controller/hash-password"
	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/user-controller/user-dtos"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/middlewares/auth"
)

var (
	ErrInvalidPassword   = errors.New("invalid username or password")
	ErrGenerateAuthToken = errors.New("failed to generate authentication token")
)

type LoginUseCase struct {
	userRepository repositories.UserRepository
}

func NewLoginUseCase(userRepository repositories.UserRepository) *LoginUseCase {
	return &LoginUseCase{
		userRepository: userRepository,
	}
}

func (useCase *LoginUseCase) Execute(request *userdtos.LoginInputDTO) (string, error) {
	existingUser, err := useCase.userRepository.FindByEmail(request.Email)
	if err != nil {
		return "", err
	}

	if existingUser == nil {
		return "", ErrInvalidPassword
	}

	if !hashpassword.VerifyPassword(request.Password, existingUser.Password) {
		return "", ErrInvalidPassword
	}

	token, err := auth.GenerateAuthToken(existingUser.ID)
	if err != nil {
		return "", ErrGenerateAuthToken
	}

	return token, nil
}
