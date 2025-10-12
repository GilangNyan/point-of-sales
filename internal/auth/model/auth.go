package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginResponse struct {
	Token        string    `json:"token"`
	RefreshToken string    `json:"refreshToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
	User         UserInfo  `json:"user"`
}

type UserInfo struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	FullName string   `json:"fullName"`
	Roles    []string `json:"roles"`
	IsActive bool     `json:"isActive"`
}

type JWTClaims struct {
	UserID   string   `json:"id"`
	Email    string   `json:"email"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}
