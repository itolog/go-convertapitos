package google

import "github.com/gofiber/fiber/v2/middleware/session"

type GoogleResponse struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Verified   bool   `json:"verified_email"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
}

// SessionStore app wide session store
var SessionStore *session.Store
