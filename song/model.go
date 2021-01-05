package song

import "time"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type SongResp struct {
	Title    string    `json:"title"`
	Length   time.Time `json:"length"`
	Artist   User      `json:"artist"`
	Composer User      `json:"composer"`
	Lyrics   string    `json:"lyrics"`
	Path     string    `json:"path"`
	Image    string    `json:img`
}
