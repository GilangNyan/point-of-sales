package model

import "gilangnyan/point-of-sales/package/template"

type UserRole struct {
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
	template.Base
}

type UserWithRoles struct {
	UserID   string   `json:"user_id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}
