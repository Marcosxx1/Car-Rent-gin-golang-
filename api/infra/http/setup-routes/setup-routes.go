package setuproutes

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/middlewares"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/setup-routes/routes"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	// Create the auth group
	authGroup := router.Group("/api/v1").Use(middlewares.JWTMiddleware(), middlewares.SanitizeMiddleware())
	/* 	onlyAdminRoutes := router.Group("/admin/v1").Use()
	 */ // Cast authGroup to *gin.RouterGroup, if we try to use authGroup directly, it will
	// show us an error because authGroup is type of gin.IRoutes and our routes are expecting
	// *gin.RouterGroup
	authGroupRoutes := authGroup.(*gin.RouterGroup)
	routes.AuthRoutes(router)                                     // don't need authentication
	routes.SetupCategoryRoutes(router /* authGroupRoutes */)      // need to be admin
	routes.SetupCarRoutes(router, authGroupRoutes)                //don't need authentication
	routes.SetupSpecificationRoutes(router /* authGroupRoutes */) //don't need authentication
	routes.SetupUserRoutes(router /* authGroupRoutes */)          // can be admin or normal user
	routes.SetupMaintenanceRoutes(authGroupRoutes)                // needs to be admin
	routes.SetupReviewRoutes(router, authGroupRoutes)                     // need to be logged user
	routes.SetupOrderRoutes(authGroupRoutes, router)              // needs to be admin
}
