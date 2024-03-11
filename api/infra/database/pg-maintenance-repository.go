package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain/enums"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	"gorm.io/gorm"
)

type PGMaintenanceRepository struct{}

// /**
//   - Creates a new maintenance record in the database.
//     *
//   - @param maintenance The maintenance record to create.
//   - @return error If any error occurs while creating the maintenance record.
//     */
func (r *PGMaintenanceRepository) CreateMaintenance(maintenance *domain.Maintenance) error {
	return dbconfig.Postgres.Create(maintenance).Error
}

// /**
//   - Retrieves a single maintenance record by its ID.
//     *
//   - @param id The ID of the maintenance record to retrieve.
//   - @return (*domain.Maintenance, error) A pointer to the retrieved maintenance record and any error encountered.
//     */
func (r *PGMaintenanceRepository) GetMaintenanceByID(id string) (*domain.Maintenance, error) {
	var maintenance domain.Maintenance

	rows, err := dbconfig.Postgres.Raw("SELECT * FROM maintenances WHERE id = ?", id).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}
		return nil, gorm.ErrRecordNotFound
	}

	if err := dbconfig.Postgres.ScanRows(rows, &maintenance); err != nil {
		return nil, err
	}

	return &maintenance, nil
}

// /**
//   - Updates an existing maintenance record based on ID and provided information.
//     *
//   - @param maintenance The maintenance record with updated information.
//   - @param id The ID of the maintenance record to update.
//   - @return error If any error occurs while updating the maintenance record.
//     */
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

// /**
//   - Deletes a maintenance record based on ID.
//     *
//   - @param id The ID of the maintenance record to delete.
//   - @return error If any error occurs while deleting the maintenance record.
//     */
func (r *PGMaintenanceRepository) DeleteMaintenance(id string) error {
	result := dbconfig.Postgres.Where("id = ?", id).Delete(&domain.Maintenance{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("maintenance not found")
	}

	return nil
}

// /**
//   - Retrieves all maintenance records paginated by provided page and page size.
//   - Eagerly loads associated parts for each record.
//     *
//   - @param page The page number (starting from 1).
//   - @param pageSize The number of records per page.
//   - @return ([]*domain.Maintenance, error) A slice of pointers to retrieved maintenance records and any error encountered.
//     */
func (r *PGMaintenanceRepository) ListAllMaintenances(page, pageSize int) ([]*domain.Maintenance, error) {
	var maintenances []*domain.Maintenance

	offset := (page - 1) * pageSize

	err := dbconfig.Postgres.Offset(offset).Limit(pageSize).Preload("Parts").Find(&maintenances).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*domain.Maintenance{}, nil
		}
		return nil, err
	}

	return maintenances, nil
}

//   - Retrieves all maintenance records associated with a specific car ID.
//   - Eagerly loads associated parts for each record.
//     *
//
// Parameters:
//   - carID The ID of the car to retrieve associated maintenances for.
//
// Returns:
//   - ([]*domain.Maintenance, error) A slice of pointers to retrieved maintenance records and any error encountered.
func (r *PGMaintenanceRepository) GetMaintenancesByCarID(carID string) ([]*domain.Maintenance, error) {
	var maintenances []*domain.Maintenance

	err := dbconfig.Postgres.Where("car_id = ?", carID).Preload("Parts").Find(&maintenances).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*domain.Maintenance{}, nil
		}
		return nil, err
	}

	return maintenances, nil
}

// GetScheduledMaintenances retrieves an slice of maintenances
//
// Example:
//
//	maintenances := GetScheduledMaintenances()
//
// Returns
//   - []*domain.Maintenance
//   - error
func (r *PGMaintenanceRepository) GetScheduledMaintenances() ([]*domain.Maintenance, error) {
	var scheduledMaintenances []*domain.Maintenance

	err := dbconfig.Postgres.Where("scheduled_maintenance = ?", true).Preload("Parts").Find(&scheduledMaintenances).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*domain.Maintenance{}, nil
		}
		return nil, err
	}

	return scheduledMaintenances, nil
}

// - Retrieves all maintenance records with a specific maintenance status.
// - Eagerly loads associated parts for each record.
//
// - @param status The maintenance status to filter by.
//
// - @return ([]*domain.Maintenance, error) A slice of pointers to retrieved maintenance records and any error encountered.
func (r *PGMaintenanceRepository) GetMaintenancesByStatus(status enums.MaintenanceStatus) ([]*domain.Maintenance, error) {
	var maintenances []*domain.Maintenance

	// Enable Gorm logging
	db := dbconfig.Postgres.Debug()

	err := db.Where("maintenance_status = ?", status).Preload("Parts").Find(&maintenances).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*domain.Maintenance{}, nil
		}
		return nil, err
	}

	return maintenances, nil
}

func (r *PGMaintenanceRepository) GetLatestMaintenanceByCar(carID string) (*domain.Maintenance, error) {
	var latestMaintenance domain.Maintenance

	err := dbconfig.Postgres.
		Where("car_id = ?", carID).
		Order("maintenance_completion_date desc").
		Limit(1).
		Preload("Parts").
		First(&latestMaintenance).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &latestMaintenance, nil
}

// GetMaintenancesByDateRange retrieves all maintenance records within a specified date range.
// Eagerly loads associated parts for each record.
//
// Parameters:
//   - startDate The start date of the range in the format "2006-01-02".
//   - endDate The end date of the range in the format "2006-01-02".
//
// Returns:
//   - ([]*domain.Maintenance, error) A slice of pointers to retrieved maintenance records and any error encountered.
func (r *PGMaintenanceRepository) GetMaintenancesByDateRange(startDate, endDate string) ([]*domain.Maintenance, error) {
	var maintenances []*domain.Maintenance

	startDateTime, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start date: %w", err)
	}

	endDateTime, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end date: %w", err)
	}

	err = dbconfig.Postgres.
		Where("maintenance_completion_date BETWEEN ? AND ?", startDateTime, endDateTime).
		Preload("Parts").
		Find(&maintenances).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*domain.Maintenance{}, nil
		}
		return nil, err
	}

	return maintenances, nil
}
