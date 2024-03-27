package database

import (
	"errors"
	"fmt"
	"strings"

	orderdto "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos/order"
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

// GetOrdersByOptions retrieves orders from the database based on provided options.
//
// Parameters:
//   - options: An OrderOutputDTO containing filtering options.
//
// Returns:
//   - []*domain.Order: A slice of domain.Order pointers representing the retrieved orders.
//   - error: An error if the operation fails, nil otherwise.
//
// The provided options are used to construct conditions for filtering the orders. If any of the
// filtering criteria are provided (ID, UserID, CarID, RentalStartDate, RentalEndDate), the function
// constructs corresponding conditions and applies them to the database query. If no filtering
// criteria are provided, all orders are retrieved.
// The function returns a slice of domain.Order pointers representing the retrieved orders
// along with an error. If no orders match the specified criteria, the function returns nil for
// the slice of orders and nil error.
func (r *PGOrderRepository) GetOrdersByOptions(options *orderdto.OrderOutputDTO) ([]*domain.Order, error) {
	var orders []*domain.Order
	query := dbconfig.Postgres

	conditions := []string{}

	/*  values := []interface{}{} cria um novo slice vazio, pronto
	    para receber valores. Se a declaração fosse simplesmente var
	    values []interface{}, values seria inicializado como nil, e
	   	não poderia ser usado diretamente para adicionar valores usando o operador .... */
			values := []interface{}{}

	// Construct conditions based on non-zero values in options
	if options.ID != "" {
		conditions = append(conditions, "id ILIKE ?")
		values = append(values, "%"+options.ID+"%")
	}
	if options.UserID != "" {
		conditions = append(conditions, "user_id ILIKE ?")
		values = append(values, "%"+options.UserID+"%")
	}
	if options.CarID != "" {
		conditions = append(conditions, "car_id ILIKE ?")
		values = append(values, "%"+options.CarID+"%")
	}
	if !options.RentalStartDate.IsZero() {
		conditions = append(conditions, "rental_start_date = ?")
		values = append(values, options.RentalStartDate)
	}
	if !options.RentalEndDate.IsZero() {
		conditions = append(conditions, "rental_end_date = ?")
		values = append(values, options.RentalEndDate)
	}

	if len(conditions) > 0 {
		query = query.Where(strings.Join(conditions, " OR "), values...) //https://yourbasic.org/golang/three-dots-ellipsis/
	}

	result := query.Find(&orders)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return orders, nil
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
