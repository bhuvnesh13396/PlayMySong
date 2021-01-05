package song

import (
	"context"
	"time"

	"github.com/bhuvnesh13396/PlayMySong/common/err"
	"github.com/bhuvnesh13396/PlayMySong/common/id"
	"github.com/bhuvnesh13396/PlayMySong/model"
)

var (
	errInvalidArgument = err.New(101, "Invalid Arguments.")
)

type Service interface {
	Get(ctx context.Context, songName string) (res SongResp, err error)
	Get1(ctx context.Context, id string) (song model.Song, err error)
	Add(ctx context.Context, title string, length time.Time, artistID string, composerID string, lyrics string, path string, img string) (songID string, err error)
	List(ctx context.Context) (res []SongResp, err error)
	Update(ctx context.Context, id string, title string) (err error)
}

type service struct {
	songRepo model.SongRepo
}

func NewService(songRepo model.SongRepo) Service {
	return &service{
		songRepo: songRepo,
	}
}

func (s *service) Get(ctx context.Context, songName string) (song SongResp, err error) {
	if len(songName) < 1 {
		err = errInvalidArgument
		return
	}

	tempSong, err := s.songRepo.Get(songName)
	if err != nil {
		return
	}

	song = SongResp{
		ID:         tempSong.ID,
		Title:      tempSong.Title,
		Length:     tempSong.Length,
		ArtistID:   tempSong.ArtistID,
		ComposerID: tempSong.ComposerID,
		Lyrics:     tempSong.Lyrics,
		Path:       tempSong.Path,
		Image:      tempSong.Image,
	}

	return
}

func (s *service) Get1(ctx context.Context, Id string) (song SongResp, err error) {
	return nil, nil
}

func (s *service) Update(ctx context.Context, id string, title string) (err error) {
	return nil, nil
}

func (s *service) List(ctx context.Context) (res []SongResp, err error) {
	songs, err := s.songRepo.GetAll()
	if err != nil {
		return
	}

	for i := range songs {
		s := &songs[i]

		song := SongResp{
			ID:         s.ID,
			Title:      s.Title,
			Length:     s.Length,
			ArtistID:   s.ArtistID,
			ComposerID: s.ComposerID,
			Lyrics:     s.Lyrics,
			Path:       s.Path,
			Image:      s.Image,
		}

		res = append(res, song)
	}

	return
}

func (s *service) Add(ctx context.Context, title string, length time.Time, artistID string, composerID string, lyrics string, path string, img string) (songID string, err error) {
	song := model.Song{
		ID:          id.New(),
		UserID:      userID,
		Title:       title,
		Description: description,
	}

	err = s.songRepo.Add(song)
	if err != nil {
		return
	}

	return song.ID, nil
}
