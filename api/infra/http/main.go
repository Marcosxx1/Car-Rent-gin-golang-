// main.go in http/server
package main

import (
	"log"
	"os"

	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
 
	router := gin.Default()
	routes.SetupCategoryRoutes(router)
	routes.SetupCarRoutes(router)

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbconfig.Connection()

	port := os.Getenv("PORT")

	log.Println("localhost:8080/api/v1/cars" + port)
	router.Run(":" + port)
}
