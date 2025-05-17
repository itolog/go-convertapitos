package common

import "github.com/itolog/go-convertapitos/src/internal/user"

type AuthResponse struct {
	AccessToken *string    `json:"access_token"`
	User        *user.User `json:"user"`
}
