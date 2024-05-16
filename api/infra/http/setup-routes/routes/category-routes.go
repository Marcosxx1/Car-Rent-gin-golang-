package routes

import (
	categoryfactory "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/factories/category"
	"github.com/gin-gonic/gin"
)

func SetupCategoryRoutes( /* router *gin.Engine */ authGroup *gin.RouterGroup) {

	authGroup.POST("/category/create", categoryfactory.PostCategoryControllerFactory)
	authGroup.GET("/category/list", categoryfactory.ListCategoriesControllerFactory)

}
