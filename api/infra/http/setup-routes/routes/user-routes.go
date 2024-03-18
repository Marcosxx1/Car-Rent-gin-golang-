package routes

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	userendpoints "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/controllers/user-controller/user-endpoints"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	userRepository := database.PGUserRepository{}

	router.GET("/api/v1/user/:id", func(contex *gin.Context) {
		userendpoints.GetUserByIdController(contex, &userRepository)
	})

	router.PATCH("/api/v1/user/:id", func(context *gin.Context) {
		userendpoints.PatchUserController(context, &userRepository)
	})

	router.PATCH("/api/v1/user/:id/change-password", func(context *gin.Context) {
		userendpoints.ChangePasswordController(context, &userRepository)
	})

}
