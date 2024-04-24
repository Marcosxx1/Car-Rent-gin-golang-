package auth

import (
	"net/http"

	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/user"
	authusecase "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/auth-use-case"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// RegisterUserController handles the HTTP POST request to create a new user.
// @Summary				Create a new user
// @Description			Create a new user with the provided information
// @ID					sign-up
// @Tags				Auth
// @Accept				json
// @Produce				json
// @Param				request				body 			userdtos.UserInputDTO	true "user information to be created"
// @Router				/signup [post]
func RegisterUserController(context *gin.Context, userUseCase *authusecase.PostUserUseCase) {
	var request *userdtos.UserInputDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		validation_errors.NewValidationError(context, http.StatusUnprocessableEntity, err)
		return
	}

	createdUser, err := userUseCase.Execute(*request)
	if err != nil {
		validation_errors.NewValidationError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, createdUser)
}

