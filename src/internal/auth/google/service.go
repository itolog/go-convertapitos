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
	"github.com/itolog/go-convertapitos/src/pkg/jwt"
	"golang.org/x/oauth2"
	"io"
)

const userUrl = "https://www.googleapis.com/oauth2/v1/userinfo"

type ServiceDeps struct {
	UserService *user.Service
}
type Service struct {
	UserService *user.Service
}

func NewService(deps ServiceDeps) *Service {
	return &Service{
		UserService: deps.UserService,
	}
}

func (service *Service) callback(ctx *fiber.Ctx, token *oauth2.Token) (*RegistrationResponse, error) {
	userInfo, err := service.getUser(token)
	if err != nil {
		return nil, err
	}

	existedUser, _ := service.UserService.FindByEmail(userInfo.Email)
	if existedUser != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, api.ErrUserAlreadyExist)
	}

	createdUser, err := service.UserService.Create(user.User{
		Name:          userInfo.Name,
		Email:         userInfo.Email,
		Picture:       userInfo.Picture,
		VerifiedEmail: userInfo.Verified,
	})
	if err != nil {
		return nil, err
	}

	jwtService, err := jwt.NewJWT()
	if err != nil {
		return nil, err
	}

	tokens, err := jwtService.GenAccessTokens(userInfo.Email)
	if err != nil {
		return nil, err
	}

	cookieStore := cookie.NewCookie()
	err = cookieStore.SetCookie(ctx, cookie.Payload{
		Key:        "refresh_token",
		Value:      tokens.RefreshToken,
		CookieName: "cookie:refresh_token",
		Expires:    jwtService.RefreshTokenExpires,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &RegistrationResponse{
		AccessToken: tokens.AccessToken,
		User:        createdUser,
	}, nil
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
