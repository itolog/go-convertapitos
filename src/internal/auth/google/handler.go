package google

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/pkg/api"
)

func GoogleAuth(c *fiber.Ctx) error {
	path := ConfigGoogle()
	url := path.AuthCodeURL("state")

	return c.Redirect(url)
}

func GoogleAuthCallback(c *fiber.Ctx) error {
	code := c.FormValue("code")
	session, err := SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	token, error := ConfigGoogle().Exchange(c.Context(), code)
	if error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(api.ApiResponse{
			Error: api.ErrorResponse{
				Code:    fiber.StatusUnauthorized,
				Message: "Failed to exchange code for token",
				Details: error.Error(),
			},
			Status: api.StatusError,
		})
	}

	user, err := GetUser(token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.ApiResponse{
			Error: api.ErrorResponse{
				Code:    fiber.StatusInternalServerError,
				Message: "Something went wrong",
				Details: err.Error(),
			},
			Status: api.StatusError,
		})
	}

	jsonBytes, err := json.Marshal(&user)
	if err != nil {
		panic(err)
	}

	session.Set("user", string(jsonBytes))
	session.Save()

	return c.Redirect("/auth/google/profile")
}

func GoogleProfile(c *fiber.Ctx) error {
	session, err := SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	user := session.Get("user")
	fmt.Println(user)

	return c.Status(fiber.StatusOK).JSON(api.ApiResponse{
		Data: fiber.Map{
			"user": user,
		},
		Status: api.StatusSuccess,
	})
}
