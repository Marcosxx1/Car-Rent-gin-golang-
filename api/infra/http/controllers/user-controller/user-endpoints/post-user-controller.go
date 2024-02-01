package userendpoints

import (
	"net/http"

	r "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	userusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/user-use-cases"
	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/user-controller/user-dtos"
	"github.com/gin-gonic/gin"
)

func RegisterUserController(context *gin.Context, userRepo r.UserRepository) {

	var request *userdtos.UserInputDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
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

	createdUser, err := userusecases.PostUserUseCase(*user, userRepo)
	println(err)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, createdUser)
	}
}
