package repository

import (
	"context"
	"database/sql"
)

type UserRoleRepository interface {
	AssignRoles(ctx context.Context, userID string, roleIDs []string) error
	AssignRolesWithTx(ctx context.Context, tx *sql.Tx, userID string, roleIDs []string) error
	RemoveRoles(ctx context.Context, userID string, roleIDs []string) error
	RemoveAllUserRoles(ctx context.Context, userID string) error
	RemoveAllUserRolesWithTx(ctx context.Context, tx *sql.Tx, userID string) error
	GetUserRoles(ctx context.Context, userID string) ([]string, error)
	HasRole(ctx context.Context, userID string, roleID string) (bool, error)
}
