package routes

import (
	category_endpoints "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/category-controller/category-endpoints"
	"github.com/gin-gonic/gin"
)

func SetupCategoryRoutes(router *gin.Engine){
	router.POST("/api/v1/category/create", category_endpoints.PostCategoryController)
}