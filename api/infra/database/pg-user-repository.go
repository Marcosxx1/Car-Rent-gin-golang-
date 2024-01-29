package database

import (
	"errors"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	"gorm.io/gorm"
)

type PGUserRepository struct{}

func (repo *PGUserRepository) PostUser(userData *domain.User) error {
	return dbconfig.Postgres.Create(&userData).Error
}

func (repo *PGUserRepository) FindByEmail(email string) (*domain.User, error) {
	var user *domain.User
	err := dbconfig.Postgres.Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

/* 	GetById(id string) (*domain.User, error)
 */
func (repo *PGUserRepository) GetById(id string) (*domain.User, error) {
	var user *domain.User
	err := dbconfig.Postgres.Where("id = ?", id).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
