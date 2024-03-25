package userendpoints

import (
	"net/http"

	authdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/auth"
	repositories "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	authusecase "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/auth-use-case"
	"github.com/gin-gonic/gin"
)

func ChangePasswordController(context *gin.Context, userRepo repositories.UserRepository) {
	userID := context.Param("id")

	var request authdto.ChangePasswordDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := authusecase.ChangePasswordUseCase(userID, request.CurrentPassword, request.NewPassword, userRepo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
