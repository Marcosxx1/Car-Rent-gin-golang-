package userendpoints

import (
	"net/http"

	r "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	userusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/user-use-cases"
	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/user-controller/user-dtos"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

func LoginHandlerController(context *gin.Context, UserRepository r.UserRepository) {
	var request *userdtos.LoginInputDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	loginUseCase := userusecases.NewLoginUseCase(UserRepository)
	token, err := loginUseCase.Execute(request)

	if err != nil {
		switch err {
		case userusecases.ErrInvalidPassword:
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		case userusecases.ErrGenerateAuthToken:
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate authentication token"})
		default:
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}
