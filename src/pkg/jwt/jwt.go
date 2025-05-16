package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/itolog/go-convertapitos/src/pkg/environments"
	"github.com/itolog/go-convertapitos/src/pkg/timeutils"
	"time"
)

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

func (j *JWT) CreateToken(payload Payload, duration time.Duration) (string, error) {
	expirationTime := time.Now().Add(duration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	})

	s, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *JWT) Verify(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	claim, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claim, nil
}

func (j *JWT) GenAccessTokens(payload Payload) (tokens *AccessTokens, err error) {
	accessToken, err := j.CreateToken(payload, j.AccessTokenExpires)
	if err != nil {
		return &AccessTokens{}, err
	}

	refreshToken, err := j.CreateToken(payload, j.RefreshTokenExpires)
	if err != nil {
		return &AccessTokens{}, err
	}

	return &AccessTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
