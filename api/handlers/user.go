package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammetburakgolec/InvestHub-Backend/api/models"
	"github.com/muhammetburakgolec/InvestHub-Backend/helpers"
)

func Login(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": helpers.PARSE_ERROR})
	}
	var user models.User

	if err := user.FindByUsername(input.Username); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": helpers.USER_NOT_FOUND})
	}
	if !helpers.ChechPasswordHash(input.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": helpers.PASSWORD_ERROR})
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

func GetUser(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": helpers.PARSE_ERROR})
	}

	var user models.User

	if err := user.GetByUserID(input.Id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": helpers.USER_NOT_FOUND})
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": helpers.PARSE_ERROR})
	}

	hashedPassword, err := helpers.HashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": helpers.PASSWORD_HASH_ERROR})
	}

	input.Password = hashedPassword

	if err := input.Register(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": helpers.USER_CREATE_ERROR})
	}

	UserCustom := models.User{
		Id:         input.Id,
		Username:   input.Username,
		Password:   input.Password,
		GroupId:    input.GroupId,
		IsAdmin:    input.IsAdmin,
		IsInvestor: input.IsInvestor,
		IsStudent:  input.IsStudent,
	}

	return c.JSON(UserCustom)
}

func GetByAllByGroupId(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": helpers.PARSE_ERROR})
	}

	var user models.User

	if err := user.GetByAllGroupId(input.GroupId); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": helpers.USER_NOT_FOUND})
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
