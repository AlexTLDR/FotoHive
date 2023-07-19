package models

type Session struct {
	ID        int
	UserID    int
	TokenHash string
}
