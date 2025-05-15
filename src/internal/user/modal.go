package user

import (
	"github.com/itolog/go-convertapitos/src/pkg/db"
)

type User struct {
	db.Model
	Name          string `json:"name"`
	Email         string `json:"email" gorm:"uniqueIndex"`
	VerifiedEmail bool   `json:"verified_email"`
	Password      string `json:"password"`
	Picture       string `json:"picture"`
}
