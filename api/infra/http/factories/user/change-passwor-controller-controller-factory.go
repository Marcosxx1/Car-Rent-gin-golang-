package factory

import (
	authusecase "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/auth-use-case"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	usercontroller "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/user-controller"
	"github.com/gin-gonic/gin"
)

func ChangePasswordControllerFactory(context *gin.Context) {
	userRepository := database.NewPGUserRepository()
	userUseCase := authusecase.NewChangePasswordUseCase(userRepository)

	usercontroller.ChangePasswordController(context, userUseCase)

}
