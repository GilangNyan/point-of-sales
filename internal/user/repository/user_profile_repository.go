package repository

import (
	"context"
	"gilangnyan/point-of-sales/internal/user/model"
)

type UserProfileRepository interface {
	Create(ctx context.Context, data model.UserProfile) (string, error)
	Update(ctx context.Context, userID string, data model.UserProfile) (string, error)
	Delete(ctx context.Context, userID string) error
}
