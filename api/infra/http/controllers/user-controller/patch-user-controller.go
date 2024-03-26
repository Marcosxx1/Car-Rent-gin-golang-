package usercontroller

import (
	"net/http"

	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/user"
	userusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/user-use-cases"
	"github.com/gin-gonic/gin"
)

func PatchUserController(context *gin.Context, userUseCase *userusecases.PutUserUseCase) {

	var request userdtos.UserUpdateDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := context.Param("id")

	updatedUser, err := userUseCase.Execute(id, &request)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		context.JSON(http.StatusOK, updatedUser)
	}
}
