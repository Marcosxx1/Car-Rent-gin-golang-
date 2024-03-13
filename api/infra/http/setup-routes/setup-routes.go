package setuproutes

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/setup-routes/routes"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	routes.SetupCategoryRoutes(router)
	routes.SetupCarRoutes(router)
	routes.SetupSpecificationRoutes(router)
	routes.SetupUserRoutes(router)
	routes.SetupMaintenanceRoutes(router)
	routes.SetupReviewRoutes(router)
}
