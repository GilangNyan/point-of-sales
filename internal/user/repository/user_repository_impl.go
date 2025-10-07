package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gilangnyan/point-of-sales/internal/user/model"
)

const (
	FindAllQuery        = `SELECT u.id, u.username, u.email, up.full_name, up.date_of_birth, up.phone_number, up.address, up.profile_picture, u.is_active, u.is_blocked FROM users u LEFT JOIN user_profiles up ON u.id = up.user_id`
	FindByIDQuery       = `SELECT u.id, u.username, u.email, up.full_name, up.date_of_birth, up.phone_number, up.address, up.profile_picture, u.is_active, u.is_blocked FROM users u LEFT JOIN user_profiles up ON u.id = up.user_id WHERE u.id = $1`
	FindByEmailQuery    = `SELECT u.id, u.username, u.email, up.full_name, up.date_of_birth, up.phone_number, up.address, up.profile_picture, u.is_active, u.is_blocked FROM users u LEFT JOIN user_profiles up ON u.id = up.user_id WHERE u.email = $1`
	FindByUsernameQuery = `SELECT u.id, u.username, u.email, up.full_name, up.date_of_birth, up.phone_number, up.address, up.profile_picture, u.is_active, u.is_blocked FROM users u LEFT JOIN user_profiles up ON u.id = up.user_id WHERE u.username = $1`
	CreateQuery         = `INSERT INTO users (username, email, password, is_active, is_blocked) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	UpdateQuery         = `UPDATE users SET username = $1, email = $2, is_active = $3, is_blocked = $4 WHERE id = $5`
	DeleteQuery         = `DELETE FROM users WHERE id = $1`
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func (r *UserRepositoryImpl) FindAll(ctx context.Context) ([]*model.UserWithProfile, error) {
	rows, err := r.db.QueryContext(ctx, FindAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanUsers(rows)
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, id string) (*model.UserWithProfile, error) {
	row := r.db.QueryRowContext(ctx, FindByIDQuery, id)

	return ScanUser(row)
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*model.UserWithProfile, error) {
	row := r.db.QueryRowContext(ctx, FindByEmailQuery, email)
	fmt.Printf("Result: %v\n", row)

	return ScanUser(row)
}

func (r *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (*model.UserWithProfile, error) {
	row := r.db.QueryRowContext(ctx, FindByUsernameQuery, username)
	fmt.Printf("Result: %v\n", row)

	return ScanUser(row)
}

func (r *UserRepositoryImpl) Create(ctx context.Context, data model.User) (string, error) {
	var id string
	err := r.db.QueryRowContext(ctx, CreateQuery, data.Username, data.Email, data.Password, data.IsActive, data.IsBlocked).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, id string, data model.User) (string, error) {
	_, err := r.db.ExecContext(ctx, UpdateQuery, data.Username, data.Email, data.IsActive, data.IsBlocked, id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, DeleteQuery, id)
	return err
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}
