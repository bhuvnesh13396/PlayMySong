package model

import (
	"errors"
	"time"
)

var (
	SongNotFound = errors.New("Song not found")
)

type Song struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Length   time.Time `json:"length`
	Artist   string    `json:"artist_id"`
	Composer string    `json:"composer_id"`
	Lyrics   string    `json:"lyrics"`
	Path     string    `json:"path"`
	Image    string    `json:"img"`
}

type SongRepo interface {
	Add(song Song) (err error)
	Get(ID string) (song Song, err error)
	List() (song []Song, err error)
	Update(song Song, Title string) (err error)
}
