package role

import (
	"database/sql/driver"
	"fmt"
	"github.com/goccy/go-json"

	"github.com/itolog/go-convertapitos/backend/pkg/db"
)

// Permission represents a single permission for an entity
// @Description Permission object with CRUD operations for a specific entity
type Permission struct {
	Entity string `json:"entity" validate:"required,max=70" example:"users"`
	Read   bool   `json:"read" example:"true"`
	Create bool   `json:"create" example:"true"`
	Update bool   `json:"update" example:"false"`
	Delete bool   `json:"delete" example:"false"`
}

// PermissionsList represents a list of permissions
type PermissionsList []Permission

// Value implements driver.Valuer interface for writing to database
func (p PermissionsList) Value() (driver.Value, error) {
	if len(p) == 0 {
		return "[]", nil
	}
	return json.Marshal(p)
}

// Scan implements sql.Scanner interface for reading from database
func (p *PermissionsList) Scan(value interface{}) error {
	if value == nil {
		*p = PermissionsList{}
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("cannot scan %T into PermissionsList", value)
	}

	return json.Unmarshal(bytes, p)
}

// Role represents a user role with permissions
// @Description Role object containing name and associated permissions
type Role struct {
	db.Model
	Name        string          `json:"name" gorm:"uniqueIndex" example:"admin"`
	Permissions PermissionsList `json:"permissions" gorm:"type:jsonb" swaggertype:"array,object"`
}
