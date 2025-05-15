package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokens struct {
	AccessToken  string
	RefreshToken string
}

type Deps struct {
	Secret              string
	AccessTokenExpires  time.Duration
	RefreshTokenExpires time.Duration
}
type JWT struct {
	Secret              string
	AccessTokenExpires  time.Duration
	RefreshTokenExpires time.Duration
}

func NewJWT(des Deps) *JWT {
	return &JWT{
		Secret:              des.Secret,
		AccessTokenExpires:  des.AccessTokenExpires,
		RefreshTokenExpires: des.RefreshTokenExpires,
	}
}

func (j *JWT) Create(payload string, duration time.Duration) (string, error) {
	expirationTime := time.Now().Add(duration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": payload,
		"exp":   expirationTime.Unix(),
	})

	s, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *JWT) GenAccessTokens(payload string) (tokens *AccessTokens, err error) {
	accessToken, err := j.Create(payload, j.AccessTokenExpires)
	if err != nil {
		return &AccessTokens{}, err
	}

	refreshToken, err := j.Create(payload, j.RefreshTokenExpires)
	if err != nil {
		return &AccessTokens{}, err
	}

	return &AccessTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
