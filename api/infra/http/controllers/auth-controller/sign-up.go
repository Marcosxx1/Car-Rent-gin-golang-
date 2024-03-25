package auth

import (
	"net/http"

	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/user"
	authusecase "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/auth-use-case"
	"github.com/gin-gonic/gin"
)

func RegisterUserController(context *gin.Context, userUseCase *authusecase.PostUserUseCase) {

	var request *userdtos.UserInputDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := userUseCase.Execute(*request)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, createdUser)
	}
}
