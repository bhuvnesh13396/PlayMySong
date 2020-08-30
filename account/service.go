package account

import (
	"context"
	"errors"

	"github.com/bhuvnesh13396/PlayMySong/model"

	"github.com/bhuvnesh13396/PlayMySong/common/id"

	"golang.org/x/crypto/bcrypt"
)

var (
	errInvalidArgument = errors.New("invalid argument")
)

type Service interface {
	Get(ctx context.Context, username string) (account model.Account, err error)
	Get1(ctx context.Context, id string) (account model.Account, err error)
	Add(ctx context.Context, name string, username string, password string) (err error)
	Update(ctx context.Context, username string, name string) (err error)
	List(ctx context.Context) (account []model.Account, err error)
}

type service struct {
	accountRepo model.AccountRepo
}

func NewService(accountRepo model.AccountRepo) Service {
	return &service{
		accountRepo: accountRepo,
	}
}

func (s *service) Get1(ctx context.Context, username string) (account model.Account, err error) {
	if len(username) < 1 {
		err = errInvalidArgument
		return
	}
	return s.accountRepo.Get1(username)
}

func (s *service) Get(ctx context.Context, id string) (account model.Account, err error) {
	if len(id) < 1 {
		err = errInvalidArgument
		return
	}
	return s.accountRepo.Get(id)
}

func (s *service) Add(ctx context.Context, name string, username string, password string) (err error) {
	// TODO return proper error messages
	if len(name) < 1 || len(username) < 4 || len(password) < 8 {
		return errInvalidArgument
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return
	}

	acc := model.Account{
		ID:       id.New(),
		Name:     name,
		Username: username,
		Password: string(hash),
	}

	return s.accountRepo.Add(acc)
}

func (s *service) Update(ctx context.Context, username string, name string) (err error) {
	if len(username) < 1 || len(name) < 1 {
		err = errInvalidArgument
		return
	}

	return s.accountRepo.Update(username, name)

}

func (s *service) List(ctx context.Context) (accounts []model.Account, err error) {
	return s.accountRepo.GetAll()
}
