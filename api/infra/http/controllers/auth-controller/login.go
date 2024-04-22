package auth

import (
	"net/http"

	authdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/auth"
	authusecase "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/auth-use-case"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/validation_errors"
	"github.com/gin-gonic/gin"
)

// RegisterUserController handles the HTTP POST request to create a new maintenance.
// @Summary				Create a new user
// @Description			Create a new maintenance with the provided information
// @ID					login
// @Tags				Auth
// @Accept				json
// @Produce				json
// @Param				request				body 			authdto.LoginInputDTO	true "information to log in"
// @Router				/login [post]
func LoginHandlerController(context *gin.Context, loginUseCase *authusecase.LoginUseCase) {
	var request authdto.LoginInputDTO

	if err := context.ShouldBindJSON(&request); err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}
 
	token, err := loginUseCase.Execute(&request)
 	if err != nil {
		validation_errors.NewError(context, http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, token)
}
