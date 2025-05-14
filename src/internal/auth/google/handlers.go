package google

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/pkg/api"
	"golang.org/x/oauth2"
	"io"
)

const userUrl = "https://www.googleapis.com/oauth2/v1/userinfo"

type authGoogleHandler struct{}

func newAuthGoogleHandler() *authGoogleHandler {
	return &authGoogleHandler{}
}

func (google authGoogleHandler) auth(c *fiber.Ctx) error {
	path := ConfigGoogle()
	url := path.AuthCodeURL("state")

	return c.Redirect(url)
}

func (google authGoogleHandler) callback(c *fiber.Ctx) error {
	code := c.FormValue("code")
	session, err := SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	token, err := ConfigGoogle().Exchange(c.Context(), code)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(api.DataResponse{
			Error: api.ErrorResponse{
				Code:    fiber.StatusUnauthorized,
				Message: "Failed to exchange code for token",
				Details: err.Error(),
			},
			Status: api.StatusError,
		})
	}

	user, err := google.getUser(token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.DataResponse{
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
	err = session.Save()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.DataResponse{
			Error: api.ErrorResponse{
				Code:    fiber.StatusInternalServerError,
				Message: "Something went wrong",
				Details: err.Error(),
			},
			Status: api.StatusError,
		})
	}

	return c.Redirect("/auth/google/profile")
}

func (google authGoogleHandler) profile(c *fiber.Ctx) error {
	session, err := SessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	user := session.Get("user")

	return c.Status(fiber.StatusOK).JSON(api.DataResponse{
		Data: fiber.Map{
			"user": user,
		},
		Status: api.StatusSuccess,
	})
}

func (google authGoogleHandler) getUser(token *oauth2.Token) (ResponseGoogle, error) {
	client := ConfigGoogle().Client(context.Background(), token)

	response, err := client.Get(userUrl)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(response.Body)

	if err != nil {
		return ResponseGoogle{}, fmt.Errorf("%v", err)
	}

	var user ResponseGoogle

	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		return ResponseGoogle{}, fmt.Errorf("%v", err)
	}
	return user, nil
}
