package routes

import (
	authfactory "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/factories/auth"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	router.POST("/login", authfactory.Login)
	router.POST("/signup", authfactory.Signup)
	//router.GET("/logout", authfactory.Logout)
}
