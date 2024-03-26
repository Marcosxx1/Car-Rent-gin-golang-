package usercontroller

import (
	"net/http"

	authdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/auth"
	authusecase "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/auth-use-case"
	"github.com/gin-gonic/gin"
)

func ChangePasswordController(context *gin.Context, userUseCase *authusecase.ChangePasswordUseCase) {
	userID := context.Param("id")

	var request authdto.ChangePasswordDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := userUseCase.Execute(userID, request.CurrentPassword, request.NewPassword)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
