package userusecases

import (
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/error_handling"
	hashpassword "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/user-controller/hash-password"
	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/user-controller/user-dtos"
	"github.com/rs/xid"
)

func PostUserUseCase(registerRequest userdtos.UserInputDTO,
	userRepository repositories.UserRepository) (*userdtos.UserOutPutDTO, error) {

	existUser, err := userRepository.FindByEmail(registerRequest.Email)

	if existUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", registerRequest.Email)
	}

	if err != nil {
		return nil, err
	}

	hashedPassword, err := hashpassword.HashPassword(registerRequest.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	newUser := &domain.User{
		ID:       xid.New().String(),
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: hashedPassword,
		Role:     registerRequest.Role,
		Status:   registerRequest.Status,
		Avatar:   registerRequest.Avatar,
	}

	if err := error_handling.ValidateStruct(newUser); err != nil {
		return nil, err
	}

	if err := userRepository.PostUser(newUser); err != nil {
		return nil, fmt.Errorf("failed to singin user: %w", err)
	}

	outPut := &userdtos.UserOutPutDTO{
		ID:     newUser.ID,
		Name:   newUser.Name,
		Email:  newUser.Email,
		Role:   newUser.Role,
		Status: newUser.Status,
		Avatar: newUser.Avatar,
	}

	return outPut, nil
}