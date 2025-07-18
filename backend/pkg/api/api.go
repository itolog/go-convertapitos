package api

import "github.com/itolog/go-convertapitos/backend/pkg/validation"

type StatusType string

const (
	StatusSuccess StatusType = "success"
	StatusError   StatusType = "error"
)

type ErrorResponse struct {
	Code    int                      `json:"code,omitempty"`
	Message string                   `json:"message"`
	Details string                   `json:"details,omitempty"`
	Fields  []validation.ErrorFields `json:"fields,omitempty"`
}

type Response struct {
	Data   any            `json:"data,omitempty"`
	Error  *ErrorResponse `json:"error,omitempty"`
	Meta   *Meta          `json:"meta,omitempty"`
	Status StatusType     `json:"status"`
}

type Meta struct {
	Items int64 `json:"items"`
	Pages int   `json:"pages"`
}

type ResponseData struct {
	Data   any        `json:"data,omitempty"`
	Meta   *Meta      `json:"meta,omitempty"`
	Status StatusType `json:"status"`
}

type ResponseError struct {
	Error  *ErrorResponse `json:"error,omitempty"`
	Status StatusType     `json:"status"`
}
