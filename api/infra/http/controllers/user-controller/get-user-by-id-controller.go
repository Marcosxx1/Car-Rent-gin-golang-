package usercontroller

import (
	"log"
	"net/http"

	userusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/user-use-cases"
	"github.com/gin-gonic/gin"
)

// @Summary Get user by ID
// @Description Retrieve user details by ID
// @Tags User
// @Accept json
// @Produce json
// @Security 		BearerAuth
// @Param id path string true "User ID"
// @Security ApiKeyAuth
// @Success 200 {object} userdtos.UserOutPutDTO
// @Router /user/{id} [get]
func GetUserByIdController(context *gin.Context, userUseCase *userusecases.GetUserByIdUseCase) {

	id := context.Param("id")

	user, err := userUseCase.Execute(id)
	if err != nil {
		log.Println("Error finding car", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, user)
}
