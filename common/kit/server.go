package kit

import (
	"context"
	"net/http"
)

type RequestFunc func(context.Context, *http.Request) context.Context

type ServerOption func(*server)

type server struct {
	e      Endpoint
	dec    DecodeRequestFunc
	enc    EncodeResponseFunc
	before []RequestFunc
}

func NewServer(e Endpoint, dec DecodeRequestFunc, enc EncodeResponseFunc, options ...ServerOption) *server {
	s := &server{
		e:   e,
		dec: dec,
		enc: enc,
	}
	for _, option := range options {
		option(s)
	}
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	for _, f := range s.before {
		ctx = f(ctx, r)
	}

	req, err := s.dec(ctx, r)
	if err != nil {
		s.enc(ctx, w, nil, err)
		return
	}

	resp, err := s.e(ctx, req)
	err = s.enc(ctx, w, resp, err)
	if err != nil {
		return
	}
}

func ServerBefore(before ...RequestFunc) ServerOption {
	return func(s *server) {
		s.before = append(s.before, before...)
	}
}
