package song

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

func (s *logSvc) Get(ctx context.Context, songName string) (song SongResp, err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Get",
			"songName", songName,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Get(ctx, songName)
}

func (s *logSvc) Get1(ctx context.Context, ID string) (song SongResp, err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Get",
			"ID", ID,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Get(ctx, ID)
}

func (s *logSvc) Add(ctx context.Context, title string, length time.Time, artistID string, composerID string, lyrics string, path string, img string) (songID string, err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Add",
			"title", title,
			"length", length,
			"artistID", artistID,
			"composerID", composerID,
			"lyrics", lyrics,
			"path", path,
			"image", img,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Add(ctx, title, length, artistID, composerID, lyrics, path, img)
}

func (s *logSvc) Update(ctx context.Context, id string, title string) (err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Update",
			"id", id,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Update(ctx, id, title)
}

func (s *logSvc) List(ctx context.Context) (res []SongResp, err error) {
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
