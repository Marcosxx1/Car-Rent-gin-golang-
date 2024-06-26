package authfactory

import (
	authusecase "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/auth-use-case"
	authautho "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/auth-autho"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	auth "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/auth-controller"
	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	userRepository := database.NewPGUserRepository()
	passwordRepository := authautho.NewPasswordRepository()
	authRepository := authautho.NewAuthRepository()
	userUseCase := authusecase.NewLoginUseCase(userRepository, passwordRepository, authRepository)

	auth.LoginHandlerController(context, userUseCase)
}
