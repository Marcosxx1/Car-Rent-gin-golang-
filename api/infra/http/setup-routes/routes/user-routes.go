package routes

import (
	factory "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/factories/user"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(authGroup *gin.RouterGroup) {

	authGroup.GET("/user/:id", factory.GetUserByIdControllerFactory)
	authGroup.PATCH("/user/:id", factory.PatchUserControllerFactory)
	authGroup.PATCH("/user/change-password", factory.ChangePasswordControllerFactory)

}
