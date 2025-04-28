package repository

import (
	"context"
	"ebook-api/db"
	"ebook-api/models"
)

type OrderRepository interface {
	Create(ctx context.Context, order models.Order) (models.Order, error)
	GetAll(ctx context.Context) ([]models.Order, error)
}

type orderRepository struct{}

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}

func (r *orderRepository) Create(ctx context.Context, order models.Order) (models.Order, error) {
	query := `
	INSERT INTO "order" (user_id, payment_status, total_payment)
	VALUES ($1, $2, $3)
	RETURNING order_id, created_at
	`

	// Menyimpan order_id dan created_at
	err := db.Conn.QueryRow(ctx, query, order.UserID, order.PaymentStatus, order.TotalPayment).
		Scan(&order.OrderID, &order.CreatedAt)

	if err != nil {
		return models.Order{}, err
	}

	return order, nil

}

func (r *orderRepository) GetAll(ctx context.Context) ([]models.Order, error) {
	query := `SELECT order_id, user_id, payment_status, total_payment, created_at FROM "order"`

	rows, err := db.Conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.OrderID, &order.UserID, &order.PaymentStatus, &order.TotalPayment, &order.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
