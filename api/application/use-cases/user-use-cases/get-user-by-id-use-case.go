package userusecases

import (
	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/user"
	repositories "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
)

func GetUserByIdUseCase(id string, repo repositories.UserRepository) (*userdtos.UserOutPutDTO, error) {

	existUser, err := repo.GetById(id)
	if err != nil {
		return nil, err
	}
	println(existUser)
	//fmt.Printf("%+v\n", existUser)

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
