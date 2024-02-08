package userusecases

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/user-controller/user-dtos"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
)

func PutUserUseCase(id string, putRequest userdtos.UserUpdateDTO,
	userRepository repositories.UserRepository) (*userdtos.UserOutPutDTO, error) {

	userToBeUpdated := userdtos.UserUpdateDTO{
		ID:     id,
		Name:   putRequest.Name,
		Email:  putRequest.Email,
		Status: putRequest.Status,
		Avatar: putRequest.Avatar,
	}

	if err := validation_errors.ValidateStruct(userToBeUpdated); err != nil {
		return nil, err
	}
	/* convert userToBeUpdated to domain.User */
	userToBeUpdatedDomain := &domain.User{
		ID:     id,
		Name:   userToBeUpdated.Name,
		Email:  userToBeUpdated.Email,
		Status: userToBeUpdated.Status,
		Avatar: userToBeUpdated.Avatar,
	}

	userUpdated, err := userRepository.Update(id, userToBeUpdatedDomain)

	if err != nil {
		return nil, err
	}

	userToBeReturned := &userdtos.UserOutPutDTO{
		ID:     userUpdated.ID,
		Name:   userToBeUpdated.Name,
		Email:  userToBeUpdated.Email,
		Status: userToBeUpdated.Status,
		Avatar: userToBeUpdated.Avatar,
	}

	return userToBeReturned, nil
}
