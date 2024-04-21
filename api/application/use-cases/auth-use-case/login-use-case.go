package authusecase

import (
	"errors"

	authdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/auth"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
)

type LoginUseCase struct {
	userRepository     repositories.UserRepository
	passwordRepository repositories.PasswordRepository
	authRepository     repositories.AuthRepository
}

func NewLoginUseCase(userRepository repositories.UserRepository, passwordRepository repositories.PasswordRepository, authRepository repositories.AuthRepository) *LoginUseCase {
	return &LoginUseCase{
		userRepository:     userRepository,
		passwordRepository: passwordRepository,
		authRepository:     authRepository,
	}
}

func (useCase *LoginUseCase) Execute(request *authdto.LoginInputDTO) (string, error) {

	if err := validation_errors.ValidateStruct(request); err != nil {
		return "", err
	}

	existingUser, err := useCase.userRepository.FindByEmail(request.Email)
	if err != nil {
		return "", err
	}

	if existingUser.ID == "" {
		return "", errors.New("user not found")
	}

	if !useCase.passwordRepository.VerifyPassword(request.Password, existingUser.Password) {
		return "", errors.New("invalid username or password")
	}

	token, err := useCase.authRepository.GenerateAuthToken(existingUser.ID, string(existingUser.Role))
	if err != nil {
		return "", errors.New("failed to generate authentication token")
	}

	//useCase.context.SetCookie("token", token, int(expirationTime.Unix()), "/", "localhost", false, true)
	//useCase.context.JSON(200, gin.H{"success": "user logged in"})

	return token, nil
}
