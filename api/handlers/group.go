package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammetburakgolec/InvestHub-Backend/api/models"
	"github.com/muhammetburakgolec/InvestHub-Backend/helpers"
)

func GetGroup(c *fiber.Ctx) error {
	var input models.Group
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": helpers.PARSE_ERROR})
	}

	var group models.Group

	if err := group.GetByGroupID(input.GroupId); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": helpers.GROUP_NOT_FOUND})
	}

	return c.JSON(group)
}

func CreateGroup(c *fiber.Ctx) error {
	var input models.Group
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": helpers.PARSE_ERROR})
	}

	if err := input.CreateGroup(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": helpers.GROUP_CREATE_ERROR})
	}

	return c.JSON(input)
}
