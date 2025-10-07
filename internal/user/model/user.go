package model

import (
	"gilangnyan/point-of-sales/package/template"
)

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsActive  bool   `json:"is_active"`
	IsBlocked bool   `json:"is_blocked"`
	template.Base
}

type UserProfile struct {
	ID             string  `json:"id"`
	FullName       string  `json:"full_name"`
	DateOfBirth    *string `json:"date_of_birth"`
	PhoneNumber    *string `json:"phone_number"`
	Address        *string `json:"address"`
	ProfilePicture *string `json:"profile_picture"`
	UserID         string  `json:"user_id"`
	template.Base
}

type UserWithProfile struct {
	ID             string  `json:"id"`
	Username       string  `json:"username"`
	Email          string  `json:"email"`
	FullName       string  `json:"full_name"`
	DateOfBirth    *string `json:"date_of_birth"`
	PhoneNumber    *string `json:"phone_number"`
	Address        *string `json:"address"`
	ProfilePicture *string `json:"profile_picture"`
	IsActive       bool    `json:"is_active"`
	IsBlocked      bool    `json:"is_blocked"`
}
