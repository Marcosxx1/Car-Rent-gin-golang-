package database

import (
	"errors"
	"fmt"

	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	"gorm.io/gorm"
)

type PGOrderRepository struct{}

func NewPGOrderRepository() repositories.OrderRepository {
	return &PGOrderRepository{}
}

// CreateOrder creates a new order in the database.
//
// Parameters:
//   - order: The order to be created.
//
// Returns:
//   - error: An error if the operation fails, nil otherwise.
func (r *PGOrderRepository) CreateOrder(order *domain.Order) error {
	return dbconfig.Postgres.Create(&order).Error
}

// GetOrderByID retrieves an order from the database by its ID.
//
// Parameters:
//   - orderID: The ID of the order to retrieve.
//
// Returns:
//   - *domain.Order: A pointer to the retrieved order, nil if not found.
//   - error: An error if the operation fails, nil otherwise.
func (r *PGOrderRepository) GetOrderByID(orderID string) (*domain.Order, error) {
	var order domain.Order
	result := dbconfig.Postgres.First(&order, "id = ?", orderID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &order, nil
}

// UpdateOrder updates an existing order in the database.
//
// Parameters:
//   - id: The ID of the order to update.
//   - order: A pointer to the updated order.
//
// Returns:
//   - error: An error if the operation fails, nil otherwise.
func (r *PGOrderRepository) UpdateOrder(id string, order *domain.Order) error {
	result := dbconfig.Postgres.Model(&domain.Order{}).Where("id = ?", id).Updates(order)
	if result.Error != nil {
		return fmt.Errorf("failed to update order record: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no order record found with id: %s", id)
	}
	return nil
}

func (r *PGOrderRepository) DeleteOrder(orderID string) error {
	result := dbconfig.Postgres.Where("id = ?", orderID).Delete(&domain.Order{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("order not found")
	}

	return nil
}
