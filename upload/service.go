package upload

import (
	"context"

	"github.com/bhuvnesh13396/PlayMySong/model"
)

type Service interface {
	Add(ctx context.Context, songID string, songFile []byte) error
}

type service struct {
	data model.Uploader
}

func NewService(data model.Uploader) Service {
	return &service{
		data: data,
	}
}

func (s *service) Add(ctx context.Context, songID string, songFile []byte) error {

	// Create new directory at given location
	path := "/home/bhuvi/work/src/github.com/bhuvnesh13396/PlayMySong/song_data/raw"

	data := model.SongUpload{
		SongID:   songID,
		SongFile: songFile,
		Path:     path,
	}
	//fmt.Println("In service ::: Data ", data)
	// Copy song at the given location
	err := s.data.Upload(data)
	return err
}
