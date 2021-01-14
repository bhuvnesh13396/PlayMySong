package category

import (
	"context"
	"time"

	"github.com/bhuvnesh13396/PlayMySong/common/kit"
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

func (s *logSvc) Get(ctx context.Context, ID string) (category CategoryResp, err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Get",
			"id", ID,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Get(ctx, ID)
}

func (s *logSvc) Add(ctx context.Context, title string, type string, songIDs []string) (categoryID string, err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Add",
			"title", title,
			"type", type,
			"songIDs", songIDs,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Add(ctx, title, type, songIDs)
}

func (s *logSvc) Update(ctx context.Context, ID string, songsIDs []string) (err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Update",
			"id", ID,
			"songIDs", songsIDs,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Update(ctx, ID, songsIDs)
}

func (s *logSvc) List(ctx context.Context) (categories []CategoryResp, err error) {
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
