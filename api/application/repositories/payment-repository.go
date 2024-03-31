package repositories

import (
	paymentdtos "github.com/Marcosxx1/Car-Rent-gin-golang-/api/application/dtos"
	"github.com/Marcosxx1/Car-Rent-gin-golang-/api/domain"
)

type PaymentRepository interface {
	CreatePayment(payment *domain.Payment) error
	GetPaymentsByOptions(options *paymentdtos.PaymentOutput) ([]*domain.Payment, error)
	UpdatePayment(id string, payment *domain.Payment) error
	DeletePayment(paymentID string) error

	/* 	// Métodos adicionais
	   	// talvez seja mais adequado apenas uma rota GET, assim condensamos todas
	   	// estas por um filtro/opções
	   	ListPaymentsByOrderID(ctx context.Context, orderID string) ([]*domain.Payment, error)
	   	ListPaymentsByStatus(ctx context.Context, status string) ([]*domain.Payment, error)
	   	GetTotalPaymentsByOrderID(ctx context.Context, orderID string) (float64, error)
	   	GetLatestPaymentByOrderID(ctx context.Context, orderID string) (*domain.Payment, error)
	   	GetPaymentsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*domain.Payment, error)
	   	CountPaymentsByStatus(ctx context.Context, status string) (int, error)
	   	GetTotalPayments(ctx context.Context) (float64, error) */
}
