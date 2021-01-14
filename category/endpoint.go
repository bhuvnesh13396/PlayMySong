package category

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
	Category CategoryResp `json:"category"`
}

func MakeGetEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequest)
		category, err := s.Get(ctx, req.ID)
		return getResponse{Category: category}, err
	}
}

func (e GetEndpoint) Get(ctx context.Context, ID string) (category CategoryResp, err error) {
	request := getRequest{
		ID: ID,
	}
	response, err := e(ctx, request)
	resp := response.(getResponse)
	return resp.Category, err
}

type addRequest struct {
	Title       string   `json:"title"`
	Type        string   `json:"type"`
	SongIDs     []string `json:"song_ids"`
}

type addResponse struct {
	ID string `json:"id"`
}

func MakeAddEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest)
		id, err := s.Add(ctx, req.Title, req.Type, req.SongIDs)
		return addResponse{ID: id}, err
	}
}

func (e AddEndpoint) Add(ctx context.Context, title string, type string, songIDs []string) (err error) {
	request := addRequest{
		Title:       title,
		Type:        type,
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
	Categories []CategoryResp `json:"categories"`
}

func MakeListEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		categories, err := s.List(ctx)
		return listResponse{Categories: categories}, err
	}
}

func (e ListEndpoint) List(ctx context.Context) (res []CategoryResp, err error) {
	request := getRequest{}
	response, err := e(ctx, request)
	resp := response.(listResponse)
	return resp.Categories, err
}
