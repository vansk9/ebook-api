package dto

import "time"

type OrderItemResponse struct {
	ProductName string    `json:"product_name"`
	TotalPrice  float64   `json:"total_price"`
	CreatedAt   time.Time `json:"created_at"` // Menambahkan created_at di sini
}

type OrderResponse struct {
	OrderID       int                 `json:"order_id"`
	UserID        int                 `json:"user_id"`
	PaymentStatus string              `json:"payment_status"`
	TotalPayment  float64             `json:"total_payment"`
	CreatedAt     time.Time           `json:"created_at"`
	Items         []OrderItemResponse `json:"items"` // Response Items
}
