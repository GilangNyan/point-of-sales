package repository

import (
	"context"
	"database/sql"
	"gilangnyan/point-of-sales/internal/user/model"
)

const (
	FindAllQuery        = `SELECT id, username, email, password, is_active, is_blocked FROM users`
	FindByIDQuery       = `SELECT id, username, email, password, is_active, is_blocked FROM users WHERE id = $1`
	FindByEmailQuery    = `SELECT id, username, email, password, is_active, is_blocked FROM users WHERE email = $1`
	FindByUsernameQuery = `SELECT id, username, email, password, is_active, is_blocked FROM users WHERE username = $1`
	CreateQuery         = `INSERT INTO users (username, email, password, is_active, is_blocked) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	UpdateQuery         = `UPDATE users SET username = $1, email = $2, password = $3, is_active = $4, is_blocked = $5 WHERE id = $6`
	DeleteQuery         = `DELETE FROM users WHERE id = $1`
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func (r *UserRepositoryImpl) FindAll(ctx context.Context) ([]*model.User, error) {
	rows, err := r.db.QueryContext(ctx, FindAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanUsers(rows)
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, id string) (*model.User, error) {
	row := r.db.QueryRowContext(ctx, FindByIDQuery, id)

	return ScanUser(row)
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	row := r.db.QueryRowContext(ctx, FindByEmailQuery, email)

	return ScanUser(row)
}

func (r *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	row := r.db.QueryRowContext(ctx, FindByUsernameQuery, username)

	return ScanUser(row)
}

func (r *UserRepositoryImpl) Create(ctx context.Context, data model.User) (*model.User, error) {
	var id string
	err := r.db.QueryRowContext(ctx, CreateQuery, data.Username, data.Email, data.Password, data.IsActive, data.IsBlocked).Scan(&id)
	if err != nil {
		return nil, err
	}
	data.ID = id
	return &data, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, id string, data model.User) (*model.User, error) {
	_, err := r.db.ExecContext(ctx, UpdateQuery, data.Username, data.Email, data.Password, data.IsActive, data.IsBlocked, id)
	if err != nil {
		return nil, err
	}

	data.ID = id
	return &data, nil
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, DeleteQuery, id)
	return err
}

func NewUserRepository(sql *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: sql,
	}
}
