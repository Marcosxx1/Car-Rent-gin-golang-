package authusecase

import (
	"errors"

	authdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/auth"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/auth-controller/auth"
	hashpassword "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/auth-controller/hash-password"
	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidPassword   = errors.New("invalid username or password")
	ErrGenerateAuthToken = errors.New("failed to generate authentication token")
)
/* type UserRepository interface {
	PostUser(userData *domain.User) error
	GetById(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Update(id string, data *domain.User) (*domain.User, error)
	UpdatePassword(id, newPassword string) error
}
 */
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

func (useCase *LoginUseCase) Execute(request *authdto.LoginInputDTO) (string, error) {
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

	token, err := auth.GenerateAuthToken(existingUser.ID, string(existingUser.Role))
	if err != nil {
		return "", ErrGenerateAuthToken
	}

	//useCase.context.SetCookie("token", token, int(expirationTime.Unix()), "/", "localhost", false, true)
	//useCase.context.JSON(200, gin.H{"success": "user logged in"})

	return token, nil
}
