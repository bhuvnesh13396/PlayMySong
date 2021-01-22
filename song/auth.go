package song

import (
	"context"
	"time"

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

func (s *authSvc) Get(ctx context.Context, songName string) (res SongResp, err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.Get(ctx, songName)
}

func (s *authSvc) Get1(ctx context.Context, id string) (song SongResp, err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.Get1(ctx, id)
}

func (s *authSvc) Add(ctx context.Context, title string, length time.Time, artistID string, composerID string, lyrics string, path string, img string) (songID string, err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.Add(ctx, title, length, artistID, composerID, lyrics, path, img)
}

func (s *authSvc) List(ctx context.Context) (res []SongResp, err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.List(ctx)
}

func (s *authSvc) Update(ctx context.Context, id string, title string) (err error) {
	_, err = s.verifyToken(ctx)
	if err != nil {
		return
	}

	return s.Service.Update(ctx, id, title)
}
