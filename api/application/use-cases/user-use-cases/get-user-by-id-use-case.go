package userusecases

import (
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

func (userRepo *GetUserByIdUseCase) Execute(id string) (*userdtos.UserOutPutDTO, error) {

	existUser, err := userRepo.userRepository.GetById(id)
	if err != nil {
		return nil, err
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
