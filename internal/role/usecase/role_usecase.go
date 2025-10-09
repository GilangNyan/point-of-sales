package usecase

import (
	"context"
	"gilangnyan/point-of-sales/internal/role/model"
	"gilangnyan/point-of-sales/package/request"
	"gilangnyan/point-of-sales/package/response"
)

type RoleUsecase interface {
	FindAll(ctx context.Context, params *request.PaginationParams) (*response.PaginationResponse[*model.Role], error)
	Create(ctx context.Context, data model.CreateRoleDto) (string, error)
	Update(ctx context.Context, id string, data model.UpdateRoleDto) (string, error)
	Delete(ctx context.Context, id string) error
}
