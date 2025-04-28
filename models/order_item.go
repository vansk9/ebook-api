package models

type OrderItem struct {
	ID          int
	OrderID     int
	ProductName string
	Quantity    int
	Price       float64
	TotalPrice  float64
}
