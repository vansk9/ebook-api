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

	orderRepo := repository.NewOrderRepository()
	orderService := service.NewOrderService(orderRepo)
	orderController := controller.NewOrderController(orderService)

	app.Post("/users", userController.Create) // Tambah route POST untuk membuat user
	app.Get("/users", userController.GetAll)  // Tambah route GET untuk mendapatkan semua user

	app.Get("/orders", orderController.GetAll)  // Tambah route GET untuk mendapatkan semua order
	app.Post("/orders", orderController.Create) // Tambah route POST untuk membuat order

	log.Println("Server running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
