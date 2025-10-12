package jwt

import "gilangnyan/point-of-sales/internal/auth/model"

type JWTService interface {
	GenerateToken(userID, email, username string, roles []string) (*model.LoginResponse, error)
	ValidateToken(tokenString string) (*model.JWTClaims, error)
	RefreshToken(refreshTokenString string) (*model.LoginResponse, error)
}
