package song

import (
	"context"
	"time"

	"github.com/bhuvnesh13396/PlayMySong/account"
	"github.com/bhuvnesh13396/PlayMySong/common/err"
	"github.com/bhuvnesh13396/PlayMySong/common/id"
	"github.com/bhuvnesh13396/PlayMySong/model"
)

var (
	errInvalidArgument = err.New(101, "Invalid Arguments.")
	errNoComposerFound = err.New(102, "No Composer Found")
	errNoArtistFound   = err.New(103, "No Artist Found")
)

type Service interface {
	Get(ctx context.Context, songName string) (res SongResp, err error)
	Get1(ctx context.Context, id string) (song SongResp, err error)
	Add(ctx context.Context, title string, length time.Time, artistID string, composerID string, lyrics string, path string, img string) (songID string, err error)
	List(ctx context.Context) (res []SongResp, err error)
	Update(ctx context.Context, id string, title string) (err error)
}

type service struct {
	songRepo   model.SongRepo
	accountSvc account.Service
}

func NewService(songRepo model.SongRepo, accountSvc account.Service) Service {
	return &service{
		songRepo:   songRepo,
		accountSvc: accountSvc,
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

	composer, err := s.accountSvc.Get1(ctx, tempSong.ComposerID)
	if err != nil {
		err = errNoComposerFound
		return
	}

	artist, err := s.accountSvc.Get1(ctx, tempSong.ArtistID)
	if err != nil {
		err = errNoArtistFound
		return
	}

	song = SongResp{
		Title:  tempSong.Title,
		Length: tempSong.Length,

		Artist: User{
			ID:       artist.ID,
			Name:     artist.Name,
			Username: artist.Username,
		},

		Composer: User{
			ID:       composer.ID,
			Name:     composer.Name,
			Username: composer.Username,
		},

		Lyrics: tempSong.Lyrics,
		Path:   tempSong.Path,
		Image:  tempSong.Image,
	}

	return song, nil
}

func (s *service) Get1(ctx context.Context, ID string) (song SongResp, err error) {

	tempSong, err := s.songRepo.Get1(ID)
	if err != nil {
		return
	}

	composer, err := s.accountSvc.Get1(ctx, tempSong.ComposerID)
	if err != nil {
		return SongResp{}, err
	}

	artist, err := s.accountSvc.Get1(ctx, tempSong.ArtistID)
	if err != nil {
		return SongResp{}, err
	}

	song = SongResp{
		Title:  tempSong.Title,
		Length: tempSong.Length,

		Artist: User{
			ID:       artist.ID,
			Name:     artist.Name,
			Username: artist.Username,
		},

		Composer: User{
			ID:       composer.ID,
			Name:     composer.Name,
			Username: composer.Username,
		},

		Lyrics: tempSong.Lyrics,
		Path:   tempSong.Path,
		Image:  tempSong.Image,
	}

	return song, nil
}

func (s *service) Update(ctx context.Context, ID string, title string) (err error) {
	err = s.songRepo.Update(ID, title)
	return
}

func (s *service) List(ctx context.Context) (res []SongResp, err error) {
	songs, err := s.songRepo.List()
	if err != nil {
		return
	}

	for i := range songs {
		song := &songs[i]

		artist, err := s.accountSvc.Get1(ctx, song.ArtistID)
		if err != nil {
			return []SongResp{}, err
		}
		composer, err := s.accountSvc.Get1(ctx, song.ComposerID)
		if err != nil {
			return []SongResp{}, err
		}

		createSong := SongResp{
			Title:  song.Title,
			Length: song.Length,

			Artist: User{
				ID:       artist.ID,
				Name:     artist.Name,
				Username: artist.Username,
			},

			Composer: User{
				ID:       composer.ID,
				Name:     composer.Name,
				Username: composer.Username,
			},

			Lyrics: song.Lyrics,
			Path:   song.Path,
			Image:  song.Image,
		}

		res = append(res, createSong)
	}

	return
}

func (s *service) Add(ctx context.Context, title string, length time.Time, artistID string, composerID string, lyrics string, path string, img string) (songID string, err error) {

	song := model.Song{
		ID:         id.New(),
		Title:      title,
		Length:     length,
		ArtistID:   artistID,
		ComposerID: composerID,
		Lyrics:     lyrics,
		Path:       path,
		Image:      img,
	}

	err = s.songRepo.Add(song)
	if err != nil {
		return
	}

	return song.ID, nil
}
