package playlist

import (
	"context"

	"github.com/bhuvnesh13396/PlayMySong/common/kit"
)

type GetEndpoint kit.Endpoint
type AddEndpoint kit.Endpoint
type UpdateEndpoint kit.Endpoint
type ListEndpoint kit.Endpoint

type Endpoint struct {
	GetEndpoint
	AddEndpoint
	UpdateEndpoint
	ListEndpoint
}

type getRequest struct {
	ID string `schema:"id"`
}

type getResponse struct {
	Playlist PlaylistResp `json:"playlist"`
}

func MakeGetEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequest)
		playlist, err := s.Get(ctx, req.ID)
		return getResponse{Playlist: playlist}, err
	}
}

func (e GetEndpoint) Get(ctx context.Context, ID string) (playlist PlaylistResp, err error) {
	request := getRequest{
		ID: ID,
	}
	response, err := e(ctx, request)
	resp := response.(getResponse)
	return resp.Playlist, err
}

type addRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	SongIDs     []string `json:"song_ids"`
}

type addResponse struct {
	ID string `json:"id"`
}

func MakeAddEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest)
		id, err := s.Add(ctx, req.Title, req.Description, req.SongIDs)
		return addResponse{ID: id}, err
	}
}

func (e AddEndpoint) Add(ctx context.Context, title string, description string, songIDs []string) (err error) {
	request := addRequest{
		Title:       title,
		Description: description,
		SongIDs:     songIDs,
	}
	_, err = e(ctx, request)
	return err
}

type updateRequest struct {
	ID      string   `json:"id"`
	SongIDs []string `json:"song_ids"`
}

type updateResponse struct {
}

func MakeUpdateEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequest)
		err := s.Update(ctx, req.ID, req.SongIDs)
		return updateResponse{}, err
	}
}

func (e UpdateEndpoint) Update(ctx context.Context, ID string, songIDs [] string) (err error) {
	request := updateRequest{
		ID:    ID,
		SongIDs: songIDs,
	}

	_, err = e(ctx, request)
	return err
}

type listRequest struct {
}

type listResponse struct {
	Playlists []PlaylistResp `json:"playlists"`
}

func MakeListEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		playlists, err := s.List(ctx)
		return listResponse{Playlists: playlists}, err
	}
}

func (e ListEndpoint) List(ctx context.Context) (res []PlaylistResp, err error) {
	request := getRequest{}
	response, err := e(ctx, request)
	resp := response.(listResponse)
	return resp.Playlists, err
}
