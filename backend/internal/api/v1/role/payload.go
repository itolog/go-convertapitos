package role

// CreateRequest represents the request body for creating a role
// @Description Request payload for creating a new role
type CreateRequest struct {
	Name        string       `json:"name" validate:"required,max=70" example:"manager"`
	Permissions []Permission `json:"permissions" validate:"required,min=1,dive"`
}

// UpdateRequest represents the request body for updating a role
// @Description Request payload for updating an existing role
type UpdateRequest struct {
	Name        string       `json:"name" validate:"required,max=70" example:"manager"`
	Permissions []Permission `json:"permissions" validate:"required,min=1,dive"`
}
