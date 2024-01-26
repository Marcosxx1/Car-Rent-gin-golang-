package database

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
)

/* package repositories

import "github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"

type SpecificationRepository interface {
	FindSpecificationByName(name string) (*domain.Specification, error)
	GetAll() ([]*domain.Specification, error)
	PostSpecification(specification *domain.Specification) error
}
*/

type PGSpecification struct{}

func (repo *PGSpecification) FindSpecificationByName(name string) (*domain.Specification, error) {
	var specification domain.Specification
	err := dbconfig.Postgres.Where("name = ?", name).First(&specification).Error
	if err != nil {
		return nil, err
	}
	return &specification, nil
}

func (repo *PGSpecification) GetAll() ([]*domain.Specification, error) {
	var specifications []*domain.Specification
	err := dbconfig.Postgres.Find(&specifications).Error

	if err != nil {
		return nil, err
	}
	return specifications, nil
}

func (repo *PGSpecification) PostSpecification(specification *domain.Specification) error {
	return dbconfig.Postgres.Create(specification).Error
}
