package model

import "gilangnyan/point-of-sales/package/template"

type UserRole struct {
	UserID string `json:"userId"`
	RoleID string `json:"roleId"`
	template.Base
}

type UserWithRoles struct {
	UserID   string   `json:"userId"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}
