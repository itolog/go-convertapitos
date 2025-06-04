package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/backend/middleware"
	"github.com/itolog/go-convertapitos/backend/pkg/api"
	"github.com/itolog/go-convertapitos/backend/pkg/environments"
	"github.com/itolog/go-convertapitos/backend/pkg/req"
	"github.com/rs/zerolog"
	"github.com/shareed2k/goth_fiber"
)

const RedirectKey = "redirect_to"

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
func (h *Handler) Login(c *fiber.Ctx) error {
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

	userInfo, err := h.AuthService.Login(c, payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(api.Response{
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
func (h *Handler) Register(c *fiber.Ctx) error {
	payload, err := req.DecodeBody[RegisterRequest](c)
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

	data, err := h.AuthService.Register(payload)
	if err != nil {
		statusCode := api.GetErrorCode(err)
		return fiber.NewError(statusCode, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(api.Response{
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
func (h *Handler) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refreshToken")
	if refreshToken == "" {
		return fiber.NewError(fiber.StatusUnauthorized, api.ErrUnauthorized)
	}

	user, err := h.AuthService.RefreshToken(c, refreshToken)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(api.Response{
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
func (h *Handler) Logout(c *fiber.Ctx) error {
	h.AuthService.Logout(c)

	return c.Status(fiber.StatusOK).JSON(api.Response{
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
func (h *Handler) OAuthLogin(c *fiber.Ctx) error {
	redirectTo := c.Query(RedirectKey)

	if redirectTo != "" {
		c.Cookie(&fiber.Cookie{
			Name:     RedirectKey,
			Value:    redirectTo,
			HTTPOnly: true,
			Secure:   !environments.IsDev(),
			SameSite: "Lax",
			MaxAge:   300,
		})
	}

	return goth_fiber.BeginAuthHandler(c)
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
func (h *Handler) OAuthCallback(c *fiber.Ctx) error {
	user, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		h.CustomLogger.Error().Err(err).Msg("Error completing user auth")
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_, err = h.AuthService.OAuthCallback(c, user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	redirectTo := c.Cookies(RedirectKey)
	if redirectTo == "" {
		redirectTo = "/"
	}

	c.Cookie(&fiber.Cookie{
		Name:     RedirectKey,
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	})

	return c.Redirect(redirectTo)

}
