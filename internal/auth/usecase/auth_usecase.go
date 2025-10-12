package usecase

import (
	"context"
	"gilangnyan/point-of-sales/internal/auth/model"
)

type AuthUsecase interface {
	Login(ctx context.Context, req model.LoginDto) (*model.LoginResponse, error)
	ActivateUser(ctx context.Context, req model.ActivateUserDto) error
}
