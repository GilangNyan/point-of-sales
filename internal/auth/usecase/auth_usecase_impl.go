package usecase

import (
	"context"
	"errors"
	"gilangnyan/point-of-sales/internal/auth/model"
	"gilangnyan/point-of-sales/internal/auth/repository"
	userRoleRepo "gilangnyan/point-of-sales/internal/user_role/repository"
	"gilangnyan/point-of-sales/package/jwt"
	"gilangnyan/point-of-sales/package/utils"
)

type AuthUsecaseImpl struct {
	authRepo     repository.AuthRepository
	userRoleRepo userRoleRepo.UserRoleRepository
	jwtService   jwt.JWTService
}

func (a *AuthUsecaseImpl) Login(ctx context.Context, req model.LoginDto) (*model.LoginResponse, error) {
	user, err := a.userRoleRepo.FindUserByUsernameWithRoles(ctx, req.Username)
	if err != nil {
		return nil, errors.New("Invalid email or password")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("Invalid email or password")
	}

	loginResponse, err := a.jwtService.GenerateToken(
		user.UserID,
		user.Email,
		user.Username,
		user.Roles,
	)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}
	return loginResponse, nil
}

func (a *AuthUsecaseImpl) ActivateUser(ctx context.Context, req model.ActivateUserDto) error {
	panic("unimplemented")
}

func NewAuthUsecase(authRepo repository.AuthRepository, userRoleRepo userRoleRepo.UserRoleRepository, jwtService jwt.JWTService) AuthUsecase {
	return &AuthUsecaseImpl{
		authRepo:     authRepo,
		userRoleRepo: userRoleRepo,
		jwtService:   jwtService,
	}
}
