package models

import (
	"database/sql"
	"fmt"

	"github.com/AlexTLDR/WebDev/rand"
)

const (
	// The minimum number of bytes to be used for each session token
	MinBytesPerToken = 32
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
	// MinBytesPerToken is the minimum number of bytes to be used when generating a new session token
	// If this value is not set or less than the MinBytesPerToken, the MinBytesPerToken will be used
	BytesPerToken int
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(ss.BytesPerToken)
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
