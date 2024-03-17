package routes

import (
	categoryfactory "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/factories/category"
	"github.com/gin-gonic/gin"
)

func SetupCategoryRoutes(router *gin.Engine) {

	router.POST("/category/create", categoryfactory.PostCategoryControllerFactory)
	router.GET("/category/list", categoryfactory.ListCategoriesControllerFactory)

}
