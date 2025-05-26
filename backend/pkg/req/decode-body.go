package req

import "github.com/gofiber/fiber/v2"

func DecodeBody[T any](c *fiber.Ctx) (*T, error) {
	payload := new(T)

	err := c.BodyParser(payload)
	if err != nil {
		return payload, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return payload, nil
}
