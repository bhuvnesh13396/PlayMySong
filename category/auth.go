package category

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

func (s *authSvc) Get(ctx context.Context, ID string) (category CategoryResp, err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.Get(ctx, ID)
}

func (s *authSvc) Add(ctx context.Context, title string, category_type string, songIDs []string) (categoryID string, err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.Add(ctx, title, category_type, songIDs)
}

func (s *authSvc) List(ctx context.Context) (categories []CategoryResp, err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.List(ctx)
}

func (s *authSvc) Update(ctx context.Context, ID string, songsIDs []string) (err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.Update(ctx, ID, songsIDs)
}
