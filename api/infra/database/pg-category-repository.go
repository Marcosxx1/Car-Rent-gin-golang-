package database

import (
	"errors"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	"gorm.io/gorm"
)

type PGCategory struct{}

func (repo *PGCategory) PostCategory(category *domain.Category) error {
	return dbconfig.Postgres.Create(category).Error
}

func (repo *PGCategory) FindCategoryByName(name string) (*domain.Category, error) {
	var category domain.Category
	err := dbconfig.Postgres.Where("name = ?", name).First(&category).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &category, nil
}
