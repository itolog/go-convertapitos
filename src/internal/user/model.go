package user

import (
	"github.com/itolog/go-convertapitos/src/pkg/db"
)

type User struct {
	db.Model
	Name          string `json:"name"`
	Email         string `json:"email" gorm:"uniqueIndex"`
	VerifiedEmail bool   `json:"verified_email"`
	Password      string `json:"password,omitempty"`
	Picture       string `json:"picture"`
}

type CreateRequest struct {
	Name          string `json:"name" validate:"required,max=70"`
	Email         string `json:"email" validate:"required,email"`
	VerifiedEmail bool   `json:"verified_email" validate:"boolean"`
	Password      string `json:"password" validate:"required,min=6,max=128"`
	Picture       string `json:"picture"`
}

type UpdateRequest struct {
	Name          string `json:"name" validate:"omitempty,max=70"`
	Email         string `json:"email" validate:"omitempty,email"`
	VerifiedEmail bool   `json:"verified_email" validate:"omitempty,boolean"`
	Password      string `json:"password" validate:"omitempty,min=6,max=128"`
	Picture       string `json:"picture"`
}

type FindAllResponse struct {
	Users []User
	Count *int64
}
