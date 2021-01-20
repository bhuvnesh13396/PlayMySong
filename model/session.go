package model

import (
	"errors"
	"time"
)

var (
	ErrSessionNotFound = errors.New("session not found")
)

type SessionRepo interface {
	Add(s Session) error
	Get(token string) (s Session, err error)
}

type Session struct {
	Token      string
	UserID     string
	ExpiryDate time.Time
}
