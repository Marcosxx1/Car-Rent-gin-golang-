package usercontroller

import (
	"net/http"

	userdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/user"
	userusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/user-use-cases"
	"github.com/gin-gonic/gin"
)

// @Summary Update user by ID
// @Description Update user details by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body userdtos.UserUpdateDTO true "Update User Request"
// @Security Bearer
// @Param Authorization header string true "Authorization"
// @Success 200 {object} userdtos.UserUpdateDTO
// @Router /user/{id} [patch]
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
