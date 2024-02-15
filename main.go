package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammetburakgolec/InvestHub-Backend/api/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Post("/api/login", handlers.Login)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}
