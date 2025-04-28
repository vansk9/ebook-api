package controller

import (
	"context"
	"ebook-api/dto"
	"ebook-api/models"
	"ebook-api/service"

	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	service service.OrderService
}

func NewOrderController(service service.OrderService) *OrderController {
	return &OrderController{service: service}
}

func (ctrl *OrderController) Create(c *fiber.Ctx) error {
	var req dto.CreateOrderRequest

	// Parse body sekali saja
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse order JSON",
		})
	}

	// Mapping ke models
	order := models.Order{
		UserID:        req.UserID,
		PaymentStatus: req.PaymentStatus,
	}

	var items []models.OrderItem
	for _, itemReq := range req.Items {
		items = append(items, models.OrderItem{
			ProductName: itemReq.ProductName,
			Quantity:    itemReq.Quantity,
			Price:       itemReq.Price,
			// TotalPrice dan CreatedAt akan dihitung di service
		})
	}

	createdOrder, err := ctrl.service.CreateOrder(context.Background(), order, items)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createdOrder)
}

func (ctrl *OrderController) GetAll(c *fiber.Ctx) error {
	orders, err := ctrl.service.GetAllOrders(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(orders)
}
