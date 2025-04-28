package controller

import (
	"ebook-api/models"
	"ebook-api/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) Create(c *fiber.Ctx) error {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	// Membuat user model dari request
	user := models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	// Memanggil service untuk membuat user
	createdUser, err := uc.service.CreateUser(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Mapping data dari model User ke UserResponse
	userResponse := models.UserResponse{
		Success: true,
		Message: "User Berhasil dibuat",
		Data: models.UserData{
			ID:    createdUser.ID,
			Name:  createdUser.Name,
			Email: createdUser.Email,
		},
	}

	// Mengembalikan response yang sudah dimapping
	return c.Status(fiber.StatusCreated).JSON(userResponse)
}

func (uc *UserController) GetAll(c *fiber.Ctx) error {
	// Mendapatkan data user dari service
	users, err := uc.service.GetAllUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Mapping data model User ke UserResponse
	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, models.UserResponse{
			Success: true,
			Message: "User ditemukan",
			Data: models.UserData{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			},
		})
	}

	// Mengembalikan response yang sudah dimapping
	return c.JSON(userResponses)
}
