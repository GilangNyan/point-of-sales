package repository

import (
	"database/sql"
	"gilangnyan/point-of-sales/internal/user/model"
)

func ScanUser(row *sql.Row) (*model.User, error) {
	var user model.User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.IsActive, &user.IsBlocked); err != nil {
		return nil, err
	}
	return &user, nil
}

func ScanUsers(rows *sql.Rows) ([]*model.User, error) {
	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.IsActive, &user.IsBlocked); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
