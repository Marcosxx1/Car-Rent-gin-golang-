package usercontroller

import (
	"net/http"

	authdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/auth"
	authusecase "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/auth-use-case"
	"github.com/gin-gonic/gin"
)

// @Summary Change user password
// @Description Change user password by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body authdto.ChangePasswordDTO true "Change Password Request"
// @Security ApiKeyAuth
// @Router /user/change-password [post]
func ChangePasswordController(context *gin.Context, userUseCase *authusecase.ChangePasswordUseCase) {
	userID := context.GetString("user_id")
	if userID == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}
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
