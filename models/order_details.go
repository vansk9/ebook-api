package models

type OrderDetails struct {
	OrderDetailsID int `json:"order_details_id"`
	EbookID        int `json:"ebook_id"`
	OrderID        int `json:"order_id"`
	Price          int `json:"price"`
	Quantity       int `json:"quantity"`
}
