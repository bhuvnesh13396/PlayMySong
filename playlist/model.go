package playlist

import "time"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type Song struct {
	ID     string    `json:"id"`
	Title  string    `json:"title"`
	Length time.Time `json:"length"`
	Path   string    `json:"path"`
	Image  string    `json:"img"`
}

type PlaylistResp struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Songs       []Song `json:"songs"`
}
