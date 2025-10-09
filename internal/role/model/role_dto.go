package model

type CreateRoleDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateRoleDto struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
