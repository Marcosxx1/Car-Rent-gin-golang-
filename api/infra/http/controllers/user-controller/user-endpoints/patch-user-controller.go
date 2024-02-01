package userendpoints

import (
	"net/http"

	r "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	userusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/user-use-cases"
	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/user-controller/user-dtos"
	"github.com/gin-gonic/gin"
)

func PatchUserController(context *gin.Context, userRepository r.UserRepository) {

	var request userdtos.UserInputDTO
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

	user := &userdtos.UserInputDTO{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Role:     request.Role,
		Status:   request.Status,
		Avatar:   request.Avatar,
	}

	updatedUser, err := userusecases.PutUserUseCase(id, *user, userRepository)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, updatedUser)
	}
}
