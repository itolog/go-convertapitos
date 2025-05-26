package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/backend/common"
	"github.com/itolog/go-convertapitos/backend/internal/api/v1/user"
	"github.com/itolog/go-convertapitos/backend/pkg/api"
	"github.com/itolog/go-convertapitos/backend/pkg/authorization"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(ctx *fiber.Ctx, payload *LoginRequest) (*common.AuthResponse, error)
	Register(ctx *fiber.Ctx, payload *RegisterRequest) (*common.AuthResponse, error)
	Logout(ctx *fiber.Ctx)
	RefreshToken(ctx *fiber.Ctx, refreshToken string) (*common.RefreshResponse, error)
}

type ServiceDeps struct {
	UserService   user.IUserService
	Authorization *authorization.Authorization
}
type Service struct {
	UserService   user.IUserService
	Authorization *authorization.Authorization
}

func NewService(deps ServiceDeps) *Service {
	return &Service{
		UserService:   deps.UserService,
		Authorization: deps.Authorization,
	}
}

func (service *Service) Login(ctx *fiber.Ctx, payload *LoginRequest) (*common.AuthResponse, error) {
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

	return &common.AuthResponse{
		AccessToken: accessToken,
		User:        existedUser,
	}, nil
}

func (service *Service) Register(ctx *fiber.Ctx, payload *RegisterRequest) (*common.AuthResponse, error) {
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
	return &common.AuthResponse{
		AccessToken: accessToken,
		User:        created,
	}, nil
}

func (service *Service) Logout(ctx *fiber.Ctx) {
	service.Authorization.SetCookie(ctx, "refreshToken", 0)
}

func (service *Service) RefreshToken(ctx *fiber.Ctx, refreshToken string) (*common.RefreshResponse, error) {
	verify, err := service.Authorization.VerifyToken(refreshToken)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	existedUser, _ := service.UserService.FindByEmail(verify.Email)
	if existedUser == nil {
		return nil, fiber.NewError(fiber.StatusNotFound, api.ErrWrongCredentials)
	}

	accessToken, err := service.Authorization.SetAuth(ctx, existedUser.Email)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	existedUser.Password = ""
	return &common.RefreshResponse{
		AccessToken: accessToken,
	}, nil
}
