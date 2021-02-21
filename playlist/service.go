package playlist

import (
	"context"

	"github.com/bhuvnesh13396/PlayMySong/common/err"
	"github.com/bhuvnesh13396/PlayMySong/common/id"
	"github.com/bhuvnesh13396/PlayMySong/model"
	"github.com/bhuvnesh13396/PlayMySong/song"
)

var (
	errInvalidArgument = err.New(301, "Invalid Arguments.")
	errNoPlayListFound = err.New(302, "No playlist found.")
)

type Service interface {
	Get(ctx context.Context, ID string) (playlist PlaylistResp, err error)
	Add(ctx context.Context, title string, description string, songIDs []string) (playlistID string, err error)
	List(ctx context.Context) (playlists []PlaylistResp, err error)
	Update(ctx context.Context, ID string, songsIDs []string) (err error)
}

type service struct {
	songSvc      song.Service
	playlistRepo model.PlaylistRepo
}

func NewService(songSvc song.Service, playlistRepo model.PlaylistRepo) Service {
	return &service{
		songSvc:      songSvc,
		playlistRepo: playlistRepo,
	}
}

func (s *service) Get(ctx context.Context, ID string) (playlist PlaylistResp, err error) {
	if len(ID) < 1 {
		err = errInvalidArgument
		return
	}

	dbPlaylist, err := s.playlistRepo.Get(ID)
	if err != nil {
		err = errNoPlayListFound
		return
	}

	songIDs := dbPlaylist.SongIDs
	songsInPlaylist := make([]Song, 0)

	for _, songID := range songIDs {
		song, err := s.songSvc.Get1(ctx, songID)
		if err != nil {
			return PlaylistResp{}, err
		}

		songResp := Song{
			ID:     songID,
			Title:  song.Title,
			Length: song.Length,
			Path:   song.Path,
			Image:  song.Image,
		}

		songsInPlaylist = append(songsInPlaylist, songResp)
	}

	playlist = PlaylistResp{
		Title:       dbPlaylist.Title,
		Description: dbPlaylist.Description,
		Songs:       songsInPlaylist,
	}

	return
}

func (s *service) Add(ctx context.Context, title string, description string, songIDs []string) (playlistID string, err error) {
	playlist := model.Playlist{
		ID:          id.New(),
		Title:       title,
		Description: description,
		SongIDs:     songIDs,
	}

	err = s.playlistRepo.Add(playlist)
	if err != nil {
		return
	}

	return playlist.ID, nil
}

func (s *service) List(ctx context.Context) (playlists []PlaylistResp, err error) {
	dbPlaylists, err := s.playlistRepo.List()
	if err != nil {
		return
	}

	for _, playlist := range dbPlaylists {
		songIDs := playlist.SongIDs
		songsInPlaylist := make([]Song, 0)

		for _, songID := range songIDs {
			song, err := s.songSvc.Get1(ctx, songID)
			if err != nil {
				return nil, err
			}

			songResp := Song{
				ID:     songID,
				Title:  song.Title,
				Length: song.Length,
				Path:   song.Path,
				Image:  song.Image,
			}

			songsInPlaylist = append(songsInPlaylist, songResp)
		}

		newPlaylist := PlaylistResp{
			Title:       playlist.Title,
			Description: playlist.Description,
			Songs:       songsInPlaylist,
		}

		playlists = append(playlists, newPlaylist)
	}

	return
}

func (s *service) Update(ctx context.Context, ID string, songIDs []string) (err error) {
	err = s.playlistRepo.Update(ID, songIDs)
	return
}
