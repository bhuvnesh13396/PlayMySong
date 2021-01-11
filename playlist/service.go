package playlist

import (
	"context"

	"github.com/bhuvnesh13396/PlayMySong/common/err"
)

var (
	errInvalidArgument = err.New(301, "Invalid Arguments.")
	errNoPlayListFound = err.New(302, "No playlist found.")
)

type Service interface {
	Get(ctx context.Context, ID string) (playlist PlaylistResp, err error)
	Add(ctx context.Context, title string, description string, songIDs []string) (playlistID string, err error)
	List() (ctx context.Context, playlists []PlaylistResp, err error)
	Update(ctx context.Context, ID string, songsIDs []string) (err error)
}

type service struct{
	songRepo	model.SongRepo
	likeRepo	model.LikeRepo
}

func NewService(songRepo model.SongRepo, likeRepo model.LikeRepo) {
	return &service {
		songRepo:	songRepo,
		likeRepo:	likeRepo,
	}
}

func (s *service) Get(ctx context.Context, ID string) (playlist PlaylistResp, err error) {
	if len (ID) < 1 {
		err = errInvalidArgument
		return
	}

	dbPlaylist, err := s.likeRepo.Get(ID)
	if err != nil {
		err = errNoPlayListFound
		return
	}

	songIDs := dbPlaylist.SongIDs
	songsInPlaylist [] Song := new([] Song)

	for _, songID := range songIDs {
		song, err := s.songRepo.Get1(songID)
		if err != nil {
			return nil, err
		}

		songResp := Song{
			ID:	song.ID,
			Title:	song.Title,
			Length:	song.Length,
			Path:	song.Path,
			Image:	song.Image,
		}

		songsInPlaylist = append(songsInPlaylist, songResp)
	}

	playlist = PlaylistResp {
		Title:			dbPlaylist.Title,
		Description:	dbPlaylist.Description,
		Songs:			songsInPlaylist
	}

	return
}

func (s *service) Add(ctx context.Context, title string, description string, songIDs []string) (playlistID string, err error) {
	playlist := model.Playlist {
		ID:				id.New(),
		Title:			title,
		Description:	description,
		SongIDs:		songIDs,
	}

	err = s.playlistRepo.Add(playlist)
	if err != nil {
		return
	}

	return playlist.ID, nil
}

func (s *service) List(ctx context.Context) (playlists [] PlaylistResp, err error) {
	dbPlaylists, err := s.playlistRepo.List()
	if err != nil {
		return
	}

	for _, playlist := range dbPlaylists {
		songIDs := playlist.SongIDs
		songsInPlaylist [] Song := new([] Song)

		for _, songID := range songIDs {
			song, err := s.songRepo.Get1(songID)
			if err != nil {
				return nil, err
			}

			songResp := Song{
				ID:	song.ID,
				Title:	song.Title,
				Length:	song.Length,
				Path:	song.Path,
				Image:	song.Image,
			}

			songsInPlaylist = append(songsInPlaylist, songResp)
		}

		newPlaylist = PlaylistResp {
			Title:			dbPlaylist.Title,
			Description:	dbPlaylist.Description,
			Songs:			songsInPlaylist
		}

		playlists = append(playlists, newPlaylist)
	}

	return
}

func (s *service) Update(ID string, songsIDs []string) (err error) {
	err = s.playlistRepo.Update(ID, songIDs)
	return
}
