package auth

import (
	"net/http"

	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/user"
	authusecase "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/auth-use-case"
	"github.com/gin-gonic/gin"
)

// RegisterUserController handles the HTTP POST request to create a new maintenance.
// @Summary				Create a new user
// @Description		Create a new maintenance with the provided information
// @ID						post-user
// @Tags					User
// @Accept				json
// @Produce				json
// @Param					request				body 			userdtos.UserInputDTO	true "user information to be created"
// @Router				/signup [post]
func RegisterUserController(context *gin.Context, userUseCase *authusecase.PostUserUseCase) {
	var request *userdtos.UserInputDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := userUseCase.Execute(*request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, createdUser)
}
