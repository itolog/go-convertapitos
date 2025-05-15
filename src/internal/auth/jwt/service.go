package jwt

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/internal/user"
	"github.com/itolog/go-convertapitos/src/pkg/api"
	"github.com/itolog/go-convertapitos/src/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type ServiceDeps struct {
	*configs.Config
	UserRepository *user.Repository
}
type Service struct {
	*configs.Config
	UserRepository *user.Repository
}

func NewService(deps ServiceDeps) *Service {
	return &Service{
		UserRepository: deps.UserRepository,
		Config:         deps.Config,
	}
}

func (service *Service) login(payload *LoginRequest) (*user.User, error) {

	existedUser, _ := service.UserRepository.FindByEmail(payload.Email)
	if existedUser == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, api.ErrWrongCredentials)
	}

	err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(payload.Password))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, api.ErrWrongCredentials)

	}
	existedUser.Password = ""

	return existedUser, nil
}

func (service *Service) register(payload *RegisterRequest) (*AuthResponse, error) {
	existedUser, _ := service.UserRepository.FindByEmail(payload.Email)
	if existedUser != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, api.ErrUserAlreadyExist)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	created, err := service.UserRepository.Create(&user.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	created.Password = ""
	return &AuthResponse{
		AccessToken: "accessToken",
		User:        created,
	}, nil
}

func (service *Service) SetAuthTokens(payload string) {
	jwtService := jwt.NewJWT(jwt.Deps{
		Secret:              service.Auth.JwtSecret,
		AccessTokenExpires:  service.Auth.AccessTokenExpires,
		RefreshTokenExpires: service.Auth.RefreshTokenExpires,
	})

	tokens, err := jwtService.GenAccessTokens(payload)
	if err != nil {
		return
	}

	fmt.Println(tokens.RefreshToken)
}
