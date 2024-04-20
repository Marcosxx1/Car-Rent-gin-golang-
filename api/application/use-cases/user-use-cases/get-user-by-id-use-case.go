package userusecases

import (
	"errors"

	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/user"
	repositories "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
)

type GetUserByIdUseCase struct {
	userRepository repositories.UserRepository
}

func NewGetUserByIdUseCase(userRepository repositories.UserRepository) *GetUserByIdUseCase {
	return &GetUserByIdUseCase{
		userRepository: userRepository,
	}
}

func (useCase *GetUserByIdUseCase) Execute(id string) (*userdtos.UserOutPutDTO, error) {
	existUser, err := useCase.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	if existUser == nil {
		return nil, errors.New("user not found")
	}

	userToBeReturned := &userdtos.UserOutPutDTO{
		ID:        existUser.ID,
		Name:      existUser.Name,
		Email:     existUser.Email,
		Status:    existUser.Status,
		Avatar:    existUser.Avatar,
		CreatedAt: existUser.CreatedAt,
	}

	return userToBeReturned, nil
}
