package routes

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database"
	category_endpoints "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/category-controller/category-endpoints"
	"github.com/gin-gonic/gin"
)

func SetupCategoryRoutes(router *gin.Engine) {
	
	categoryRepository := database.PGCategory{} 

	router.POST("/api/v1/category/create", func(context *gin.Context) {
		category_endpoints.PostCategoryController(context, &categoryRepository)
	})
}