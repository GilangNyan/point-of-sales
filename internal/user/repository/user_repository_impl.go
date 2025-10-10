package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gilangnyan/point-of-sales/internal/user/model"
	"gilangnyan/point-of-sales/package/request"
)

const (
	FindAllQuery        = `SELECT u.id, u.username, u.email, up.full_name, up.date_of_birth, up.phone_number, up.address, up.profile_picture, u.is_active, u.is_blocked FROM users u LEFT JOIN user_profiles up ON u.id = up.user_id`
	FindByIDQuery       = `SELECT u.id, u.username, u.email, up.full_name, up.date_of_birth, up.phone_number, up.address, up.profile_picture, u.is_active, u.is_blocked FROM users u LEFT JOIN user_profiles up ON u.id = up.user_id WHERE u.id = $1`
	FindByEmailQuery    = `SELECT u.id, u.username, u.email, up.full_name, up.date_of_birth, up.phone_number, up.address, up.profile_picture, u.is_active, u.is_blocked FROM users u LEFT JOIN user_profiles up ON u.id = up.user_id WHERE u.email = $1`
	FindByUsernameQuery = `SELECT u.id, u.username, u.email, up.full_name, up.date_of_birth, up.phone_number, up.address, up.profile_picture, u.is_active, u.is_blocked FROM users u LEFT JOIN user_profiles up ON u.id = up.user_id WHERE u.username = $1`
	CreateQuery         = `INSERT INTO users (username, email, password, is_active, is_blocked) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	UpdateQuery         = `UPDATE users SET username = $1, email = $2, is_active = $3, is_blocked = $4 WHERE id = $5 RETURNING id`
	DeleteQuery         = `DELETE FROM users WHERE id = $1`
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func (r *UserRepositoryImpl) FindAll(ctx context.Context, params *request.PaginationParams) ([]*model.UserWithProfile, int64, error) {
	whereClause := ""
	args := []interface{}{}
	argIndex := 1

	orderClause := fmt.Sprintf("ORDER BY u.%s %s", params.SortBy, params.SortDir)

	countQuery := `SELECT COUNT(DISTINCT u.id) FROM users u LEFT JOIN user_profiles up ON u.id = up.user_id ` + whereClause
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

	return ScanUsers(rows, total)
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

func (r *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, data model.User) (string, error) {
	var id string
	// err := r.db.QueryRowContext(ctx, CreateQuery, data.Username, data.Email, data.Password, data.IsActive, data.IsBlocked).Scan(&id)
	err := tx.QueryRowContext(ctx, CreateQuery, data.Username, data.Email, data.Password, data.IsActive, data.IsBlocked).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, id string, data model.User) (string, error) {
	var userID string
	// _, err := r.db.ExecContext(ctx, UpdateQuery, data.Username, data.Email, data.IsActive, data.IsBlocked, id)
	err := tx.QueryRowContext(ctx, UpdateQuery, data.Username, data.Email, data.IsActive, data.IsBlocked, id).Scan(&userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id string) error {
	// _, err := r.db.ExecContext(ctx, DeleteQuery, id)
	_, err := tx.ExecContext(ctx, DeleteQuery, id)
	return err
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}
