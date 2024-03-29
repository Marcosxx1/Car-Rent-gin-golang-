package database

import (
	"errors"
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	"gorm.io/gorm"
)

type PGUserRepository struct{}

func NewPGUserRepository() repositories.UserRepository {
	return &PGUserRepository{}
}

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

func (repo *PGUserRepository) Update(id string, data *domain.User) (*domain.User, error) {
	dbconfig.Postgres.Model(&data).Where("id = ?", id).Updates(&data)
	return data, nil
}

func (repo *PGUserRepository) UpdatePassword(id, newPassword string) error {
	result := dbconfig.Postgres.Model(&domain.User{}).Where("id = ?", id).Update("password", newPassword)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user with ID %s not found", id)
	}

	return nil
}
