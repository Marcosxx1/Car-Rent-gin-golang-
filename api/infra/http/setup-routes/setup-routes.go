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
	authGroupRoutes := authGroup.(*gin.RouterGroup)

	routes.SetupCategoryRoutes(router /* authGroupRoutes */)
	routes.SetupCarRoutes(router, authGroupRoutes)
	routes.SetupSpecificationRoutes(router /* authGroupRoutes */)
	routes.SetupUserRoutes(router /* authGroupRoutes */)
	routes.SetupMaintenanceRoutes(authGroupRoutes)
	routes.SetupReviewRoutes(authGroupRoutes)
	routes.SetupOrderRoutes(authGroupRoutes)
}
