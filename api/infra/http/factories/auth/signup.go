package authfactory

import (
	authusecase "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/auth-use-case"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/auth-controller"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	userRepository := database.NewPGUserRepository()

	userUseCase := authusecase.NewPostUserUseCase(userRepository)

	auth.RegisterUserController(context, userUseCase)
}
