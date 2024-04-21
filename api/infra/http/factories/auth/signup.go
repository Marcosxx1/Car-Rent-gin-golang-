package authfactory

import (
	authusecase "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/auth-use-case"
	authautho "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/auth-autho"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/auth-controller"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	userRepository := database.NewPGUserRepository()
	passwordRepository := authautho.NewPasswordRepository()
	userUseCase := authusecase.NewPostUserUseCase(userRepository, passwordRepository)

	auth.RegisterUserController(context, userUseCase)
}
