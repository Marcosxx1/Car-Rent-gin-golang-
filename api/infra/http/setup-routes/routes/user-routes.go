package routes

import (
	factory "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/factories/user"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	
	router.GET("/user/:id", factory.GetUserByIdControllerFactory)
	router.PATCH("/user/:id", factory.PatchUserControllerFactory)
	router.PATCH("/user/:id/change-password", factory.ChangePasswordControllerFactory)

}
