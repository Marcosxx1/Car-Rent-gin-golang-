package usercontroller

import (
	"log"
	"net/http"

	userusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/user-use-cases"
	"github.com/gin-gonic/gin"
)

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
