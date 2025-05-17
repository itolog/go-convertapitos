package google

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/common"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/internal/user"
	"github.com/itolog/go-convertapitos/src/pkg/api"
	"github.com/itolog/go-convertapitos/src/pkg/authorization"

	"golang.org/x/oauth2"
	"io"
)

const userUrl = "https://www.googleapis.com/oauth2/v1/userinfo"

type ServiceDeps struct {
	UserService   *user.Service
	Authorization *authorization.Authorization
}
type Service struct {
	UserService   *user.Service
	Authorization *authorization.Authorization
}

func NewService(deps ServiceDeps) *Service {
	return &Service{
		UserService:   deps.UserService,
		Authorization: deps.Authorization,
	}
}

func (service *Service) callback(ctx *fiber.Ctx, token *oauth2.Token) (*common.AuthResponse, error) {
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

	accessToken, err := service.Authorization.SetAuth(ctx, userInfo.Email)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &common.AuthResponse{
		AccessToken: accessToken,
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
