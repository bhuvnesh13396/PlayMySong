package account

import (
	"context"
	"time"

	"github.com/bhuvnesh13396/PlayMySong/common/kit"
	"github.com/bhuvnesh13396/PlayMySong/model"
)

type GetEndpoint kit.Endpoint
type Get1Endpoint kit.Endpoint
type AddEndpoint kit.Endpoint
type UpdateEndpoint kit.Endpoint
type ListEndpoint kit.Endpoint

type Endpoint struct {
	GetEndpoint
	Get1Endpoint
	AddEndpoint
	UpdateEndpoint
	ListEndpoint
}

type getRequest struct {
	SongName string `schema:"songName"`
}

type getResponse struct {
	Song model.Song `json:"song"`
}

func MakeGetEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequest)
		song, err := s.Get(ctx, req.SongName)
		return getResponse{Song: song}, err
	}
}

func (e GetEndpoint) Get(ctx context.Context, songName string) (song model.Song, err error) {
	request := getRequest{
		SongName: songName,
	}
	response, err := e(ctx, request)
	resp := response.(getResponse)
	return resp.Song, err
}

type get1Request struct {
	ID string `schema:"id"`
}

type get1Response struct {
	Song model.Song `json:"song"`
}

func MakeGet1Endpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(get1Request)
		song, err := s.Get1(ctx, req.ID)
		return get1Response{Song: song}, err
	}
}

func (e Get1Endpoint) Get1(ctx context.Context, id string) (account model.Account, err error) {
	request := get1Request{
		ID: id,
	}
	response, err := e(ctx, request)
	resp := response.(get1Response)
	return resp.Song, err
}

type addRequest struct {
	Title      string    `json:"title"`
	Length     time.Time `json:"length"`
	ArtistID   string    `json:"artist_id"`
	ComposerID string    `json:"composer_id"`
	Lyrics     string    `json:"lyrics"`
	Path       string    `json:"path"`
	Img        string    `json:img`
}

type addResponse struct {
	ID string `json:"id"`
}

func MakeAddEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest)
		id, err := s.Add(ctx, req.Title, req.Length, req.ArtistID, req.ComposerID, req.Lyrics, req.Path, req.Img)
		return addResponse{ID: id}, err
	}
}

func (e AddEndpoint) Add(ctx context.Context, title string, length time.Time, artistID string, composerID string, lyrics string, path string, img string) (err error) {
	request := addRequest{
		Title:      title,
		Length:     length,
		ArtistID:   artistID,
		ComposerID: composerID,
		Lyrics:     lyrics,
		Path:       path,
		Img:        img,
	}
	_, err = e(ctx, request)
	return err
}

type updateRequest struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type updateResponse struct {
}

func MakeUpdateEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequest)
		err := s.Update(ctx, req.Id, req.Title)
		return updateResponse{}, err
	}
}

func (e UpdateEndpoint) Update(ctx context.Context, id string, title string) (err error) {
	request := updateRequest{
		Id:    id,
		Title: title,
	}

	_, err = e(ctx, request)
	return err
}

type listRequest struct {
}

type listResponse struct {
	Songs []SongResp `json:"Songs"`
}

func MakeListEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		songs, err := s.List(ctx)
		return listResponse{Songs: songs}, err
	}
}

func (e ListEndpoint) List(ctx context.Context) (res []SongResp, err error) {
	request := getRequest{}
	response, err := e(ctx, request)
	resp := response.(listResponse)
	return resp.Songs, err
}
