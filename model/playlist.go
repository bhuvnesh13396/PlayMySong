package model

import (
	"errors"
)

var (
	PlaylistNotFound = errors.New("Playlist not found")
)

type Playlist struct {
	ID      		string   `json:"id"`
	Title   		string   `json:"title"`
	Description	string	`json:"descritption"`
	Access  		string   `json:"access"` // "Public/Private/Shared/Featured"
	SongIDs 		[]string `json:"songs_ids"`
}

type PlaylistRepo interface {
	Add(playlist Playlist) (err error)
	Get(ID string) (playlist Playlist, err error)
	List() (playlists []Playlist, err error)
	Update(ID string, songIDs []string) (err error)
}
