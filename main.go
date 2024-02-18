package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammetburakgolec/InvestHub-Backend/api/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.GetHome)
	app.Post("/api/login", handlers.Login)
	// app.Post("/api/register", handlers.Register)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":5001")
}
