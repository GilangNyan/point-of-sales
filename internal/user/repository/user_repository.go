package repository

import (
	"context"
	"gilangnyan/point-of-sales/internal/user/model"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]*model.User, error)
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, data model.User) (*model.User, error)
	Update(ctx context.Context, id string, data model.User) (*model.User, error)
	Delete(ctx context.Context, id string) error
}
