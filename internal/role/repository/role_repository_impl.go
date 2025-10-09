package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gilangnyan/point-of-sales/internal/role/model"
	"gilangnyan/point-of-sales/package/request"
)

const (
	FindAllQuery  = `SELECT id, name, description FROM roles`
	FindByIDQuery = `SELECT id, name, description FROM roles WHERE id = $1`
	CreateQuery   = `INSERT INTO roles (name, description) VALUES ($1, $2) RETURNING id`
	UpdateQuery   = `UPDATE roles SET name = $1, description = $2 WHERE id = $3`
	DeleteQuery   = `DELETE FROM roles WHERE id = $1`
)

type RoleRepositoryImpl struct {
	db *sql.DB
}

func (r *RoleRepositoryImpl) FindAll(ctx context.Context, params *request.PaginationParams) ([]*model.Role, int64, error) {
	whereClause := ""
	args := []interface{}{}
	argIndex := 1

	orderClause := fmt.Sprintf("ORDER BY %s %s", params.SortBy, params.SortDir)

	countQuery := `SELECT COUNT(*) FROM roles ` + whereClause
	var total int64
	err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	dataQuery := FindAllQuery + whereClause + " " + orderClause + fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, params.GetLimit(), params.GetOffset())

	rows, err := r.db.QueryContext(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	return ScanRoles(rows, total)
}

func (r *RoleRepositoryImpl) FindByID(ctx context.Context, id string) (*model.Role, error) {
	row := r.db.QueryRowContext(ctx, FindByIDQuery, id)

	return ScanRole(row)
}

func (r *RoleRepositoryImpl) Create(ctx context.Context, data model.Role) (string, error) {
	var id string
	err := r.db.QueryRowContext(ctx, CreateQuery, data.Name, data.Description).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *RoleRepositoryImpl) Update(ctx context.Context, id string, data model.Role) (string, error) {
	_, err := r.db.ExecContext(ctx, UpdateQuery, data.Name, data.Description, id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *RoleRepositoryImpl) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, DeleteQuery, id)
	return err
}

func NewRoleRepository(db *sql.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db: db,
	}
}
