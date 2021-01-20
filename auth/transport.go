package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/bhuvnesh13396/PlayMySong/common/kit"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var (
	ctx = context.Background()
)

func NewHandler(s Service) http.Handler {
	r := mux.NewRouter()

	signin := kit.NewServer(
		MakeSigninEndpoint(s),
		DecodeSigninRequest,
		kit.EncodeJSONResponse,
	)

	verifyToken := kit.NewServer(
		MakeVerifyTokenEndpoint(s),
		DecodeVerifyTokenRequest,
		kit.EncodeJSONResponse,
	)

	r.Handle("/auth/signin", signin).Methods(http.MethodPost)
	r.Handle("/auth/verify", verifyToken).Methods(http.MethodGet)

	return r
}

func DecodeSigninRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req signinRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func DecodeSigninResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp signinResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}

func DecodeVerifyTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req verifyTokenRequest
	err := schema.NewDecoder().Decode(&req, r.URL.Query())
	return req, err
}

func DecodeVerifyTokenResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp verifyTokenResponse
	err := kit.DecodeResponse(ctx, r, &resp)
	return resp, err
}
