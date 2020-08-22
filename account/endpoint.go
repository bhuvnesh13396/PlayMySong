package account

import (
	"context"

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
	Username string `schema:"username"`
}

type getResponse struct {
	Account model.Account `json:"account"`
}

func MakeGetEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequest)
		account, err := s.Get(ctx, req.Username)
		return getResponse{Account: account}, err
	}
}

func (e GetEndpoint) Get(ctx context.Context, username string) (account model.Account, err error) {
	request := getRequest{
		Username: username,
	}
	response, err := e(ctx, request)
	resp := response.(getResponse)
	return resp.Account, err
}

type get1Request struct {
	ID string `schema:"id"`
}

type get1Response struct {
	Account model.Account `json:"account"`
}

func MakeGet1Endpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(get1Request)
		account, err := s.Get1(ctx, req.ID)
		return get1Response{Account: account}, err
	}
}

func (e Get1Endpoint) Get1(ctx context.Context, id string) (account model.Account, err error) {
	request := get1Request{
		ID: id,
	}
	response, err := e(ctx, request)
	resp := response.(get1Response)
	return resp.Account, err
}

type addRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type addResponse struct {
}

func MakeAddEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest)
		err := s.Add(ctx, req.Name, req.Username, req.Password)
		return addResponse{}, err
	}
}

func (e AddEndpoint) Add(ctx context.Context, name string, username string, password string) (err error) {
	request := addRequest{
		Name:     name,
		Username: username,
		Password: password,
	}
	_, err = e(ctx, request)
	return err
}

type updateRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

type updateResponse struct {
}

func MakeUpdateEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequest)
		err := s.Update(ctx, req.Username, req.Name)
		return updateResponse{}, err
	}
}

func (e UpdateEndpoint) Update(ctx context.Context, username string, name string) (err error) {
	request := updateRequest{
		Username: username,
		Name:     name,
	}

	_, err = e(ctx, request)
	return err
}

type listRequest struct {
}

type listResponse struct {
	Accounts []model.Account `json:"accounts"`
}

func MakeListEndpoint(s Service) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		accounts, err := s.List(ctx)
		return listResponse{Accounts: accounts}, err
	}
}

func (e ListEndpoint) List(ctx context.Context) (res []model.Account, err error) {
	request := getRequest{}
	response, err := e(ctx, request)
	resp := response.(listResponse)
	return resp.Accounts, err
}
