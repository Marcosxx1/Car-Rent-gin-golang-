package http

import (
	"log"
	"os"

	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	endpoints "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/car-controller/endpoints/car-endpoints"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Init() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/api/v1/cars", endpoints.ListCarController)
	router.GET("/api/v1/cars/:id", endpoints.FindCarByIdController)

	router.POST("/api/v1/cars/create", endpoints.RegisterCarController)

	router.DELETE("/api/v1/cars/delete/:id", endpoints.DeleteCarController)
	
	router.PUT("/api/v1/cars/update/:id", endpoints.UpdateCarController)

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbconfig.Connection()

	port := os.Getenv("PORT")

	log.Println("Running server in port " + port)
	router.Run(":" + port)
}
