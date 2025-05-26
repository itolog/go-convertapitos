package req

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itolog/go-convertapitos/backend/pkg/api"
	"github.com/itolog/go-convertapitos/backend/pkg/validation"
)

func ValidateBody[T any](payload *T) (*api.ErrorResponse, bool) {
	validator := validation.NewValidator()
	validationErrors := validator.Validate(payload)

	if len(validationErrors) > 0 {
		return &api.ErrorResponse{
			Message: "Validation error",
			Code:    fiber.StatusBadRequest,
			Fields:  validationErrors,
		}, false
	}

	return &api.ErrorResponse{}, true
}
