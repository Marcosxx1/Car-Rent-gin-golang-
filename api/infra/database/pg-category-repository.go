package database

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
)

type PGCategory struct {}

/* type CategoryPort interface {
	FindByName(name string) (*domain.Category, error)
    List(page, limit int) ([]*domain.Category, error)
    RegisterCategory(category *domain.Category) (*domain.Category, error)
} */

func (repo * PGCategory) RegisterCategory(category domain.Category) (*domain.Category, error) {
	dbconfig.Postgres.Create(&category)
	return &category, nil

}