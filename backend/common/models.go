package common

import (
	"github.com/itolog/go-convertapitos/backend/internal/api/v1/user"
)

type AuthResponse struct {
	AccessToken *string    `json:"accessToken"`
	User        *user.User `json:"user"`
}

type RefreshResponse struct {
	AccessToken *string `json:"accessToken"`
}
