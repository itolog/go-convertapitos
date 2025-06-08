package user

import (
	"github.com/google/uuid"
	"github.com/itolog/go-convertapitos/backend/internal/api/v1/role"
	"github.com/itolog/go-convertapitos/backend/pkg/db"
	"time"
)

type User struct {
	db.Model
	Name          string     `json:"name"`
	Email         string     `json:"email" gorm:"uniqueIndex"`
	VerifiedEmail bool       `json:"verifiedEmail" gorm:"default:false"`
	Password      string     `json:"password,omitempty"`
	Picture       string     `json:"picture"`
	RoleID        string     `json:"roleId"`
	Role          role.Role  `json:"role" gorm:"foreignKey:RoleID"`
	AuthMethod    AuthMethod `json:"authMethod"`
	Accounts      []Account  `json:"accounts" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Account struct {
	db.Model
	Provider     string    `json:"provider"`
	ProviderID   string    `json:"providerId"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
	UserID       uuid.UUID
}
