package jwt

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=128"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,max=70"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=128"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type RegisterResponse struct {
	AccessToken string `json:"access_token"`
}
