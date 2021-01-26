package upload

import (
	"context"

	"github.com/bhuvnesh13396/PlayMySong/common/kit"
)

type AddEndpoint kit.Endpoint

type Endpoint struct {
	AddEndpoint
}

type addRequest struct {
	SongID   string
	SongFile []byte
}

type addResponse struct {
}

func MakeAddEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest)
		err := s.Add(ctx, req.SongID, req.SongFile)
		return addResponse{}, err
	}
}

func (e AddEndpoint) Add(ctx context.Context, song_id string, song_file []byte) (err error) {
	request := addRequest{
		SongID:   song_id,
		SongFile: song_file,
	}
	_, err = e(ctx, request)
	return err
}
