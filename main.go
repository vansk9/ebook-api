package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"ebook-api/controller"
	"ebook-api/db"
	"ebook-api/repository"
	"ebook-api/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := db.ConnectDB(); err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer db.Conn.Close()

	app := fiber.New()

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Inisialisasi repository -> service -> controller
	orderRepo := repository.NewOrderRepository(db.Conn)
	orderService := service.NewOrderService(orderRepo)
	orderController := controller.NewOrderController(orderService)

	app.Post("/users", userController.Create)
	app.Get("/users", userController.GetAll)

	app.Get("/orders", orderController.GetAll)
	app.Post("/orders", orderController.Create)

	log.Println("Server running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
