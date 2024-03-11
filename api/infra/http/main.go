// main.go in http/server
package main

import (
	"log"
	"os"

	_ "github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	_ "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/docs"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @title Serviço de locação de carros
// @version 1.0
// @description Serviço utilizando o framework gin

// @host localhost:8080
func setupServer(db *gorm.DB) error {
	router := gin.Default()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupCategoryRoutes(router)
	routes.SetupCarRoutes(router)
	routes.SetupSpecificationRoutes(router)
	routes.SetupUserRoutes(router)
	routes.SetupMaintenanceRoutes(router)
	routes.SetupReviewRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running at port:", port)
	return router.Run(":" + port)
}
