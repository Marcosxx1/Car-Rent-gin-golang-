package userusecases

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/error_handling"
	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/user-controller/user-dtos"
)

func PutUserUseCase(id string, putRequest userdtos.UserInputDTO,
	userRepository repositories.UserRepository) (*userdtos.UserOutPutDTO, error) {

	userToBeUpdated := domain.User{
		ID:     id,
		Name:   putRequest.Name,
		Email:  putRequest.Email,
		Role:   putRequest.Role,
		Status: putRequest.Status,
		Avatar: putRequest.Avatar,
	}

	if err := error_handling.ValidateStruct(userToBeUpdated); err != nil {
		return nil, err
	}

	userUpdated, err := userRepository.Update(id, &userToBeUpdated)

	if err != nil {
		return nil, err
	}

	userToBeReturned := &userdtos.UserOutPutDTO{
		ID:     userUpdated.ID,
		Name:   userToBeUpdated.Name,
		Email:  userToBeUpdated.Email,
		Role:   userToBeUpdated.Role,
		Status: userToBeUpdated.Status,
		Avatar: userToBeUpdated.Avatar,
	}

	return userToBeReturned, nil
}
