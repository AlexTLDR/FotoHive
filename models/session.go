package models

import "database/sql"

type Session struct {
	ID     int
	UserID int
	// Token is only set when creating a new session
	// When looking up a session, this will be left empty, as only the hash of a session is stored
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	// TODO: Create the session token
	// TODO: Implement SessionService.Create
	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionService.User
	return nil, nil
}
