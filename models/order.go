package models

import "time"

type Order struct {
	OrderID       int       `json:"order_id"`
	UserID        int       `json:"user_id"`
	PaymentStatus string    `json:"payment_status"`
	TotalPayment  float64   `json:"total_payment"`
	CreatedAt     time.Time `json:"created_at"`
}

type OrderResponse struct {
	OrderID       int     `json:"order_id"`
	UserID        int     `json:"user_id"`
	PaymentStatus string  `json:"payment_status"`
	TotalPayment  float64 `json:"total_payment"`
}
