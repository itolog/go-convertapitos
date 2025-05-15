package google

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/pkg/api"
	"golang.org/x/oauth2"
	"io"
	"time"
)

const userUrl = "https://www.googleapis.com/oauth2/v1/userinfo"

func (handler *Handler) login(c *fiber.Ctx) error {
	//from := c.Query("from")
	path := configs.ConfigGoogle()
	url := path.AuthCodeURL("state")
	//fmt.Println("from", from)
	//fmt.Println(url)
	return c.Redirect(url)
}

func (handler *Handler) callback(c *fiber.Ctx) error {
	code := c.FormValue("code")

	token, err := configs.ConfigGoogle().Exchange(c.Context(), code)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(api.Response[any]{
			Error: &api.ErrorResponse{
				Message: "Unauthorized",
				Details: err.Error(),
				Code:    fiber.StatusUnauthorized,
			},
			Status: api.StatusError,
		})
	}

	user, err := handler.getUser(token)
	if err != nil {
		return err
	}

	jsonBytes, err := json.Marshal(&user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(api.Response[any]{
			Error: &api.ErrorResponse{
				Message: "Data encoding error",
				Details: err.Error(),
				Code:    fiber.StatusInternalServerError,
			},
			Status: api.StatusError,
		})
	}

	err = handler.setCookie(c, CookiePayload{
		Token:     "user",
		Value:     jsonBytes,
		KeyLookup: "cookie:user_session",
		Expires:   time.Duration(token.ExpiresIn),
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(api.Response[any]{
			Error: &api.ErrorResponse{
				Message: "Data encoding error",
				Details: err.Error(),
				Code:    fiber.StatusInternalServerError,
			},
			Status: api.StatusError,
		})
	}

	return c.Redirect("/")
}

func (handler *Handler) getUser(token *oauth2.Token) (ResponseGoogle, error) {
	client := configs.ConfigGoogle().Client(context.Background(), token)

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
