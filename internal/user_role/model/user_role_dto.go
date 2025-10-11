package model

type AssignRolesDto struct {
	UserID  string   `json:"userId" binding:"required,uuid"`
	RoleIDs []string `json:"roleIds" binding:"required,uuid"`
}
