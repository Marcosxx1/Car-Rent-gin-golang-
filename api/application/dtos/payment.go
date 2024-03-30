package paymentdtos

import "time"

type PaymentInput struct {
	OrderID       uint      `json:"order_id"`
	PaymentDate   time.Time `json:"payment_date"`
	PaymentAmount float64   `json:"payment_amount"`
	PaymentStatus string    `json:"payment_status"`
}

// PaymentOutput representa a estrutura de saída para retornar informações sobre um pagamento.
type PaymentOutput struct {
	ID            string    `json:"id"`
	OrderID       uint      `json:"order_id"`
	PaymentDate   time.Time `json:"payment_date"`
	PaymentAmount float64   `json:"payment_amount"`
	PaymentStatus string    `json:"payment_status"`
}
