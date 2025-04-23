package models

type Order struct {
	OrderID       int    `json:"order_id"`
	UserID        int    `json:"user_id"`
	PaymentStatus string `json:"payment_status"`
	TotalPayment  int    `json:"total_payment"`
	CreatedAt     string `json:"created_at"`
}
