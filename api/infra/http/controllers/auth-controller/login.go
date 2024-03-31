package auth

import (
	"net/http"

	authdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/auth"
	authusecase "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/auth-use-case"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// LoginHandlerController handles the HTTP POST request for login.
// @Summary				Log in the application
// @Description			If the provided e-mail and password are correct an jwt token will be generated with the user id
// @ID					login
// @Tags				Auth
// @Accept				json
// @Produce				json
// @Param				request				body 			authdto.LoginInputDTO	true "user information to login"
// @Router				/login 			[post]
func LoginHandlerController(context *gin.Context, loginUseCase *authusecase.LoginUseCase) {
	var request *authdto.LoginInputDTO

	if err := context.ShouldBindJSON(&request); err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := loginUseCase.Execute(request)

	if err != nil {
		switch err {
		case authusecase.ErrInvalidPassword:
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		case authusecase.ErrGenerateAuthToken:
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate authentication token"})
		default:
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}
