package repository

import (
	"database/sql"
	"gilangnyan/point-of-sales/internal/user/model"
)

func ScanUser(row *sql.Row) (*model.UserWithProfile, error) {
	var user model.UserWithProfile
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.FullName, &user.DateOfBirth, &user.PhoneNumber, &user.Address, &user.ProfilePicture, &user.IsActive, &user.IsBlocked); err != nil {
		return nil, err
	}
	return &user, nil
}

func ScanUsers(rows *sql.Rows) ([]*model.UserWithProfile, error) {
	var users []*model.UserWithProfile
	for rows.Next() {
		var user model.UserWithProfile
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.FullName, &user.DateOfBirth, &user.PhoneNumber, &user.Address, &user.ProfilePicture, &user.IsActive, &user.IsBlocked); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
