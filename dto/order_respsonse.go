package dto

import "time"

type OrderItemResponse struct {
	ID          int       `json:"id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	TotalPrice  float64   `json:"total_price"`
	CreatedAt   time.Time `json:"created_at"`
}

type OrderResponse struct {
	OrderID       int                 `json:"order_id"`
	UserID        int                 `json:"user_id"`
	PaymentStatus string              `json:"payment_status"`
	TotalPayment  float64             `json:"total_payment"`
	CreatedAt     time.Time           `json:"created_at"`
	Items         []OrderItemResponse `json:"items"`
}
