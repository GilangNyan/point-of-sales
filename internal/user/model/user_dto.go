package model

type CreateUserDto struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UpdateUserDto struct {
	Username  *string `json:"username,omitempty"`
	Email     *string `json:"email,omitempty"`
	Password  *string `json:"password,omitempty"`
	IsActive  *bool   `json:"is_active,omitempty"`
	IsBlocked *bool   `json:"is_blocked,omitempty"`
}
