package repository

import (
	"context"
	"database/sql"
	"gilangnyan/point-of-sales/internal/user/model"
)

type UserProfileRepositoryImpl struct {
	db *sql.DB
}

const (
	CreateUserProfileQuery = `INSERT INTO user_profiles (full_name, user_id, phone_number, date_of_birth, address) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	UpdateUserProfileQuery = `UPDATE user_profiles SET full_name=$1, phone_number=$2, date_of_birth=$3, address=$4 WHERE user_id=$5 RETURNING id`
	DeleteUserProfileQuery = `DELETE FROM user_profiles WHERE user_id=$1`
)

func (r *UserProfileRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, data model.UserProfile) (string, error) {
	var id string
	err := tx.QueryRowContext(ctx, CreateUserProfileQuery, data.FullName, data.UserID, data.PhoneNumber, data.DateOfBirth, data.Address).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *UserProfileRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, userID string, data model.UserProfile) (string, error) {
	var id string
	// err := r.db.QueryRowContext(ctx, UpdateUserProfileQuery, data.FullName, data.PhoneNumber, data.DateOfBirth, data.Address, userID).Scan(&id)
	err := tx.QueryRowContext(ctx, UpdateUserProfileQuery, data.FullName, data.PhoneNumber, data.DateOfBirth, data.Address, userID).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *UserProfileRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userID string) error {
	// _, err := r.db.ExecContext(ctx, DeleteUserProfileQuery, userID)
	_, err := tx.ExecContext(ctx, DeleteUserProfileQuery, userID)
	return err
}

func NewUserProfileRepository(db *sql.DB) UserProfileRepository {
	return &UserProfileRepositoryImpl{
		db: db,
	}
}
