package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

// Auth Errors
const (
	ErrUserAlreadyExist      = "user already exists"
	ErrWrongCredentials      = "wrong email or password"
	ErrUnauthorized          = "unauthorized"
	ErrMustBeANumber         = "must be a number"
	ErrMissingOrMalformedJWT = "missing or malformed jwt"
	ErrInvalidToken          = "invalid or expired token"
)

func GetErrorCode(err error) int {
	var fiberErr *fiber.Error
	ok := errors.As(err, &fiberErr)
	statusCode := fiber.StatusInternalServerError

	if ok {
		statusCode = fiberErr.Code
	}

	return statusCode
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	if err != nil {
		return ctx.Status(code).JSON(Response{
			Error: &ErrorResponse{
				Message: err.Error(),
				Code:    code,
			},
			Status: StatusError,
		})
	}

	return nil
}
