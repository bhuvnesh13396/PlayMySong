package model

import (
	"errors"
	"time"
)

var (
	SongNotFound = errors.New("Song not found")
)

type Song struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Length     time.Time `json:"length`
	ArtistID   string    `json:"artist_id"`
	ComposerID string    `json:"composer_id"`
	Lyrics     string    `json:"lyrics"`
	Path       string    `json:"path"`
	Image      string    `json:"img"`
}

type SongRepo interface {
	Add(song Song) (err error)
	Get(songName string) (song Song, err error)
	Get1(ID string) (song Song, err error)
	List() (song []Song, err error)
	Update(ID string, Title string) (err error)
}
