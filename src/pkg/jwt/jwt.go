package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/itolog/go-convertapitos/src/pkg/environments"
	"github.com/itolog/go-convertapitos/src/pkg/timeutils"
	"time"
)

type AccessTokens struct {
	AccessToken  string
	RefreshToken string
}

type JWT struct {
	Secret              string `env:"JWT_SECRET"`
	AccessTokenExpires  time.Duration
	RefreshTokenExpires time.Duration
}

func NewJWT() (*JWT, error) {
	accessTokenTTL := environments.GetEnv("JWT_ACCESS_TOKEN_TTL")
	refreshTokenTTL := environments.GetEnv("JWT_REFRESH_TOKEN_TTL")

	accessTokenDuration, err := timeutils.ParseDuration(accessTokenTTL)
	if err != nil {
		return nil, err
	}

	refreshTokenDuration, err := timeutils.ParseDuration(refreshTokenTTL)
	if err != nil {
		return nil, err
	}
	return &JWT{
		Secret:              environments.GetEnv("JWT_SECRET"),
		AccessTokenExpires:  accessTokenDuration,
		RefreshTokenExpires: refreshTokenDuration,
	}, nil
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
