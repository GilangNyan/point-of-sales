package usecase

import (
	"context"
	"gilangnyan/point-of-sales/internal/user/model"
	"gilangnyan/point-of-sales/package/request"
	"gilangnyan/point-of-sales/package/response"
)

type UserUsecase interface {
	FindAll(ctx context.Context, params *request.PaginationParams) (*response.PaginationResponse[*model.UserWithProfile], error)
	Create(ctx context.Context, data model.CreateUserDto) (string, error)
	Update(ctx context.Context, id string, data model.UpdateUserDto) (string, error)
	Delete(ctx context.Context, id string) error
}
