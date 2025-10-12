package jwt

import (
	"errors"
	"gilangnyan/point-of-sales/internal/auth/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTServiceImpl struct {
	secretKey      string
	issuer         string
	expirationTime time.Duration
	refreshExpTime time.Duration
}

func (j *JWTServiceImpl) GenerateToken(userID string, email string, username string, roles []string) (*model.LoginResponse, error) {
	expiresAt := time.Now().Add(j.expirationTime)
	refreshExpiresAt := time.Now().Add(j.refreshExpTime)

	claims := &model.JWTClaims{
		UserID:   userID,
		Email:    email,
		Username: username,
		Roles:    roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.issuer,
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return nil, err
	}

	refreshClaims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(refreshExpiresAt),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    j.issuer,
		Subject:   userID,
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(j.secretKey))
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Token:        accessToken,
		RefreshToken: refreshTokenString,
		ExpiresAt:    expiresAt,
		User: model.UserInfo{
			ID:       userID,
			Username: username,
			Email:    email,
			Roles:    roles,
		},
	}, nil
}

func (j *JWTServiceImpl) ValidateToken(tokenString string) (*model.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*model.JWTClaims)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	return claims, nil
}

func (j *JWTServiceImpl) RefreshToken(refreshTokenString string) (*model.LoginResponse, error) {
	panic("unimplemented")
}

func NewJWTService(secretKey, issuer string, expTime, refreshExpTime time.Duration) JWTService {
	return &JWTServiceImpl{
		secretKey:      secretKey,
		issuer:         issuer,
		expirationTime: expTime,
		refreshExpTime: refreshExpTime,
	}
}
