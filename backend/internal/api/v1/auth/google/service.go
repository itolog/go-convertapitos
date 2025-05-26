package google

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/backend/common"
	"github.com/itolog/go-convertapitos/backend/configs"
	"github.com/itolog/go-convertapitos/backend/internal/api/v1/user"
	"github.com/itolog/go-convertapitos/backend/pkg/api"
	"github.com/itolog/go-convertapitos/backend/pkg/authorization"
	"github.com/rs/zerolog"

	"golang.org/x/oauth2"
	"io"
)

const userUrl = "https://www.googleapis.com/oauth2/v1/userinfo"

type IGoogleService interface {
	Callback(ctx *fiber.Ctx, token *oauth2.Token) (*common.AuthResponse, error)
}

type ServiceDeps struct {
	UserService   user.IUserService
	Authorization *authorization.Authorization
	customLogger  zerolog.Logger
}
type Service struct {
	UserService   user.IUserService
	Authorization *authorization.Authorization
	customLogger  zerolog.Logger
}

func NewService(deps ServiceDeps) *Service {
	return &Service{
		UserService:   deps.UserService,
		Authorization: deps.Authorization,
		customLogger:  deps.customLogger,
	}
}

func (service *Service) Callback(ctx *fiber.Ctx, token *oauth2.Token) (*common.AuthResponse, error) {
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
			service.customLogger.Error().Msg(fmt.Sprintf("Error closing response body %v", err.Error()))
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
