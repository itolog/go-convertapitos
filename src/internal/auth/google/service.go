package google

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/internal/user"
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

func (service *Service) callback(ctx *fiber.Ctx, token *oauth2.Token) error {
	userInfo, err := service.getUser(token)
	if err != nil {
		return err
	}

	jsonBytes, err := json.Marshal(&userInfo)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	cookieStore := cookie.NewCookie()
	err = cookieStore.SetCookie(ctx, cookie.Payload{
		Key:        "user",
		Value:      jsonBytes,
		CookieName: "cookie:user_session",
		Expires:    time.Duration(token.ExpiresIn),
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
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
