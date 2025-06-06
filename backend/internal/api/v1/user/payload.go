package user

type CreateRequest struct {
	Name          string `json:"name" validate:"required,max=70"`
	Email         string `json:"email" validate:"required,email"`
	VerifiedEmail bool   `json:"verifiedEmail" validate:"boolean"`
	Password      string `json:"password" validate:"required,min=6,max=128"`
	Picture       string `json:"picture"`
}

type UpdateRequest struct {
	Name          string     `json:"name" validate:"omitempty,max=70"`
	Email         string     `json:"email" validate:"omitempty,email"`
	VerifiedEmail bool       `json:"verifiedEmail" validate:"omitempty,boolean"`
	Password      string     `json:"password" validate:"omitempty,min=6,max=128"`
	Picture       string     `json:"picture"`
	Role          RoleType   `json:"role" validate:"omitempty"`
	AuthMethod    AuthMethod `json:"authMethod"`
}

type FindAllResponse struct {
	Users []User
	Count *int64
}

type BatchDeleteRequest struct {
	Ids []string `json:"ids" validate:"required,min=1,dive,uuid"`
}
