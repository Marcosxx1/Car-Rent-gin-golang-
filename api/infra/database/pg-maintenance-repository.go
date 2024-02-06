package database

import (
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
)

type PGMaintenanceRepository struct{}

func (r *PGMaintenanceRepository) CreateMaintenance(maintenance *domain.Maintenance) error {
	return dbconfig.Postgres.Create(maintenance).Error
}
