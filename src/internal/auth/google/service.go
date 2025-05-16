package google

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/internal/user"
	"github.com/itolog/go-convertapitos/src/pkg/api"
	"github.com/itolog/go-convertapitos/src/pkg/cookie"
	"golang.org/x/oauth2"
	"io"
	"time"
)

const userUrl = "https://www.googleapis.com/oauth2/v1/userinfo"

type ServiceDeps struct {
	UserRepository *user.Repository
}
type Service struct {
	UserRepository *user.Repository
}

func NewService(deps ServiceDeps) *Service {
	return &Service{
		UserRepository: deps.UserRepository,
	}
}

func (service *Service) login(c *fiber.Ctx) error {
	from := c.Query("from", "/")
	path := configs.ConfigGoogle()
	url := path.AuthCodeURL(from)

	return c.Redirect(url)
}

func (service *Service) callback(c *fiber.Ctx) error {
	code := c.FormValue("code")
	from := c.Query("state")

	token, err := configs.ConfigGoogle().Exchange(c.Context(), code)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(api.Response{
			Error: &api.ErrorResponse{
				Message: "Unauthorized",
				Details: err.Error(),
				Code:    fiber.StatusUnauthorized,
			},
			Status: api.StatusError,
		})
	}

	userInfo, err := service.getUser(token)
	if err != nil {
		return err
	}

	jsonBytes, err := json.Marshal(&userInfo)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(api.Response{
			Error: &api.ErrorResponse{
				Message: "Data encoding error",
				Details: err.Error(),
				Code:    fiber.StatusInternalServerError,
			},
			Status: api.StatusError,
		})
	}

	cookieStore := cookie.NewCookie()
	err = cookieStore.SetCookie(c, cookie.Payload{
		Key:        "user",
		Value:      jsonBytes,
		CookieName: "cookie:user_session",
		Expires:    time.Duration(token.ExpiresIn),
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(api.Response{
			Error: &api.ErrorResponse{
				Message: "Data encoding error",
				Details: err.Error(),
				Code:    fiber.StatusInternalServerError,
			},
			Status: api.StatusError,
		})
	}

	return c.Redirect(from)
}

func (service *Service) getUser(token *oauth2.Token) (ResponseGoogle, error) {
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

	var userInfo ResponseGoogle

	err = json.NewDecoder(response.Body).Decode(&userInfo)
	if err != nil {
		return ResponseGoogle{}, fmt.Errorf("%v", err)
	}
	return userInfo, nil
}
