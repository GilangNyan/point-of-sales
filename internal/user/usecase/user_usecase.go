package usecase

import (
	"context"
	"gilangnyan/point-of-sales/internal/user/model"
)

type UserUsecase interface {
	FindAll(ctx context.Context) ([]*model.User, error)
	Create(ctx context.Context, data model.CreateUserDto) (*model.User, error)
	Update(ctx context.Context, id string, data model.UpdateUserDto) (*model.User, error)
	Delete(ctx context.Context, id string) error
}
