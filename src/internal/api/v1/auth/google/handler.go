package google

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/src/configs"
	"github.com/itolog/go-convertapitos/src/pkg/api"
	"github.com/rs/zerolog"
)

type HandlerDeps struct {
	GoogleService IGoogleService
	CustomLogger  *zerolog.Logger
}

type Handler struct {
	GoogleService IGoogleService
	CustomLogger  *zerolog.Logger
}

func NewHandler(router fiber.Router, deps HandlerDeps) {
	handler := Handler{
		GoogleService: deps.GoogleService,
		CustomLogger:  deps.CustomLogger,
	}

	router.Get("/google", handler.GoogleLogin)

	router.Get("/google/callback", handler.GoogleCallback)
}

// GoogleLogin initiates Google authentication process.
//
//	@Summary		Google Auth Login
//	@Description	Redirects the user to the Google OAuth consent page.
//	@Tags			Auth Google
//	@Produce		json
//	@Success		302	{string}	string	"Redirect to Google login"
//	@Router			/auth/google [get]
func (h *Handler) GoogleLogin(ctx *fiber.Ctx) error {
	from := ctx.Query("from", "/")
	path := configs.ConfigGoogle()
	url := path.AuthCodeURL(from)

	return ctx.Redirect(url)
}

// GoogleCallback handles the OAuth callback from Google.
//
//	@Summary		Google Auth Callback
//	@Description	Handles OAuth callback and authenticates/creates user account using Google data.
//	@Tags			Auth Google
//	@Accept			json
//	@Produce		json
//	@Param			code	query		string										true	"OAuth authorization code from Google"
//	@Success		200		{object}	api.ResponseData{data=common.AuthResponse}	"OAuth success, token and user info"
//	@Failure		400		{object}	api.ResponseError							"Failed to authenticate"
//	@Router			/auth/google/callback [get]
func (h *Handler) GoogleCallback(ctx *fiber.Ctx) error {
	code := ctx.FormValue("code")

	token, err := configs.ConfigGoogle().Exchange(ctx.Context(), code)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	createdUser, err := h.GoogleService.Callback(ctx, token)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(api.Response{
		Data:   createdUser,
		Status: api.StatusSuccess,
	})
}
