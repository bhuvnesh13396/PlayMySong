package like

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

	delete := kit.NewServer(
		MakeDeleteEndpoint(s),
		DecodeDeleteRequest,
		kit.EncodeJSONResponse,
		opts...,
	)

	update := kit.NewServer(
		MakeUpdateCountEndpoint(s),
		DecodeUpdateCountRequest,
		kit.EncodeJSONResponse,
		opts...,
	)

	r.Handle("/like", get).Methods(http.MethodGet)
	r.Handle("/like", add).Methods(http.MethodPost)
	r.Handle("/like", delete).Methods(http.MethodDelete)
	r.Handle("/like", update).Methods(http.MethodPut)

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

func DecodeDeleteRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req deleteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func DecodeDeleteResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp deleteResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}

func DecodeUpdateCountRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req updateCountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func DecodeUpdateCountResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp updateCountResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}
