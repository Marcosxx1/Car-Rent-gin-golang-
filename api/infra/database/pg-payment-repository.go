package database

import (
	"errors"
	"fmt"
	"strings"

	paymentdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/repositories"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
	dbconfig "github.com/Marcosxx1/Car-Rent-gin-golang-/api/infra/database/postgres/db-config"
	"gorm.io/gorm"
)

type PGPaymentRepository struct{}

func NewPGPaymentRepository() repositories.PaymentRepository {
	return &PGPaymentRepository{}
}

func (r *PGPaymentRepository) CreatePayment(payment *domain.Payment) error {
	return dbconfig.Postgres.Create(payment).Error
}

func (r *PGPaymentRepository) GetPaymentsByOptions(options *paymentdtos.PaymentOutput) ([]*domain.Payment, error) {
	var payments []*domain.Payment
	query := dbconfig.Postgres

	conditions := []string{}
	values := []interface{}{}

	/* 	if options.OrderID != "" {
		conditions = append(conditions, "order_id = ?")
		values = append(values, options.OrderID)
	} */
	if options.PaymentStatus != "" {
		conditions = append(conditions, "payment_status = ?")
		values = append(values, options.PaymentStatus)
	}
	if !options.PaymentDate.IsZero() {
		conditions = append(conditions, "payment_date = ?")
		values = append(values, options.PaymentDate)
	}

	if len(conditions) > 0 {
		query = query.Where(strings.Join(conditions, " AND "), values...)
	}

	result := query.Find(&payments)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return payments, nil
}

func (r *PGPaymentRepository) UpdatePayment(id string, payment *domain.Payment) error {
	result := dbconfig.Postgres.Model(&domain.Payment{}).Where("id = ?", id).Updates(payment)
	if result.Error != nil {
		return fmt.Errorf("failed to update payment record: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no payment record found with id: %s", id)
	}
	return nil
}

func (r *PGPaymentRepository) DeletePayment(paymentID string) error {
	result := dbconfig.Postgres.Where("id = ?", paymentID).Delete(&domain.Payment{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("payment not found")
	}

	return nil
}
