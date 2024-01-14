package http

import (
	"log"
	"os"

	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/http/car-controller/endpoints"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Init() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.POST("/api/v1/create", endpoints.RegisterCarController)

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbconfig.Connection()

	port := os.Getenv("PORT")

	log.Println("Running server in port " + port)
	router.Run(":" + port)
}
