package api

type StatusType string

const (
	StatusSuccess StatusType = "success"
	StatusError   StatusType = "error"
)

type ErrorResponse struct {
	Code    uint16 `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Details string `json:"details,omitempty"`
}

type Response[T any] struct {
	Data   T             `json:"data"`
	Error  ErrorResponse `json:"error,omitempty"`
	Status StatusType    `json:"status,omitempty"`
}
