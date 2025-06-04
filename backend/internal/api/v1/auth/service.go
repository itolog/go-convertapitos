package auth

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/backend/common"
	"github.com/itolog/go-convertapitos/backend/internal/api/v1/user"
	"github.com/itolog/go-convertapitos/backend/pkg/api"
	"github.com/itolog/go-convertapitos/backend/pkg/authorization"
	"github.com/markbates/goth"
	"golang.org/x/crypto/bcrypt"
)

const UserStoreKey = "user"

type IAuthService interface {
	Login(ctx *fiber.Ctx, payload *LoginRequest) (*common.AuthResponse, error)
	Register(payload *RegisterRequest) (*common.AuthResponse, error)
	Logout(ctx *fiber.Ctx)
	RefreshToken(ctx *fiber.Ctx, refreshToken string) (*common.RefreshResponse, error)
	OAuthCallback(ctx *fiber.Ctx, userInfo goth.User) (*common.AuthResponse, error)
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

	err = service.SaveUser(ctx, common.AuthResponse{
		AccessToken: accessToken,
		User:        existedUser,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &common.AuthResponse{
		AccessToken: accessToken,
		User:        existedUser,
	}, nil
}

func (service *Service) Register(payload *RegisterRequest) (*common.AuthResponse, error) {
	existedUser, _ := service.UserService.FindByEmail(payload.Email)
	if existedUser != nil {
		return nil, fiber.NewError(fiber.StatusConflict, api.ErrUserAlreadyExist)
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

	created.Password = ""
	return &common.AuthResponse{
		User: created,
	}, nil
}

func (service *Service) Logout(ctx *fiber.Ctx) {
	service.Authorization.SetCookie(ctx, authorization.CookiePayload{
		Name:    authorization.CookieTokenKey,
		Value:   "",
		Expires: 0,
	})

	service.Authorization.SetCookie(ctx, authorization.CookiePayload{
		Name:    UserStoreKey,
		Value:   "",
		Expires: 0,
	})
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

	err = service.SaveUser(ctx, common.AuthResponse{
		AccessToken: accessToken,
		User:        existedUser,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &common.RefreshResponse{
		AccessToken: accessToken,
	}, nil
}

func (service *Service) OAuthCallback(ctx *fiber.Ctx, userInfo goth.User) (*common.AuthResponse, error) {
	var userData *user.User

	existedUser, _ := service.UserService.FindByEmail(userInfo.Email)
	if existedUser != nil {
		userData = existedUser
	} else {
		createdUser, err := service.UserService.Create(user.User{
			Name:    userInfo.Name,
			Email:   userInfo.Email,
			Picture: userInfo.AvatarURL,
		})
		if err != nil {
			return nil, err
		}

		userData = createdUser
	}

	accessToken, err := service.Authorization.SetAuth(ctx, userInfo.Email)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	err = service.SaveUser(ctx, common.AuthResponse{
		AccessToken: accessToken,
		User:        userData,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &common.AuthResponse{
		AccessToken: accessToken,
		User:        userData,
	}, nil
}

func (service *Service) SaveUser(ctx *fiber.Ctx, userData common.AuthResponse) error {
	authData := common.StoredUser{
		AccessToken: userData.AccessToken,
		User: &common.UserInfo{
			Name:    userData.User.Name,
			Email:   userData.User.Email,
			Picture: userData.User.Picture,
		},
	}

	userDataJSON, err := json.Marshal(authData)
	if err != nil {
		return err
	}

	service.Authorization.SetCookie(ctx, authorization.CookiePayload{
		Name:     UserStoreKey,
		Value:    string(userDataJSON),
		Expires:  service.Authorization.JWT.RefreshTokenExpires,
		HTTPOnly: false,
	})

	return nil
}
