package kit

import (
	"context"
	"net/http"
	"net/url"
)

type ClientOption func(*client)

type client struct {
	client *http.Client
	method string
	url    *url.URL
	enc    EncodeRequestFunc
	dec    DecodeResponseFunc
	before []RequestFunc
}

func NewClient(method string, url *url.URL, enc EncodeRequestFunc, dec DecodeResponseFunc, options ...ClientOption) *client {
	c := &client{
		client: http.DefaultClient,
		method: method,
		url:    url,
		enc:    enc,
		dec:    dec,
	}
	for _, option := range options {
		option(c)
	}
	return c
}

func (c *client) Endpoint() Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		hreq, err := http.NewRequest(c.method, c.url.String(), nil)
		if err != nil {
			return nil, err
		}

		err = c.enc(ctx, hreq, request)
		if err != nil {
			return nil, err
		}

		for _, f := range c.before {
			ctx = f(ctx, hreq)
		}

		hres, err := c.client.Do(hreq)
		if err != nil {
			return nil, err
		}

		defer hres.Body.Close()

		return c.dec(ctx, hres)
	}
}

func SetClient(httpclient *http.Client) ClientOption {
	return func(c *client) {
		c.client = httpclient
	}
}

func ClientBefore(before ...RequestFunc) ClientOption {
	return func(c *client) {
		c.before = append(c.before, before...)
	}
}
