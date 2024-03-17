package setuproutes

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/middlewares"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/setup-routes/routes"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	// Create the auth group
	authGroup := router.Group("/api/v1").Use(middlewares.JWTMiddleware(), middlewares.SanitizeMiddleware())

	// Cast authGroup to *gin.RouterGroup, if we try to use authGroup directly, it will
	// show us an error because authGroup is type of gin.IRoutes and our routes are expecting
	// *gin.RouterGroup
	authGroupPtr := authGroup.(*gin.RouterGroup)

	routes.SetupCategoryRoutes(router /* authGroupPtr */)
	routes.SetupCarRoutes(router /* authGroupPtr */)
	routes.SetupSpecificationRoutes(router /* authGroupPtr */)
	routes.SetupUserRoutes(router /* authGroupPtr */)
	routes.SetupMaintenanceRoutes(router, authGroupPtr)
	routes.SetupReviewRoutes(router, authGroupPtr)
	routes.SetupOrderRoutes(router, authGroupPtr)
}
