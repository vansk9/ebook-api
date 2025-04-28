package repository

import (
	"context"
	"ebook-api/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository interface {
	Create(ctx context.Context, order models.Order) (models.Order, error)
	GetAll(ctx context.Context) ([]models.Order, error)
	SaveOrderItem(ctx context.Context, orderID int, item models.OrderItem) error
	GetItemsByOrderID(ctx context.Context, orderID int) ([]models.OrderItem, error)
}

type orderRepository struct {
	DB *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) OrderRepository {
	return &orderRepository{DB: db}
}

func (r *orderRepository) Create(ctx context.Context, order models.Order) (models.Order, error) {
	query := `
        INSERT INTO "order" (user_id, payment_status, total_payment)
        VALUES ($1, $2, $3)
        RETURNING order_id, created_at
    `
	row := r.DB.QueryRow(ctx, query, order.UserID, order.PaymentStatus, order.TotalPayment)

	var createdOrder models.Order
	err := row.Scan(&createdOrder.OrderID, &createdOrder.CreatedAt)
	if err != nil {
		return models.Order{}, err
	}

	createdOrder.UserID = order.UserID
	createdOrder.PaymentStatus = order.PaymentStatus
	createdOrder.TotalPayment = order.TotalPayment

	return createdOrder, nil
}

func (r *orderRepository) GetAll(ctx context.Context) ([]models.Order, error) {
	query := `
        SELECT order_id, user_id, payment_status, total_payment, created_at
        FROM "order"
    `
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(
			&order.OrderID,
			&order.UserID,
			&order.PaymentStatus,
			&order.TotalPayment,
			&order.CreatedAt,
		); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *orderRepository) SaveOrderItem(ctx context.Context, orderID int, item models.OrderItem) error {
	query := `
        INSERT INTO order_items (order_id, product_name, quantity, price, total_price)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := r.DB.Exec(ctx, query,
		orderID,
		item.ProductName,
		item.Quantity,
		item.Price,
		float64(item.Quantity)*item.Price, // Total Price dihitung di sini
	)
	return err
}

func (r *orderRepository) GetItemsByOrderID(ctx context.Context, orderID int) ([]models.OrderItem, error) {
	query := `
        SELECT id, product_name, quantity, price, total_price, created_at
        FROM order_items
        WHERE order_id = $1
    `
	rows, err := r.DB.Query(ctx, query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.OrderItem
	for rows.Next() {
		var item models.OrderItem
		if err := rows.Scan(
			&item.ID,
			&item.ProductName,
			&item.Quantity,
			&item.Price,
			&item.TotalPrice,
			&item.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
