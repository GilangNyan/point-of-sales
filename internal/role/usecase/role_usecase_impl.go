package usecase

import (
	"context"
	"errors"
	"gilangnyan/point-of-sales/internal/role/model"
	"gilangnyan/point-of-sales/internal/role/repository"
	"gilangnyan/point-of-sales/package/request"
	"gilangnyan/point-of-sales/package/response"
)

type RoleUsecaseImpl struct {
	repo repository.RoleRepository
}

func (r *RoleUsecaseImpl) FindAll(ctx context.Context, params *request.PaginationParams) (*response.PaginationResponse[*model.Role], error) {
	data, total, err := r.repo.FindAll(ctx, params)
	if err != nil {
		return nil, errors.New("failed to retrieve roles")
	}
	return response.NewPaginationResponse(data, total, params.Page, params.PageSize), nil
}

func (r *RoleUsecaseImpl) Create(ctx context.Context, data model.CreateRoleDto) (string, error) {
	role := &model.Role{
		Name:        data.Name,
		Description: data.Description,
	}

	createdRoleId, err := r.repo.Create(ctx, *role)
	if err != nil {
		return "", errors.New("failed to create role")
	}
	return createdRoleId, nil
}

func (r *RoleUsecaseImpl) Update(ctx context.Context, id string, data model.UpdateRoleDto) (string, error) {
	checkRole, err := r.repo.FindByID(ctx, id)
	if err != nil {
		return "", errors.New("role not found")
	}

	if data.Name != nil && *data.Name != "" {
		checkRole.Name = *data.Name
	}
	if data.Description != nil && *data.Description != "" {
		checkRole.Description = *data.Description
	}

	role := &model.Role{
		Name:        checkRole.Name,
		Description: checkRole.Description,
	}

	updatedRoleId, err := r.repo.Update(ctx, id, *role)
	if err != nil {
		return "", errors.New("failed to update role")
	}
	return updatedRoleId, nil
}

func (r *RoleUsecaseImpl) Delete(ctx context.Context, id string) error {
	checkRole, err := r.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("role not found")
	}

	err = r.repo.Delete(ctx, checkRole.ID)
	if err != nil {
		return errors.New("failed to delete role")
	}

	return nil
}

func NewRoleUsecase(repo repository.RoleRepository) RoleUsecase {
	return &RoleUsecaseImpl{
		repo: repo,
	}
}
