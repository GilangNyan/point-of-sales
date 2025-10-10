package repository

import (
	"context"
	"database/sql"
	"gilangnyan/point-of-sales/internal/user/model"
)

type UserProfileRepository interface {
	Create(ctx context.Context, tx *sql.Tx, data model.UserProfile) (string, error)
	Update(ctx context.Context, tx *sql.Tx, userID string, data model.UserProfile) (string, error)
	Delete(ctx context.Context, tx *sql.Tx, userID string) error
}
