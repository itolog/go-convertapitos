package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string `json:"name"`
	Email         string `json:"email" gorm:"uniqueIndex"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

func NewUser(name string, email string) *User {
	return &User{
		Name:  name,
		Email: email,
	}
}
