package userendpoints

import (
	"log"
	"net/http"

	r "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/gin-gonic/gin"
)

func GetUserByIdController(context *gin.Context, userRepo r.UserRepository) {

	id := context.Param("id")

	user, err := userRepo.GetById(id)
	if err != nil {
		log.Println("Error finding car", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, user)
}
