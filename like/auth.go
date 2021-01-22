package like

import (
	"context"

	"github.com/bhuvnesh13396/PlayMySong/auth"
	"github.com/bhuvnesh13396/PlayMySong/common/auth/token"
	"github.com/bhuvnesh13396/PlayMySong/common/err"
)

var (
	errAccessTokenNotFound = err.New(1, "access token not found")
)

type authSvc struct {
	Service
	authService auth.Service
}

func NewAuthService(s Service, authService auth.Service) Service {
	return &authSvc{
		Service:     s,
		authService: authService,
	}
}

func (s *authSvc) verifyToken(ctx context.Context) (userID string, err error) {
	token, ok := ctx.Value(token.ContextKey).(string)
	if !ok || len(token) < 1 {
		return "", errAccessTokenNotFound
	}

	return s.authService.VerifyToken(ctx, token)
}

func (s *authSvc) Add(ctx context.Context, activityType string, activityID string, userID string) (err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.Add(ctx, activityType, activityID, userID)
}

func (s *authSvc) Delete(ctx context.Context, activityID string, userID string) (err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.Delete(ctx, activityID, userID)
}

func (s *authSvc) Get(ctx context.Context, activityID string) (count int, err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.Get(ctx, activityID)
}

func (s *authSvc) UpdateCount(ctx context.Context, activityID string, count int) (err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

  return s.Service.UpdateCount(ctx, activityID, count)
}
