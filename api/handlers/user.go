package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammetburakgolec/InvestHub-Backend/api/models"
	"github.com/muhammetburakgolec/InvestHub-Backend/helpers"
)

func Login(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	var user models.User

	if err := user.FindByUsername(input.Username); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	if !helpers.ChechPasswordHash(input.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Incorrect password"})
	}

	token, err := helpers.GenerateToken(user.Id, user.Username, user.GroupId, user.IsAdmin, user.IsInvestor, user.IsStudent)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}

func GetHome(c *fiber.Ctx) error {
	return c.SendString("<h1>Api Service</h1>")
}

func GetProfile(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var user models.User

	if err := user.GetByID(input.Id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	userWithoutPassword := models.User{
		Id:         user.Id,
		Username:   user.Username,
		IsAdmin:    user.IsAdmin,
		IsInvestor: user.IsInvestor,
		IsStudent:  user.IsStudent,
	}

	return c.JSON(userWithoutPassword)
}

func Register(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	hashedPassword, err := helpers.HashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not hash password"})
	}

	input.Password = hashedPassword

	if err := input.Register(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
	}

	UserwithoutID := models.User{
		Username:   input.Username,
		Password:   input.Password,
		GroupId:    input.GroupId,
		IsAdmin:    input.IsAdmin,
		IsInvestor: input.IsInvestor,
		IsStudent:  input.IsStudent,
	}

	return c.JSON(UserwithoutID)
}
