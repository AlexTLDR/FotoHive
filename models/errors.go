package models

import "errors"

var (
	ErrEmailTaken = errors.New("models: email already taken")
	ErrNotFound   = errors.New("models: resource not found")
)
