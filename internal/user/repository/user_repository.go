package repository

import (
	"context"
	"gilangnyan/point-of-sales/internal/user/model"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]*model.UserWithProfile, error)
	FindByID(ctx context.Context, id string) (*model.UserWithProfile, error)
	FindByUsername(ctx context.Context, username string) (*model.UserWithProfile, error)
	FindByEmail(ctx context.Context, email string) (*model.UserWithProfile, error)
	Create(ctx context.Context, data model.User) (string, error)
	Update(ctx context.Context, id string, data model.User) (string, error)
	Delete(ctx context.Context, id string) error
}
