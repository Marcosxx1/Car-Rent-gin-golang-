package database

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
)

type PGCategory struct {}

func (repo * PGCategory) PostCategory(category domain.Category) (*domain.Category, error) {
	dbconfig.Postgres.Create(&category)
	return &category, nil

}