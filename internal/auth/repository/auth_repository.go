package repository

type AuthRepository interface {
	SetUserSession(userID int, sessionToken string) error
	GetUserSession(userID int) (string, error)
	DeleteUserSession(userID int) error
}
