package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

// Auth Errors
const (
	ErrUserAlreadyExist = "user already exists"
	ErrWrongCredentials = "wrong email or password"
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
