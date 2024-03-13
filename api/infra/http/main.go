// main.go in http/server
package main

import (
	"log"
	"os"

	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	docssetup "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/docs/setup-docs-routes"
	setuproutes "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/setup-routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	_, err := dbconfig.Connection()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	router := gin.Default()
	docssetup.Setup(router)
	setuproutes.Setup(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running at port:", port)
	router.Run(":" + port)
}
