package local

import (
	"fmt"
	"os"

	"github.com/bhuvnesh13396/PlayMySong/model"
)

type songUpload struct {
}

func NewSongUploadRepo() (*songUpload, error) {
	return &songUpload{}, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (song *songUpload) Upload(data model.SongUpload) (err error) {
	path := data.Path
	songFile := data.SongFile
	songID := data.SongID

	fmt.Println("Repo :: path ", path)
	//fmt.Println("Repo :: path ", )
	fmt.Println("Repo :: songID ", songID)
	// out, err := os.Create(path + "/" + songID)
	// if err != nil {
	// 	return err
	// }
	//defer out.Close()
	f, err := os.Create(path + "/" + songID)
	check(err)
	defer f.Close()

	noOfBytes, err := f.Write(songFile)
	check(err)
	fmt.Println("wrote %d bytes\n", noOfBytes)
	//ioutil.WriteFile(path+"/"+songID+".wav", songFile, 0644)
	return nil
}
