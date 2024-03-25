package userendpoints

import (
	"net/http"

	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/user"
	repositories "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	userusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/user-use-cases"
	"github.com/gin-gonic/gin"
)

func PatchUserController(context *gin.Context, userRepository repositories.UserRepository) {

	var request userdtos.UserUpdateDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := context.Param("id")

	userFound, err := userRepository.GetById(id)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if userFound == nil {
		context.JSON(400, gin.H{"error": "User not found"})
		return
	}

	user := &userdtos.UserUpdateDTO{
		Name:   request.Name,
		Email:  request.Email,
		Status: request.Status,
		Avatar: request.Avatar,
	}

	updatedUser, err := userusecases.PutUserUseCase(id, *user, userRepository)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, updatedUser)
	}
}
