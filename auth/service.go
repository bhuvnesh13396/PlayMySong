package auth

import (
	"context"
	"time"

	"github.com/bhuvnesh13396/PlayMySong/common/err"
	"github.com/bhuvnesh13396/PlayMySong/common/sid"
	"github.com/bhuvnesh13396/PlayMySong/model"

	"golang.org/x/crypto/bcrypt"
)

var (
	sessionDuration = 365 * 24 * time.Hour
)

var (
	errInvalidArgument = err.New(301, "invalid argument")
	errInvalidPassword = err.New(302, "invalid password")
	errInvalidToken    = err.New(303, "invalid token")
)

type Service interface {
	Signin(ctx context.Context, username string, password string) (token string, err error)
	VerifyToken(ctx context.Context, token string) (userID string, err error)
}

type service struct {
	sessionRepo model.SessionRepo
	accountRepo model.AccountRepo
}

func NewService(sessionRepo model.SessionRepo, accountRepo model.AccountRepo) Service {
	return &service{
		sessionRepo: sessionRepo,
		accountRepo: accountRepo,
	}
}

func (s *service) Signin(ctx context.Context, username string, password string) (token string, err error) {
	if len(username) < 1 || len(password) < 1 {
		err = errInvalidArgument
		return
	}

	u, err := s.accountRepo.Get1(username)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			err = errInvalidPassword
		}
		return
	}

	t, err := sid.New(24)
	if err != nil {
		return
	}

	session := model.Session{
		Token:      t,
		UserID:     u.ID,
		ExpiryDate: time.Now().Add(sessionDuration),
	}

	err = s.sessionRepo.Add(session)
	if err != nil {
		return
	}

	return t, nil
}

func (s *service) VerifyToken(ctx context.Context, token string) (userID string, err error) {
	if len(token) < 1 {
		err = errInvalidArgument
		return
	}

	session, err := s.sessionRepo.Get(token)
	if err != nil {
		return "", errInvalidToken
	}

	if session.ExpiryDate.Before(time.Now()) {
		return "", errInvalidToken
	}

	return session.UserID, nil
}
