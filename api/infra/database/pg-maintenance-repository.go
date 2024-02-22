package database

import (
	"errors"
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	"gorm.io/gorm"
)

type PGMaintenanceRepository struct{}

func (r *PGMaintenanceRepository) CreateMaintenance(maintenance *domain.Maintenance) error {
	return dbconfig.Postgres.Create(maintenance).Error
}

func (r *PGMaintenanceRepository) GetMaintenanceByID(id string) (*domain.Maintenance, error){
	var maintenance domain.Maintenance

	err := dbconfig.Postgres.Where("id = ?", id, true).First(&maintenance).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, nil
	}

	return &maintenance, nil
}

func (r *PGMaintenanceRepository) UpdateMaintenance(maintenance *domain.Maintenance, id string) error {
	result := dbconfig.Postgres.Model(&domain.Maintenance{}).Where("id = ?", id).Updates(maintenance)

	if result.Error != nil {
			return fmt.Errorf("failed to update maintenance record: %w", result.Error)
	}

	if result.RowsAffected == 0 {
			return fmt.Errorf("no maintenance record found with id: %s", id)
	}

	return nil
}

