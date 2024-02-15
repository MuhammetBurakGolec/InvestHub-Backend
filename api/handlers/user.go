// api/handlers/user.go

package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammetburakgolec/InvestHub-Backend/api/models"
	"github.com/muhammetburakgolec/InvestHub-Backend/helpers"
)

// Örnek olarak basit bir kullanıcı kontrolü
// Gerçek uygulamada bu bilgiler veritabanından alınmalı
var sampleUser = models.User{
	ID:       1,
	Username: "testuser",
	Password: "password", // Gerçek uygulamada şifre hash olarak karşılaştırılmalı
}

// Login Endpoint
func Login(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Kullanıcı adı ve şifre kontrolü
	if input.Username != sampleUser.Username || input.Password != sampleUser.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Incorrect username or password"})
	}

	// JWT Token Oluştur
	token, err := helpers.GenerateToken(input.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}
