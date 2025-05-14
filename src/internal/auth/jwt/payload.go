package jwt

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=128"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
