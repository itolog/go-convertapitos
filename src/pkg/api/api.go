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

type DataResponse struct {
	Data   any           `json:"data"`
	Error  ErrorResponse `json:"error"`
	Status StatusType    `json:"status,omitempty"`
}
