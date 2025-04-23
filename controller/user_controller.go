package controller

import (
	"context"
	"ebook-api/db"
	"ebook-api/models"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	query := `INSERT INTO "user" (name, email) VALUES ($1, $2) RETURNING user_id`
	err := db.Conn.QueryRow(context.Background(), query, user.Name, user.Email).Scan(&user.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	rows, err := db.Conn.Query(context.Background(), `SELECT user_id, name, email FROM "user"`)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserID, &user.Name, &user.Email); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		users = append(users, user)
	}

	return c.JSON(users)
}
