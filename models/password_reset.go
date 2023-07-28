package models

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	DefaultResetDuration = 1 * time.Hour
)

type PasswordReset struct {
	ID     int
	UserID int
	// Token is only set when a PasswordReset is created
	Token     string
	TokenHash string
	ExpiresAt time.Time
}

type PasswordResetService struct {
	DB *sql.DB
	// MinBytesPerToken is the minimum number of bytes to be used when generating each password reset token
	// If this value is not set or less than the MinBytesPerToken, the MinBytesPerToken will be used
	BytesPerToken int
	// Duration is the amount of time a password reset token is valid for
	// If this value is not set, it defaults to DefaultResetDuration
	Duration time.Duration
}

func (service *PasswordResetService) Create(email string) (*PasswordReset, error) {
	return nil, fmt.Errorf("TODO: create password reset")
}

func (service *PasswordResetService) Consume(token string) (*User, error) {
	return nil, fmt.Errorf("TODO: consume password reset")
}
