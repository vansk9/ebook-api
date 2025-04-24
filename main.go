package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"ebook-api/controller"
	"ebook-api/db"
	"ebook-api/repository"
	"ebook-api/service"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Cek env
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set in the environment variables")
	}

	// Connect to DB
	if err := db.ConnectDB(); err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer db.Conn.Close()

	// Dependency Injection
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Fiber instance
	app := fiber.New()

	// Routes
	app.Post("/users", userController.Create)
	app.Get("/users", userController.GetAll)

	// Run
	log.Println("Server is running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
