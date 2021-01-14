package category

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/bhuvnesh13396/PlayMySong/common/kit"

	"github.com/bhuvnesh13396/PlayMySong/common/auth/token"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var (
	ctx = context.Background()
)

func NewHandler(s Service) http.Handler {
	r := mux.NewRouter()

	opts := []kit.ServerOption{
		kit.ServerBefore(token.HTTPTokenToContext),
	}

	get := kit.NewServer(
		MakeGetEndpoint(s),
		DecodeGetRequest,
		kit.EncodeJSONResponse,
		opts...,
	)

	add := kit.NewServer(
		MakeAddEndpoint(s),
		DecodeAddRequest,
		kit.EncodeJSONResponse,
		opts...,
	)

	update := kit.NewServer(
		MakeUpdateEndpoint(s),
		DecodeUpdateRequest,
		kit.EncodeJSONResponse,
		opts...,
	)

	list := kit.NewServer(
		MakeListEndpoint(s),
		DecodeListRequest,
		kit.EncodeJSONResponse,
		opts...,
	)

	r.Handle("/category", get).Methods(http.MethodGet)
	r.Handle("/category", add).Methods(http.MethodPost)
	r.Handle("/category", list).Methods(http.MethodGet)
	r.Handle("/category", update).Methods(http.MethodPut)

	return r
}

func DecodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getRequest
	err := schema.NewDecoder().Decode(&req, r.URL.Query())
	return req, err
}

func DecodeGetResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp getResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}

func DecodeAddRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req addRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func DecodeAddResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp addResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}

func DecodeUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req updateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func DecodeUpdateResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp updateResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}

func DecodeListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req listRequest
	return req, nil
}

func DecodeListResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp listResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}
