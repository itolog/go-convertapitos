package user

import "github.com/gofiber/fiber/v2"

func (handler *Handler) findAll(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "users",
	})
}

func (handler *Handler) create(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "users",
	})
}
