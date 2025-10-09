package repository

import (
	"context"
	"gilangnyan/point-of-sales/internal/role/model"
	"gilangnyan/point-of-sales/package/request"
)

type RoleRepository interface {
	FindAll(ctx context.Context, params *request.PaginationParams) ([]*model.Role, int64, error)
	FindByID(ctx context.Context, id string) (*model.Role, error)
	Create(ctx context.Context, data model.Role) (string, error)
	Update(ctx context.Context, id string, data model.Role) (string, error)
	Delete(ctx context.Context, id string) error
}
