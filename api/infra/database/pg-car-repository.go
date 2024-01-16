package database

import (
	"errors"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	"gorm.io/gorm"
)

type PGCarRepository struct{}

func (repo *PGCarRepository) RegisterCar(car domain.Car) *domain.Car {
	dbconfig.Postgres.Create(&car)
	return &car
}
													  
func (repo *PGCarRepository) FindCarByLicensePlate(licensePlate string) (*domain.Car, error) {
	var car domain.Car
	err := dbconfig.Postgres.Where("license_plate = ?", licensePlate).First(&car).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil 
		}
		return nil, err 
	}

	return &car, nil
}

func (repo *PGCarRepository) FindAllCars() ([]*domain.Car, error) {
	var cars []*domain.Car
	err := dbconfig.Postgres.Find(&cars).Error

	if err != nil {
		return nil, err
	}

	return cars, nil
}
