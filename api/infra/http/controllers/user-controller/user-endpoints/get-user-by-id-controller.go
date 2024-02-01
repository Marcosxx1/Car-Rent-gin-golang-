package userendpoints

import (
	"log"
	"net/http"

	r "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	userusecases "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/use-cases/user-use-cases"
	"github.com/gin-gonic/gin"
)

func GetUserByIdController(context *gin.Context, userRepo r.UserRepository) {

	id := context.Param("id")

	user, err := userusecases.GetUserByIdUseCase(id, userRepo)
	if err != nil {
		log.Println("Error finding car", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, user)
}
