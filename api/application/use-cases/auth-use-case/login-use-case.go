package authusecase

import (
	"errors"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/auth-controller/auth"
	hashpassword "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/auth-controller/hash-password"
	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/user-controller/user-dtos"
	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidPassword   = errors.New("invalid username or password")
	ErrGenerateAuthToken = errors.New("failed to generate authentication token")
)

type LoginUseCase struct {
	userRepository repositories.UserRepository
	context        *gin.Context
}

func NewLoginUseCase(context *gin.Context, userRepository repositories.UserRepository) *LoginUseCase {
	return &LoginUseCase{
		userRepository: userRepository,
		context:        context,
	}
}

func (useCase *LoginUseCase) Execute(request *userdtos.LoginInputDTO) (string, error) {
	existingUser, err := useCase.userRepository.FindByEmail(request.Email)
	if err != nil {
		return "", err
	}

	if existingUser.ID == "" {
		useCase.context.JSON(400, gin.H{"error": "user does not exist"})
		return "", nil // TODO better error handling
	}

	if !hashpassword.VerifyPassword(request.Password, existingUser.Password) {
		return "", ErrInvalidPassword
	}

	token, err := auth.GenerateAuthToken(existingUser.ID)
	if err != nil {
		return "", ErrGenerateAuthToken
	}

	//useCase.context.SetCookie("token", token, int(expirationTime.Unix()), "/", "localhost", false, true)
	//useCase.context.JSON(200, gin.H{"success": "user logged in"})

	return token, nil
}
