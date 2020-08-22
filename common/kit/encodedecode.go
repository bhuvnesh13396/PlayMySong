package kit

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
)

type DecodeRequestFunc func(context.Context, *http.Request) (request interface{}, err error)

type EncodeRequestFunc func(context.Context, *http.Request, interface{}) error

type EncodeResponseFunc func(context.Context, http.ResponseWriter, interface{}, error) error

type DecodeResponseFunc func(context.Context, *http.Response) (response interface{}, err error)

type ErrResp struct {
	Error string `json:"error"`
}

func EncodeJSONRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func EncodeSchemaRequest(_ context.Context, r *http.Request, request interface{}) error {
	v := url.Values{}
	err := schema.NewEncoder().Encode(request, v)
	if err != nil {
		return err
	}
	r.URL.RawQuery = v.Encode()
	return nil
}

func EncodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}, err error) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return enc.Encode(ErrResp{err.Error()})
	}
	return enc.Encode(response)
}

func DecodeResponse(ctx context.Context, r *http.Response, resp interface{}) error {
	enc := json.NewDecoder(r.Body)
	if r.StatusCode != 200 {
		var e ErrResp
		err := enc.Decode(&e)
		if err != nil {
			return err
		}
		return errors.New(e.Error)

	}
	return enc.Decode(&resp)
}
