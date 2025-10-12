package model

import "gilangnyan/point-of-sales/package/template"

type UserRole struct {
	UserID string `json:"userId"`
	RoleID string `json:"roleId"`
	template.Base
}

type UserWithRoles struct {
	UserID   string   `json:"userId"`
	Email    string   `json:"email"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}
