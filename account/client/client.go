package client

import (
	"net/http"
	"net/url"

	"github.com/bhuvnesh13396/PlayMySong/account"
	"github.com/bhuvnesh13396/PlayMySong/common/auth/token"
	"github.com/bhuvnesh13396/PlayMySong/common/kit"
)

func New(instance string, client *http.Client) (account.Service, error) {
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}

	opts := []kit.ClientOption{
		kit.SetClient(client),
		kit.ClientBefore(token.HTTPTokenFromContext),
	}

	getEndpoint := kit.NewClient(
		http.MethodGet,
		copyURL(u, "/account"),
		kit.EncodeSchemaRequest,
		account.DecodeGetResponse,
		opts...,
	).Endpoint()

	get1Endpoint := kit.NewClient(
		http.MethodGet,
		copyURL(u, "/account/1"),
		kit.EncodeSchemaRequest,
		account.DecodeGet1Response,
		opts...,
	).Endpoint()

	addEndPoint := kit.NewClient(
		http.MethodPost,
		copyURL(u, "/account"),
		kit.EncodeJSONRequest,
		account.DecodeAddResponse,
		opts...,
	).Endpoint()

	updateEndpoint := kit.NewClient(
		http.MethodPut,
		copyURL(u, "/account"),
		kit.EncodeJSONRequest,
		account.DecodeUpdateResponse,
		opts...,
	).Endpoint()

	listEndpoint := kit.NewClient(
		http.MethodGet,
		copyURL(u, "/account/all"),
		kit.EncodeSchemaRequest,
		account.DecodeListResponse,
		opts...,
	).Endpoint()

	return &account.Endpoint{
		GetEndpoint:    account.GetEndpoint(getEndpoint),
		Get1Endpoint:   account.Get1Endpoint(get1Endpoint),
		AddEndpoint:    account.AddEndpoint(addEndPoint),
		UpdateEndpoint: account.UpdateEndpoint(updateEndpoint),
		ListEndpoint:   account.ListEndpoint(listEndpoint),
	}, nil
}

func copyURL(u *url.URL, path string) *url.URL {
	c := *u
	c.Path = path
	return &c
}
