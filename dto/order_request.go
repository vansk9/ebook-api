// dto/order_request.go
package dto

type CreateOrderRequest struct {
	UserID        int    `json:"user_id"`
	PaymentStatus string `json:"payment_status"`
	Items         []struct {
		ProductName string  `json:"product_name"`
		Quantity    int     `json:"quantity"`
		Price       float64 `json:"price"`
	} `json:"items"`
}
