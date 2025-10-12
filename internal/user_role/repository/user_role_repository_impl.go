package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gilangnyan/point-of-sales/internal/user_role/model"
)

const (
	AssignRoleQuery         = `INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2) ON CONFLICT (user_id, role_id) DO NOTHING`
	RemoveRoleQuery         = `DELETE FROM user_roles WHERE user_id = $1 AND role_id = $2`
	RemoveAllUserRolesQuery = `DELETE FROM user_roles WHERE user_id = $1`
	GetUserRolesQuery       = `SELECT r.name FROM user_roles ur JOIN roles r ON ur.role_id = r.id WHERE ur.user_id = $1`
	GetUserByUsernameQuery  = `SELECT id, email, username, password FROM users WHERE username = $1`
	HasRoleQuery            = `SELECT EXISTS(SELECT 1 FROM user_roles WHERE user_id = $1 AND role_id = $2)`
)

type UserRoleRepositoryImpl struct {
	db *sql.DB
}

func (u *UserRoleRepositoryImpl) AssignRoles(ctx context.Context, userID string, roleIDs []string) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = u.AssignRolesWithTx(ctx, tx, userID, roleIDs)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (u *UserRoleRepositoryImpl) AssignRolesWithTx(ctx context.Context, tx *sql.Tx, userID string, roleIDs []string) error {
	for _, roleID := range roleIDs {
		_, err := tx.ExecContext(ctx, AssignRoleQuery, userID, roleID)
		if err != nil {
			return fmt.Errorf("failed to assign role %s to user %s: %w", roleID, userID, err)
		}
	}
	return nil
}

func (u *UserRoleRepositoryImpl) RemoveRoles(ctx context.Context, userID string, roleIDs []string) error {
	for _, roleID := range roleIDs {
		_, err := u.db.ExecContext(ctx, RemoveRoleQuery, userID, roleID)
		if err != nil {
			return fmt.Errorf("failed to remove role %s from user %s: %w", roleID, userID, err)
		}
	}
	return nil
}

func (u *UserRoleRepositoryImpl) RemoveAllUserRoles(ctx context.Context, userID string) error {
	_, err := u.db.ExecContext(ctx, RemoveAllUserRolesQuery, userID)
	return err
}

func (u *UserRoleRepositoryImpl) RemoveAllUserRolesWithTx(ctx context.Context, tx *sql.Tx, userID string) error {
	_, err := tx.ExecContext(ctx, RemoveAllUserRolesQuery, userID)
	return err
}

func (u *UserRoleRepositoryImpl) GetUserRoles(ctx context.Context, userID string) ([]string, error) {
	rows, err := u.db.QueryContext(ctx, GetUserRolesQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []string
	for rows.Next() {
		var roleName string
		if err := rows.Scan(&roleName); err != nil {
			return nil, err
		}
		roles = append(roles, roleName)
	}

	return roles, nil
}

func (u *UserRoleRepositoryImpl) FindUserByUsernameWithRoles(ctx context.Context, username string) (model.UserWithRoles, error) {
	var user model.UserWithRoles
	err := u.db.QueryRowContext(ctx, GetUserByUsernameQuery, username).Scan(&user.UserID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return user, err
	}

	user.Roles, err = u.GetUserRoles(ctx, user.UserID)
	return user, err
}

func (u *UserRoleRepositoryImpl) HasRole(ctx context.Context, userID string, roleID string) (bool, error) {
	var exists bool
	err := u.db.QueryRowContext(ctx, HasRoleQuery, userID, roleID).Scan(&exists)
	return exists, err
}

func NewUserRoleRepository(db *sql.DB) UserRoleRepository {
	return &UserRoleRepositoryImpl{
		db: db,
	}
}
