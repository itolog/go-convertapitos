package api

import "github.com/itolog/go-convertapitos/src/pkg/validation"

type StatusType string

const (
	StatusSuccess StatusType = "success"
	StatusError   StatusType = "error"
)

type ErrorResponse struct {
	Code    uint16                   `json:"code"`
	Message string                   `json:"message"`
	Details string                   `json:"details"`
	Fields  []validation.ErrorFields `json:"fields"`
}

type Response[T any] struct {
	Data   T              `json:"data,omitempty"`
	Error  *ErrorResponse `json:"error,omitempty"`
	Status StatusType     `json:"status"`
}
