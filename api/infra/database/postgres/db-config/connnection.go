package dbconfig

import (
	"fmt"
	"os"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Postgres *gorm.DB

func Connection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.AutoMigrate(&domain.Car{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.Specification{})
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.UserCar{})
	db.AutoMigrate(&domain.Maintenance{})
	db.AutoMigrate(&domain.Part{})
	db.AutoMigrate(&domain.CarMaintenance{})
	db.AutoMigrate(&domain.Reviews{})

	Postgres = db
	return db, nil
}
