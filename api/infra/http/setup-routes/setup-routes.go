package setuproutes

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/middlewares"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/setup-routes/routes"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	// Create the auth group
	authGroup := router.Group("/api/v1").Use(middlewares.JWTMiddleware(), middlewares.SanitizeMiddleware())
	authGroupNotSanitized := router.Group("/api/v1").Use(middlewares.JWTMiddleware())
	/* 	onlyAdminRoutes := router.Group("/admin/v1").Use()
	 */ // Cast authGroup to *gin.RouterGroup, if we try to use authGroup directly, it will
	// show us an error because authGroup is type of gin.IRoutes and our routes are expecting
	// *gin.RouterGroup
	authGroupRoutes := authGroup.(*gin.RouterGroup)
	authGroupWithOutSanitization := authGroupNotSanitized.(*gin.RouterGroup)

	routes.AuthRoutes(router)                                                    // don't need authentication
	routes.SetupCategoryRoutes( /* router  */ authGroupRoutes)                   // openApi auth OK
	routes.SetupCarRoutes(router, authGroupRoutes, authGroupWithOutSanitization) // openApi auth OK
	routes.SetupSpecificationRoutes(router /* authGroupRoutes */)                // openApi auth OK
	routes.SetupUserRoutes(authGroupRoutes)                                      // openApi auth OK
	routes.SetupMaintenanceRoutes(authGroupRoutes)                               // openApi auth OK
	routes.SetupReviewRoutes(router, authGroupRoutes)                            // openApi auth OK
	routes.SetupOrderRoutes(authGroupRoutes, router)                             // openApi auth OK
}

// @Security 		BearerAuth
