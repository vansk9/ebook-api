package models

type Ebook struct {
	EbookID    int    `json:"ebook_id"`
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
	Price      int    `json:"price"`
}
