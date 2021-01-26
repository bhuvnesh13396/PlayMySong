package upload

import (
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/bhuvnesh13396/PlayMySong/common/auth/token"
	"github.com/bhuvnesh13396/PlayMySong/common/kit"

	"github.com/gorilla/mux"
)

var (
	ctx = context.Background()
)

func NewHandler(s Service) http.Handler {
	r := mux.NewRouter()

	opts := []kit.ServerOption{
		kit.ServerBefore(token.HTTPTokenToContext),
	}

	add := kit.NewServer(
		MakeAddEndpoint(s),
		DecodeAddRequest,
		kit.EncodeJSONResponse,
		opts...,
	)

	r.Handle("/upload", add).Methods(http.MethodPost)
	return r
}

func DecodeAddRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req addRequest
	// ParseMultipartForm parses a request body as multipart/form-data
	r.ParseMultipartForm(32 << 20)
	songFile, _, err := r.FormFile("song_file")
	songID := r.FormValue("song_id")

	if err != nil {
		return "", err
	}

	copiedFile, err := copySongFile(songFile)
	if err != nil {
		return "", err
	}

	defer songFile.Close() // Close the file when we finish
	req = addRequest{
		SongID:   songID,
		SongFile: copiedFile,
	}

	return req, err
}

func DecodeAddResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp addResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}

func copySongFile(songFile multipart.File) (copiedSong []byte, err error) {

	// Close fi on exit and check for its returned error
	defer func() {
		if err := songFile.Close(); err != nil {
			panic(err)
		}
	}()

	copiedSong, err = ioutil.ReadAll(songFile)
	if err != nil {
		fmt.Println(err)
	}

	return copiedSong, nil
}
