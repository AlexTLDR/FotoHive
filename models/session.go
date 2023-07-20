package models

import (
	"database/sql"
	"fmt"

	"github.com/AlexTLDR/WebDev/rand"
)

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
	token, err := rand.SessionToken()
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	// TODO: Hash the token
	session := Session{
		UserID: userID,
		Token:  token,
		// TODO: Hash the token
	}
	// TODO: Store the session in our DB
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionService.User
	return nil, nil
}
