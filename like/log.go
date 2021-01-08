package like

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

func (s *logSvc) Get(ctx context.Context, activityID string) (count int, err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Get",
			"activity_id", activityID,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Get(ctx, activityID)
}

func (s *logSvc) Add(ctx context.Context, activityType string, activityID string, userID string) (err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Add",
			"activityType", activityType,
			"activity_id", activityID,
			"user_id", userID,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Add(ctx, activityType, activityID, userID)
}

func (s *logSvc) Delete(ctx context.Context, activityID string, userID string) (err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "Delete",
			"activity_id", activityID,
			"user_id", userID,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.Delete(ctx, activityID, userID)
}

func (s *logSvc) UpdateCount(ctx context.Context, activityID string, count int) (err error) {
	defer func(t time.Time) {
		s.logger.Log(
			"ts", t,
			"method", "UpdateCount",
			"activity_id", activityID,
			"count", count,
			"took", time.Since(t),
			"err", err,
		)
	}(time.Now())
	return s.Service.UpdateCount(ctx, activityID, count)
}
