package factory

import (
	userusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/user-use-cases"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	usercontroller "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/user-controller"
	"github.com/gin-gonic/gin"
)

func PatchUserControllerFactory(context *gin.Context) {
	userRepository := database.NewPGUserRepository()
	userUseCase := userusecases.NewPutUserUseCase(userRepository)

	usercontroller.PatchUserController(context, userUseCase)

}
