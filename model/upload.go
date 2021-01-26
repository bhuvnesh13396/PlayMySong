package model

// Model for uploading song...
type SongUpload struct {
	Path     string
	SongID   string
	SongFile []byte
}

// Common interfce to upload the song on a given path
type Uploader interface {
	Upload(song SongUpload) (err error)
}
