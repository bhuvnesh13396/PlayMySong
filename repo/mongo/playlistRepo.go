package mongo

import (
	"github.com/bhuvnesh13396/PlayMySong/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type playlistRepo struct {
	db *mongo.Client
}

func NewPlaylistRepo(db *mongo.Client) (*playlistRepo, err) {
	return &playlistRepo{
		db: db,
	}, nil
}

func (repo *playlistRepo) Add(p model.Playlist) (err error) {

}

func (repo *playlistRepo) Get(ID string) (p model.Playlist, err error) {

}

func (repo *playlistRepo) Update(ID string, songIDs []string) (err error) {

}

func (repo *playlistRepo) List() (playlists []mode.Playlist, err error) {

}
