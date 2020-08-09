package song

import (
	"context"
	"time"
)

var (
	errInvalidArgument = err.New(101, "Invalid Arguments.")
)

type Service interface {
	Get(ctx context.Context, ID string) (res GetResp, err error)
	Add(ctx context.Context, title string, length time.Time, artistID string, composerID string, lyrics string, path string, img string)
	List(ctx context.Context) (res []GetResp, err error)
}

type service struct {
	songRepo model.SongRepo
}

type NewService(songRepo model.SongRepo) Service {
	return &service{
		songRepo:	songRepo,
	}
}


func (s *service) Get(ctx context.Context, ID string) (song GetResp, err error){
	if len(ID) < 1 {
		err = errInvalidArgument
		return
	}

	tempSong, err := s.songRepo.Get(ID)
	if err != nil {
		return
	}

	song = GetResp {
		ID:	tempSong.ID,
		Title: tempSong.Title,
		Length: tempSong.Length,
		ArtistID: tempSong.ArtistID,
		ComposerID: tempSong.ComposerID,
		Lyrics: tempSong.Lyrics,
		Path: tempSong.Path,
		Image: tempSong.Image,
	}

	return
}

func (s *service) List(ctx context.Context) (res [] GetResp, err error){
	songs, err := s.songRepo.GetAll()
	if err != nil{
		return
	}

	for i := range songs {
		s := &songs[i]

		song := GetResp{
			ID: s.ID,
			Title: s.Title,
			Length: s.Length,
			ArtistID: s.ArtistID,
			ComposerID: s.ComposerID,
			Lyrics: s.Lyrics,
			Path: s.Path,
			Image: s.Image,
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