package userusecases

import (
	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/user"
	repositories "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
)

type PutUserUseCase struct {
	userRepository repositories.UserRepository
}

func NewPutUserUseCase(userRepository repositories.UserRepository) *PutUserUseCase {
	return &PutUserUseCase{
		userRepository: userRepository,
	}
}
 
func (userRepo *PutUserUseCase)Execute(id string, putRequest *userdtos.UserUpdateDTO) (*userdtos.UserOutPutDTO, error) {

/* 	userFound, err := userRepo.userRepository.GetById(id)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if userFound == nil {
		context.JSON(400, gin.H{"error": "User not found"})
		return
	} */


	if err := validation_errors.ValidateStruct(putRequest); err != nil {
		return nil, err
	}
	/* convert userToBeUpdated to domain.User */
	userToBeUpdatedDomain := &domain.User{
		ID:     id,
		Name:   putRequest.Name,
		Email:  putRequest.Email,
		Status: putRequest.Status,
		Avatar: putRequest.Avatar,
	}

	userUpdated, err := userRepo.userRepository.Update(id, userToBeUpdatedDomain)

	if err != nil {
		return nil, err
	}

	userToBeReturned := &userdtos.UserOutPutDTO{
		ID:     userUpdated.ID,
		Name:   userUpdated.Name,
		Email:  userUpdated.Email,
		Status: userUpdated.Status,
		Avatar: userUpdated.Avatar,
	}

	return userToBeReturned, nil
}
