package usecase

import (
	"context"
	"gilangnyan/point-of-sales/internal/user/model"
)

type UserUsecase interface {
	FindAll(ctx context.Context) ([]*model.UserWithProfile, error)
	Create(ctx context.Context, data model.CreateUserDto) (string, error)
	Update(ctx context.Context, id string, data model.UpdateUserDto) (string, error)
	Delete(ctx context.Context, id string) error
}
