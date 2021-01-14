package category

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

type CategoryResp struct {
	Title       string `json:"title"`
	Type        string `json:"type"`
	Songs       []Song `json:"songs"`
}
