package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/backend/middleware"
	"github.com/itolog/go-convertapitos/backend/pkg/api"
	"github.com/itolog/go-convertapitos/backend/pkg/req"
	"github.com/rs/zerolog"
	"github.com/shareed2k/goth_fiber"
)

type HandlerDeps struct {
	AuthService  IAuthService
	CustomLogger *zerolog.Logger
}

type Handler struct {
	AuthService  IAuthService
	CustomLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, deps HandlerDeps) {
	handler := &Handler{
		AuthService:  deps.AuthService,
		CustomLogger: deps.CustomLogger,
	}

	setupOAuthProviders()

	router.Post("/login", handler.Login)
	router.Post("/register", handler.Register)
	router.Post("/logout", handler.Logout)
	router.Post("/refresh-token", middleware.Protected(), handler.RefreshToken)

	router.Get("/:provider", handler.OAuthLogin)
	router.Get("/:provider/callback", handler.OAuthCallback)
}

// Login handles user authentication.
//
//	@Summary		User login
//	@Description	Authenticate user with email and password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		LoginRequest								true	"User credentials"
//	@Success		200		{object}	api.ResponseData{data=common.AuthResponse}	"Successfully authenticated"
//	@Failure		400		{object}	api.ResponseError							"Invalid request or credentials"
//	@Router			/auth/login [post]
func (h *Handler) Login(ctx *fiber.Ctx) error {
	payload, err := req.DecodeBody[LoginRequest](ctx)
	if err != nil {
		return err
	}

	validateError, valid := req.ValidateBody(payload)
	if !valid {
		return ctx.Status(fiber.StatusBadRequest).JSON(api.Response{
			Error:  validateError,
			Status: api.StatusError,
		})
	}

	userInfo, err := h.AuthService.Login(ctx, payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(api.Response{
		Data:   userInfo,
		Status: api.StatusSuccess,
	})
}

// Register handles user registration.
//
//	@Summary		User registration
//	@Description	Register a new user with email and password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		RegisterRequest								true	"User registration data"
//	@Success		200		{object}	api.ResponseData{data=common.AuthResponse}	"Successfully registered"
//	@Failure		400		{object}	api.ResponseError							"Invalid request or registration error"
//	@Router			/auth/register [post]
func (h *Handler) Register(ctx *fiber.Ctx) error {
	payload, err := req.DecodeBody[RegisterRequest](ctx)
	if err != nil {
		return err
	}

	validateError, valid := req.ValidateBody(payload)
	if !valid {
		return ctx.Status(fiber.StatusBadRequest).JSON(api.Response{
			Error:  validateError,
			Status: api.StatusError,
		})
	}

	data, err := h.AuthService.Register(ctx, payload)
	if err != nil {
		statusCode := api.GetErrorCode(err)
		return fiber.NewError(statusCode, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(api.Response{
		Data:   data,
		Status: api.StatusSuccess,
	})
}

// RefreshToken handles refresh token requests.
//
//	@Summary		Refresh JWT token
//	@Description	Refresh access token using refresh token cookie
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	api.ResponseData{data=common.RefreshResponse}	"Token refreshed successfully"
//	@Failure		401	{object}	api.ResponseError								"Unauthorized or invalid refresh token"
//	@Router			/auth/refresh-token [post]
func (h *Handler) RefreshToken(ctx *fiber.Ctx) error {
	refreshToken := ctx.Cookies("refreshToken")
	if refreshToken == "" {
		return fiber.NewError(fiber.StatusUnauthorized, api.ErrUnauthorized)
	}

	user, err := h.AuthService.RefreshToken(ctx, refreshToken)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(api.Response{
		Data:   user,
		Status: api.StatusSuccess,
	})
}

// Logout godoc
//
//	@Summary		Logout user
//	@Description	Performs logout by invalidating user's authentication (such as token or session)
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	api.ResponseData{data=string}	"Logout successful"
//	@Failure		401	{object}	api.ResponseError				"Unauthorized"
//	@Router			/auth/logout [post]
func (h *Handler) Logout(ctx *fiber.Ctx) error {
	h.AuthService.Logout(ctx)

	return ctx.Status(fiber.StatusOK).JSON(api.Response{
		Data:   "User logged out.",
		Status: api.StatusSuccess,
	})
}

// OAuthLogin initiates OAuth authorization
//
//	@Summary		OAuth login
//	@Description	Start OAuth authentication with provider (google, github, ...)
//	@Tags			Auth
//	@Param			provider	path		string	true	"OAuth provider (google, github, ...)"
//	@Success		302			{object}	nil		"Redirect to OAuth provider"
//	@Failure		400			{object}	api.ResponseError
//	@Router			/auth/{provider} [get]
func (h *Handler) OAuthLogin(ctx *fiber.Ctx) error {
	return goth_fiber.BeginAuthHandler(ctx)
}

// OAuthCallback handles OAuth callback from provider
//
//	@Summary		OAuth callback
//	@Description	Handle OAuth callback and complete authentication
//	@Tags			Auth
//	@Param			provider	path		string										true	"OAuth provider (google, github, ...)"
//	@Param			code		query		string										true	"Authorization code from provider"
//	@Success		200			{object}	api.ResponseData{data=common.AuthResponse}	"Successfully authenticated"
//	@Failure		400			{object}	api.ResponseError							"Authentication failed"
//	@Router			/auth/{provider}/callback [get]
func (h *Handler) OAuthCallback(ctx *fiber.Ctx) error {
	user, err := goth_fiber.CompleteUserAuth(ctx)
	if err != nil {
		h.CustomLogger.Error().Err(err).Msg("Error completing user auth")
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(api.Response{
		Data:   user,
		Status: api.StatusSuccess,
	})
}
