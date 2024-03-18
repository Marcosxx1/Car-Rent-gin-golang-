package docssetup

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
 
func Setup(router *gin.Engine) {

   docs.SwaggerInfo.BasePath = ""
   router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
