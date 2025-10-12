package model

type LoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ActivateUserDto struct {
	Email string `json:"email" binding:"required,email"`
	Token string `json:"token" binding:"required"`
}

type RefreshTokenDto struct {
	RefreshToken string `json:"refreshToken"`
}
