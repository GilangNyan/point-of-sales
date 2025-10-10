package repository

import (
	"context"
	"database/sql"
	"gilangnyan/point-of-sales/internal/user/model"
	"gilangnyan/point-of-sales/package/request"
)

type UserRepository interface {
	FindAll(ctx context.Context, params *request.PaginationParams) ([]*model.UserWithProfile, int64, error)
	FindByID(ctx context.Context, id string) (*model.UserWithProfile, error)
	FindByUsername(ctx context.Context, username string) (*model.UserWithProfile, error)
	FindByEmail(ctx context.Context, email string) (*model.UserWithProfile, error)
	Create(ctx context.Context, tx *sql.Tx, data model.User) (string, error)
	Update(ctx context.Context, tx *sql.Tx, id string, data model.User) (string, error)
	Delete(ctx context.Context, tx *sql.Tx, id string) error
}
