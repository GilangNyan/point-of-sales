package model

type CreateUserDto struct {
	Username    string  `json:"username" binding:"required"`
	Email       string  `json:"email" binding:"required,email"`
	Password    string  `json:"password" binding:"required,min=6"`
	FullName    string  `json:"fullName" binding:"required"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	DateOfBirth *string `json:"dateOfBirth,omitempty"`
	Address     *string `json:"address,omitempty"`
}

type UpdateUserDto struct {
	Username       *string `json:"username,omitempty"`
	Email          *string `json:"email,omitempty"`
	IsActive       *bool   `json:"isActive,omitempty"`
	IsBlocked      *bool   `json:"isBlocked,omitempty"`
	FullName       *string `json:"fullName,omitempty"`
	PhoneNumber    *string `json:"phoneNumber,omitempty"`
	DateOfBirth    *string `json:"dateOfBirth,omitempty"`
	Address        *string `json:"address,omitempty"`
	ProfilePicture *string `json:"profilePicture,omitempty"`
}

type UpdateUserPasswordDto struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=6"`
}
