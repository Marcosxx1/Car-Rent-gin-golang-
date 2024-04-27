package authusecase

import (
	"errors"
	"testing"

	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/user"
	authautomocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/auth-autho/auth-auto-mocks"
	databasemocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/database-mocks"
	"github.com/stretchr/testify/assert"
)

/* package authusecase

import (
	"errors"
	"testing"

	authdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/auth"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	authautomocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/auth-autho/auth-auto-mocks"
	databasemocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/database-mocks"
	"github.com/stretchr/testify/assert"
)


func TestLoginFailedValidation(t *testing.T) {
	mockRequest1 := &authdto.LoginInputDTO{
		Email:    "",
		Password: "",
	}
	runValidationTest(t, mockRequest1, "email is required; password is required")

	mockRequest2 := &authdto.LoginInputDTO{
		Email:    "invalidemail",
		Password: "short",
	}
	runValidationTest(t, mockRequest2, "email is invalid; password is required with min 8")

	mockRequest3 := &authdto.LoginInputDTO{
		Email:    "validemail@example.com",
		Password: "short",
	}
	runValidationTest(t, mockRequest3, "password is required with min 8")


}

func runValidationTest(t *testing.T, request *authdto.LoginInputDTO, expectedError string) {
	userRepository := new(databasemocks.MockUserRepository)
	authRepository := new(authautomocks.MockAuthRepository)
	passwordRepository := new(authautomocks.MockPasswordRepository)

	useCase := NewLoginUseCase(userRepository, passwordRepository, authRepository)

	token, err := useCase.Execute(request)

	if expectedError == "" {
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
	} else {
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err.Error())
		assert.Empty(t, token)
	}
}

func TestLoginUserNotFoundWithEmail(t *testing.T) {
	userRepository := new(databasemocks.MockUserRepository)
	authRepository := new(authautomocks.MockAuthRepository)
	passwordRepository := new(authautomocks.MockPasswordRepository)

	useCase := NewLoginUseCase(userRepository, passwordRepository, authRepository)

	userRepository.On("FindByEmail", "email@example.com").Return((*domain.User)(nil), errors.New("user not found"))
	token, err := useCase.Execute(&authdto.LoginInputDTO{
		Email:    "email@example.com",
		Password: "XXXXXXXX",
	})

	assert.NotNil(t, err)
	assert.Empty(t, token)
	assert.Equal(t, "user not found", err.Error())
	userRepository.AssertExpectations(t)
}

func TestLoginEmptyUserStructNotFound(t *testing.T) {
	userRepository := new(databasemocks.MockUserRepository)
	authRepository := new(authautomocks.MockAuthRepository)
	passwordRepository := new(authautomocks.MockPasswordRepository)

	useCase := NewLoginUseCase(userRepository, passwordRepository, authRepository)

	userRepository.On("FindByEmail", "email@example.com").Return(&domain.User{ID: ""}, nil)

	token, err := useCase.Execute(&authdto.LoginInputDTO{
		Email:    "email@example.com",
		Password: "XXXXXXXX",
	})

	assert.NotNil(t, err)
	assert.Empty(t, token)
	assert.Equal(t, "user not found", err.Error())

	userRepository.AssertExpectations(t)
}

func TestLoginInvalidPassword(t *testing.T) {
	userRepository := new(databasemocks.MockUserRepository)
	authRepository := new(authautomocks.MockAuthRepository)
	passwordRepository := new(authautomocks.MockPasswordRepository)

	useCase := NewLoginUseCase(userRepository, passwordRepository, authRepository)

	userRepository.On("FindByEmail", "email@example.com").Return(&domain.User{ID: "1", Email: "email@example.com", Password: "XXXXXXXX"}, nil)
	passwordRepository.On("VerifyPassword", "XXXXXXXX", "XXXXXXXX").Return(false)

	token, err := useCase.Execute(&authdto.LoginInputDTO{
		Email:    "email@example.com",
		Password: "XXXXXXXX",
	})

	assert.NotNil(t, err)
	assert.Empty(t, token)
	assert.Equal(t, "invalid username or password", err.Error())

	userRepository.AssertExpectations(t)
	passwordRepository.AssertExpectations(t)
}

func TestLoginGenerateAuthTokenFailed(t *testing.T) {
	userRepository := new(databasemocks.MockUserRepository)
	authRepository := new(authautomocks.MockAuthRepository)
	passwordRepository := new(authautomocks.MockPasswordRepository)

	useCase := NewLoginUseCase(userRepository, passwordRepository, authRepository)

	mockUser := &domain.User{
		ID:       "12345",
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "XXXXXXXX",
		Role:     "admin",
	}

	userRepository.On("FindByEmail", "email@example.com").Return(mockUser, nil)
	passwordRepository.On("VerifyPassword", "XXXXXXXX", "XXXXXXXX").Return(true)
	authRepository.On("GenerateAuthToken", "12345", "admin").Return("", errors.New("failed to generate authentication token"))

	token, err := useCase.Execute(&authdto.LoginInputDTO{
		Email:    "email@example.com",
		Password: "XXXXXXXX",
	})

	assert.NotNil(t, err)
	assert.Empty(t, token)
	assert.Equal(t, "failed to generate authentication token", err.Error())

	userRepository.AssertExpectations(t)
	passwordRepository.AssertExpectations(t)
	authRepository.AssertExpectations(t)
}

func TestLoginSuccess(t *testing.T) {
	userRepository := new(databasemocks.MockUserRepository)
	authRepository := new(authautomocks.MockAuthRepository)
	passwordRepository := new(authautomocks.MockPasswordRepository)

	useCase := NewLoginUseCase(userRepository, passwordRepository, authRepository)

	mockUser := &domain.User{
		ID:       "12345",
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "XXXXXXXX",
		Role:     "admin",
	}

	userRepository.On("FindByEmail", "email@example.com").Return(mockUser, nil)
	passwordRepository.On("VerifyPassword", "XXXXXXXX", "XXXXXXXX").Return(true)
	authRepository.On("GenerateAuthToken", "12345", "admin").Return("token", nil)

	token, err := useCase.Execute(&authdto.LoginInputDTO{
		Email:    "email@example.com",
		Password: "XXXXXXXX",
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
	userRepository.AssertExpectations(t)
	passwordRepository.AssertExpectations(t)
	authRepository.AssertExpectations(t)
	assert.Equal(t, "token", token)
	assert.Equal(t, "12345", mockUser.ID)
	assert.Equal(t, "admin", string(mockUser.Role))
	assert.Equal(t, "johndoe@example.com", mockUser.Email)
	assert.Equal(t, "John Doe", mockUser.Name)
	assert.Equal(t, "XXXXXXXX", mockUser.Password)

}
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

*/
func TestPostUserFailedValidation(t *testing.T){
	mockRequest1 := userdtos.UserInputDTO{
		Name:     "",
		Email:    "",
		Password: "",
		Role:     "",
		Status:   true,
		Avatar:   "",
	}
	runUserValidationTest(t, mockRequest1, "name is required; email is required; password is required; role is required")


	
}
func runUserValidationTest(t *testing.T, request userdtos.UserInputDTO, expectedError string) {
	userRepository := new(databasemocks.MockUserRepository)
	passwordRepository := new(authautomocks.MockPasswordRepository)

	useCase := NewPostUserUseCase(userRepository, passwordRepository)

	userOutputDto, err := useCase.Execute(request)

	if expectedError == "" {
		assert.Nil(t, err)
		assert.NotEmpty(t, userOutputDto)
	} else {
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err.Error())
		assert.Empty(t, userOutputDto)
	}
}
/* Test failed to hash password mockRepo.Mock:
	hashedPassword, err := useCase.passwordRepository.HashPassword(userInputDto.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	} */
	func TestFailedToHashPassword(t *testing.T){
		mockRequest2 := userdtos.UserInputDTO{
			Name:     "John Doe",
			Email:    "johndoe@example.com",
			Password: "XXXXXXXX",
			Role:     "admin",
			Status:   true,
			Avatar:   "",
		}
		userRepository := new(databasemocks.MockUserRepository)
		passwordRepository := new(authautomocks.MockPasswordRepository)

		userRepository.On("HashPassword", "XXXXXXXX").Return( "", errors.New("failed to hash password"))

		useCase := NewPostUserUseCase(userRepository, passwordRepository)

		userOutputDto, err := useCase.Execute(mockRequest2)

		assert.NotNil(t, err)
		assert.Empty(t, userOutputDto)
		assert.Equal(t, "failed to hash password", err.Error())
	}