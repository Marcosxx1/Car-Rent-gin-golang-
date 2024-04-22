package authusecase

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

	/*
		 	mockRequest4 := &authdto.LoginInputDTO{ // this way we aren't getting error, but the test will not pass TODO
				Email:    "validemail@example.com",
				Password: "validpassword",
			}
			runValidationTest(t, mockRequest4, "")
	*/
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