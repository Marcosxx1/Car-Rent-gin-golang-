package userusecases

import (
	r "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/user-controller/user-dtos"
)

func GetUserByIdUseCase(id string, repo r.UserRepository) (*userdtos.UserOutPutDTO, error) {

	existUser, err := repo.FindByEmail(id)
	if err != nil {
		return nil, err
	}

	userToBeReturned := &userdtos.UserOutPutDTO{
		ID:        existUser.ID,
		Name:      existUser.Name,
		Email:     existUser.Email,
		Role:      existUser.Role,
		Status:    existUser.Status,
		Avatar:    existUser.Avatar,
		CreatedAt: existUser.CreatedAt,
	}

	return userToBeReturned, nil
}
