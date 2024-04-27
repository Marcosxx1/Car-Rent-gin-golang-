package authusecase

import (
	"errors"
	"testing"

	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/user"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	authautomocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/auth-autho/auth-auto-mocks"
	databasemocks "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/database-mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostUserFailedValidation(t *testing.T) {
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

func TestFailedToHashPassword(t *testing.T) {
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

	userRepository.On("FindByEmail", mockRequest2.Email).Return((*domain.User)(nil), errors.New("failed to check existing user"))

	useCase := NewPostUserUseCase(userRepository, passwordRepository)

	userOutputDto, err := useCase.Execute(mockRequest2)

	assert.NotNil(t, err)
	assert.Empty(t, userOutputDto)
	assert.Equal(t, "failed to check existing user", err.Error())
}

//NOTE
/* (*domain.User)(nil) it is for when we want to return nil
   (&domain.User{}) normally a return of the database when nothing is found and
   we're returning an empty struct */
func TestFailedToFetchUser(t *testing.T) {
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

	userRepository.On("FindByEmail", mockRequest2.Email).Return((&domain.User{}), nil)

	useCase := NewPostUserUseCase(userRepository, passwordRepository)

	userOutputDto, err := useCase.Execute(mockRequest2)

	assert.NotNil(t, err)
	assert.Empty(t, userOutputDto)
	assert.Equal(t, "user already exists", err.Error())
}

func TestPostUserFailedToHashPassword(t *testing.T) {
	mockRequest := userdtos.UserInputDTO{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "password123",
		Role:     "admin",
		Status:   true,
		Avatar:   "",
	}

	mockPasswordRepository := new(authautomocks.MockPasswordRepository)
	mockPasswordRepository.On("HashPassword", mockRequest.Password).Return("", errors.New("failed to hash password"))

	mockUserRepository := new(databasemocks.MockUserRepository)
	mockUserRepository.On("FindByEmail", mockRequest.Email).Return((*domain.User)(nil), nil)

	useCase := NewPostUserUseCase(mockUserRepository, mockPasswordRepository)

	userOutputDto, err := useCase.Execute(mockRequest)

	assert.NotNil(t, err)
	assert.Nil(t, userOutputDto)
	assert.Equal(t, "failed to hash password: failed to hash password", err.Error())
}

func TestPostUserFailedToSaveUser(t *testing.T) {
	mockRequest := userdtos.UserInputDTO{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "password123",
		Role:     "admin",
		Status:   true,
		Avatar:   "",
	}

	mockPasswordRepository := new(authautomocks.MockPasswordRepository)
	mockPasswordRepository.On("HashPassword", mockRequest.Password).Return("hashedPassword", nil)

	mockUserRepository := new(databasemocks.MockUserRepository)
	mockUserRepository.On("FindByEmail", mockRequest.Email).Return((*domain.User)(nil), nil)
	mockUserRepository.On("PostUser", mock.AnythingOfType("*domain.User")).Return(errors.New("failed to save user"))

	useCase := NewPostUserUseCase(mockUserRepository, mockPasswordRepository)

	userOutputDto, err := useCase.Execute(mockRequest)

	assert.NotNil(t, err)
	assert.Nil(t, userOutputDto)
	assert.Equal(t, "failed to save user: failed to save user", err.Error()) // Adjust the error message if necessary
}

func TestPostUserSuccess(t *testing.T) {
	mockRequest := userdtos.UserInputDTO{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "password123",
		Role:     "admin",
		Status:   true,
		Avatar:   "",
	}

	mockPasswordRepository := new(authautomocks.MockPasswordRepository)
	mockPasswordRepository.On("HashPassword", mockRequest.Password).Return("hashedPassword", nil)

	mockUserRepository := new(databasemocks.MockUserRepository)
	mockUserRepository.On("FindByEmail", mockRequest.Email).Return((*domain.User)(nil), nil)

	mockUserRepository.On("PostUser", mock.AnythingOfType("*domain.User")).Return(nil)

	useCase := NewPostUserUseCase(mockUserRepository, mockPasswordRepository)

	userOutputDto, err := useCase.Execute(mockRequest)

	assert.Nil(t, err)
	assert.NotNil(t, userOutputDto)

}
