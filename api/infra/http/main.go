// main.go in http/server
package main

import (
	"log"
	"os"

	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	db, err := dbconfig.Connection()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	if err := setupServer(db); err != nil {
		log.Fatal("Error setting up server:", err)
	}
}

func setupServer(db *gorm.DB) error {
	router := gin.Default()

	routes.SetupCategoryRoutes(router)
	routes.SetupCarRoutes(router)
	routes.SetupSpecificationRoutes(router)
	routes.SetupUserRoutes(router)
	routes.SetupMaintenanceRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running at port:", port)
	return router.Run(":" + port)
}
