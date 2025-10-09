package repository

import (
	"database/sql"
	"gilangnyan/point-of-sales/internal/role/model"
)

func ScanRole(row *sql.Row) (*model.Role, error) {
	var role model.Role
	if err := row.Scan(&role.ID, &role.Name, &role.Description); err != nil {
		return nil, err
	}
	return &role, nil
}

func ScanRoles(rows *sql.Rows, total int64) ([]*model.Role, int64, error) {
	var roles []*model.Role
	for rows.Next() {
		var role model.Role
		if err := rows.Scan(&role.ID, &role.Name, &role.Description); err != nil {
			return nil, 0, err
		}
		roles = append(roles, &role)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}
	return roles, total, nil
}
