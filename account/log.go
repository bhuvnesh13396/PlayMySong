package account

import (
	"context"
	"time"

	"github.com/bhuvnesh13396/PlayMySong/common/kit"
	"github.com/bhuvnesh13396/PlayMySong/model"
)

type logSvc struct {
	Service
	logger kit.Logger
}

func NewLogService(s Service, logger kit.Logger) Service {
	return &logSvc{
		Service: s,
		logger:  logger,
	}
}

func (s *logSvc) Get(ctx context.Context, username string) (account model.Account, err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Get",
			"username", username,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Get(ctx, username)
}

func (s *logSvc) Get1(ctx context.Context, id string) (account model.Account, err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Get1",
			"id", id,
			"time", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Get1(ctx, id)
}

func (s *logSvc) Add(ctx context.Context, name string, username string, password string) (err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Add",
			"name", name,
			"username", username,
			"password", password,
			"time", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Add(ctx, name, username, password)
}

func (s *logSvc) Update(ctx context.Context, username string, name string) (err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Update",
			"username", username,
			"name", name,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Update(ctx, username, name)
}
func (s *logSvc) List(ctx context.Context) (account []model.Account, err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "List",
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.List(ctx)
}
