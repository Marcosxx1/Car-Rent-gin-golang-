package userendpoints

import (
	"net/http"

	r "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	userusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/user-use-cases"
	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/user-controller/user-dtos"
	"github.com/gin-gonic/gin"
)

func ChangePasswordController(context *gin.Context, userRepo r.UserRepository) {
	// Get user ID from the URL parameter
	userID := context.Param("id")

	// Bind JSON request to DTO
	var request userdtos.ChangePasswordDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the change password use case
	err := userusecases.ChangePasswordUseCase(userID, request.CurrentPassword, request.NewPassword, userRepo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
