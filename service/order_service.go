package service

import (
	"context"
	"ebook-api/dto"
	"ebook-api/models"
	"ebook-api/repository"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order models.Order, items []models.OrderItem) (dto.OrderResponse, error)
	GetAllOrders(ctx context.Context) ([]dto.OrderResponse, error)
	SaveOrderItem(ctx context.Context, orderID int, item models.OrderItem) error
	GetItemsByOrderID(ctx context.Context, orderID int) ([]models.OrderItem, error)
}

type orderService struct {
	repo repository.OrderRepository
	repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo} // Mengembalikan pointer ke orderService
}

func (s *orderService) CreateOrder(ctx context.Context, order models.Order, items []models.OrderItem) (dto.OrderResponse, error) {
	var totalPayment float64
	var orderItems []models.OrderItem

	// Proses item dan hitung total pembayaran
	for _, item := range items {
		// Hitung total_price untuk setiap item
		item.TotalPrice = float64(item.Quantity) * item.Price
		totalPayment += item.TotalPrice
		orderItems = append(orderItems, item)
	}

	// Update total_payment pada order
	order.TotalPayment = totalPayment

	// Simpan order ke database
	newOrder, err := s.repo.Create(ctx, order)
	if err != nil {
		return dto.OrderResponse{}, err
	}

	// Simpan setiap item ke database
	for _, item := range orderItems {
		err := s.repo.SaveOrderItem(ctx, newOrder.OrderID, item)
		if err != nil {
			return dto.OrderResponse{}, err
		}
	}

	// Format response items hanya dengan product_name, total_price, dan created_at
	var orderItemResponses []dto.OrderItemResponse
	for _, item := range orderItems {
		orderItemResponses = append(orderItemResponses, dto.OrderItemResponse{
			ProductName: item.ProductName,
			TotalPrice:  item.TotalPrice,
			CreatedAt:   item.CreatedAt, // Menambahkan created_at untuk setiap item
		})
	}

	// Return response
	return dto.OrderResponse{
		OrderID:       newOrder.OrderID,
		UserID:        newOrder.UserID,
		PaymentStatus: newOrder.PaymentStatus,
		TotalPayment:  newOrder.TotalPayment,
		CreatedAt:     newOrder.CreatedAt,
		Items:         orderItemResponses,
	}, nil
}

func (s *orderService) GetAllOrders(ctx context.Context) ([]dto.OrderResponse, error) {
	orders, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []dto.OrderResponse
	for _, order := range orders {
		var orderItems []dto.OrderItemResponse
		// Ambil item untuk setiap order (misalnya dari repo)
		items, err := s.repo.GetItemsByOrderID(ctx, order.OrderID) // Memanggil method GetItemsByOrderID
		if err != nil {
			return nil, err
		}

		// Map ke order response
		for _, item := range items {
			orderItems = append(orderItems, dto.OrderItemResponse{
				ProductName: item.ProductName,
			})
		}

		// Map order ke response
		responses = append(responses, dto.OrderResponse{
			OrderID:       order.OrderID,
			UserID:        order.UserID,
			PaymentStatus: order.PaymentStatus,
			TotalPayment:  order.TotalPayment,
			CreatedAt:     order.CreatedAt,
			Items:         orderItems,
		})
	}

	return responses, nil
}
