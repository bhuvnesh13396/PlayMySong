package song

import "time"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type SongResp struct {
	Title      string    `json:"title"`
	Length     time.Time `json:"length"`
	ArtistID   User      `json:"artist"`
	ComposerID User      `json:"composer"`
	Lyrics     string    `json:"lyrics"`
	Path       string    `json:"path"`
	Img        string    `json:img`
}
