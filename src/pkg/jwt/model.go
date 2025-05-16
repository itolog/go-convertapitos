package jwt

import "github.com/golang-jwt/jwt/v5"

type Payload struct {
	Email string `json:"email"`
}
type UserClaims struct {
	Payload
	jwt.RegisteredClaims
}

type AccessTokens struct {
	AccessToken  string
	RefreshToken string
}
