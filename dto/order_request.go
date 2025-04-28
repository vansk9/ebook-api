package dto

type CreateOrderItemRequest struct {
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

type CreateOrderRequest struct {
	UserID        int                      `json:"user_id"`
	PaymentStatus string                   `json:"payment_status"`
	Items         []CreateOrderItemRequest `json:"items"`
}
