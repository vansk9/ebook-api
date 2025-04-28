package service

import (
	"context"
	"ebook-api/dto"
	"ebook-api/models"
	"ebook-api/repository"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order models.Order) (dto.OrderResponse, error)
	GetAllOrders(ctx context.Context) ([]dto.OrderResponse, error)
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) CreateOrder(ctx context.Context, order models.Order) (dto.OrderResponse, error) {
	newOrder, err := s.repo.Create(ctx, order)
	if err != nil {
		return dto.OrderResponse{}, err
	}

	// Menambahkan order ID dan informasi lain yang dibutuhkan
	return dto.OrderResponse{
		OrderID:       newOrder.OrderID,
		UserID:        newOrder.UserID,
		PaymentStatus: newOrder.PaymentStatus,
		TotalPayment:  newOrder.TotalPayment,
		CreatedAt:     newOrder.CreatedAt, // Menambahkan created_at
	}, nil
}

func (s *orderService) GetAllOrders(ctx context.Context) ([]dto.OrderResponse, error) {
	orders, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []dto.OrderResponse
	for _, order := range orders {
		responses = append(responses, dto.OrderResponse{
			OrderID:       order.OrderID,
			UserID:        order.UserID,
			PaymentStatus: order.PaymentStatus,
			TotalPayment:  order.TotalPayment,
		})
	}

	return responses, nil
}
