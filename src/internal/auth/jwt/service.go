package jwt

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/internal/user"
	"github.com/itolog/go-convertapitos/src/pkg/api"
	"github.com/itolog/go-convertapitos/src/pkg/jwt"
	"github.com/itolog/go-convertapitos/src/pkg/req"
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

func (service *Service) login(c *fiber.Ctx) error {
	payload, err := req.DecodeBody[LoginRequest](c)
	if err != nil {
		return err
	}
	validateError, valid := req.ValidateBody(payload)
	if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(api.Response{
			Error:  validateError,
			Status: api.StatusError,
		})
	}

	existedUser, _ := service.UserRepository.FindByEmail(payload.Email)
	if existedUser == nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Response{
			Error: &api.ErrorResponse{
				Message: api.ErrWrongCredentials,
			},
			Status: api.StatusError,
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(payload.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Response{
			Error: &api.ErrorResponse{
				Message: api.ErrWrongCredentials,
			},
			Status: api.StatusError,
		})
	}
	existedUser.Password = ""

	return c.Status(fiber.StatusOK).JSON(api.Response{
		Data:   existedUser,
		Status: api.StatusSuccess,
	})
}

func (service *Service) register(payload *RegisterRequest) (AuthResponse, error) {
	existedUser, _ := service.UserRepository.FindByEmail(payload.Email)
	if existedUser != nil {
		return AuthResponse{}, fmt.Errorf(api.ErrUserAlreadyExist)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return AuthResponse{}, err
	}

	created, err := service.UserRepository.Create(&user.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return AuthResponse{}, err
	}

	created.Password = ""
	return AuthResponse{
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
