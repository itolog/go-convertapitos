package user

type CreateRequest struct {
	Name          string `json:"name" validate:"required,max=70"`
	Email         string `json:"email" validate:"required,email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}
