package repository

import "database/sql"

type AuthRepositoryImpl struct {
	db *sql.DB
}

func (a *AuthRepositoryImpl) SetUserSession(userID int, sessionToken string) error {
	panic("unimplemented")
}

func (a *AuthRepositoryImpl) GetUserSession(userID int) (string, error) {
	panic("unimplemented")
}

func (a *AuthRepositoryImpl) DeleteUserSession(userID int) error {
	panic("unimplemented")
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &AuthRepositoryImpl{
		db: db,
	}
}
