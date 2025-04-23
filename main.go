package main

import (
	"ebook-api/controller"
	"ebook-api/db"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := db.ConnectDB(); err != nil {
		log.Fatal("cannot connect database:", err)
	}

	app := fiber.New()

	app.Post("/users", controller.CreateUser)
	app.Get("/users", controller.GetUsers)

	log.Fatal(app.Listen(":3000"))
}
