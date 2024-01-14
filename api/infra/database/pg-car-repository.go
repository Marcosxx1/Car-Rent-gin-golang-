package database

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
)

type PGCarRepository struct{}

func (repo *PGCarRepository) RegisterCar(car domain.Car) *domain.Car {
	dbconfig.Postgres.Create(&car)
	return &car
}
 