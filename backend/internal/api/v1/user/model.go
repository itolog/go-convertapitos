package user

import (
	"github.com/itolog/go-convertapitos/backend/pkg/db"
)

type User struct {
	db.Model
	Name          string `json:"name"`
	Email         string `json:"email" gorm:"uniqueIndex"`
	VerifiedEmail bool   `json:"verifiedEmail"`
	Password      string `json:"password,omitempty"`
	Picture       string `json:"picture"`
}
