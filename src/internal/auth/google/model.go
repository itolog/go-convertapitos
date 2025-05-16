package google

import "github.com/itolog/go-convertapitos/src/internal/user"

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type ResponseGoogle struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Verified bool   `json:"verified_email"`
	Name     string `json:"name"`

	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
}

type RegistrationResponse struct {
	AccessToken *string    `json:"access_token"`
	User        *user.User `json:"user"`
}
