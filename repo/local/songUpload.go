package local

import (
	"io/ioutil"

	"github.com/bhuvnesh13396/PlayMySong/model"
)

type songUpload struct {
}

func NewSongUploadRepo() (*songUpload, error) {
	return &songUpload{}, nil
}

func (song *songUpload) Upload(data model.SongUpload) (err error) {
	path := data.Path
	songFile := data.SongFile
	songID := data.SongID

	// out, err := os.Create(path + "/" + songID)
	// if err != nil {
	// 	return err
	// }
	//defer out.Close()
	ioutil.WriteFile(path+"/"+songID+".wav", songFile, 0644)
	return nil
}
