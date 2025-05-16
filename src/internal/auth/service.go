package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/internal/user"
	"github.com/itolog/go-convertapitos/src/pkg/api"
	"github.com/itolog/go-convertapitos/src/pkg/authorization"

	"golang.org/x/crypto/bcrypt"
)

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

func (service *Service) Login(ctx *fiber.Ctx, payload *LoginRequest) (*Response, error) {
	existedUser, _ := service.UserService.FindByEmail(payload.Email)
	if existedUser == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, api.ErrWrongCredentials)
	}

	err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(payload.Password))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, api.ErrWrongCredentials)

	}
	existedUser.Password = ""

	accessToken, err := service.Authorization.SetAuth(ctx, payload.Email)
	if err != nil {
		return nil, err
	}

	return &Response{
		AccessToken: accessToken,
		User:        existedUser,
	}, nil
}

func (service *Service) register(ctx *fiber.Ctx, payload *RegisterRequest) (*Response, error) {
	existedUser, _ := service.UserService.FindByEmail(payload.Email)
	if existedUser != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, api.ErrUserAlreadyExist)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	created, err := service.UserService.Create(user.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return nil, err
	}

	accessToken, err := service.Authorization.SetAuth(ctx, payload.Email)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	created.Password = ""
	return &Response{
		AccessToken: accessToken,
		User:        created,
	}, nil
}
