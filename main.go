package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammetburakgolec/InvestHub-Backend/api/handlers"
)

func setupRoutes(app *fiber.App) {

	// Home Routes
	app.Get("/", handlers.GetHome)

	// Authorization Routes
	app.Post("/api/login", handlers.Login)
	app.Post("/api/register", handlers.Register)

	// User Routes
	app.Get("/api/user", handlers.GetUser)

	// Group Routes
	app.Get("/api/group", handlers.GetGroup)
	app.Post("/api/group", handlers.CreateGroup)
	app.Get("/api/group/:id", handlers.GetByAllByGroupId)

	// User Routes
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":5001")
}
