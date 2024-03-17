package authfactory

import (
	authusecase "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/auth-use-case"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	auth "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/auth-controller"
	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	userRepository := database.NewPGUserRepository()

	userUseCase := authusecase.NewLoginUseCase(context, userRepository)

	auth.LoginHandlerController(context, userUseCase)
}
