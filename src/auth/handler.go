package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	path := ConfigGoogle()
	url := path.AuthCodeURL("state")

	return c.Redirect(url)
}

func Callback(c *fiber.Ctx) error {
	code := c.FormValue("code")

	token, error := ConfigGoogle().Exchange(c.Context(), code)
	if error != nil {

		return c.Status(401).JSON(fiber.Map{"error": error.Error()})
	}

	user := GetUser(token.AccessToken)
	return c.Status(200).JSON(fiber.Map{"data": user})
}
