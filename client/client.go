package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/madsaune/snowy/auth"
)

type Client struct {
	UserAgent  string
	Authorizer auth.Authorizer
	HttpClient *http.Client
}

func NewClient(authorizer auth.Authorizer) Client {
	return Client{
		UserAgent:  "Snowy (Go-http-client/1.1)",
		Authorizer: authorizer,
		HttpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *Client) Get(ctx context.Context, endpoint string, query *url.Values) (*http.Response, error) {
	newUrl, err := url.Parse(c.Authorizer.InstanceURL)
	if err != nil {
		return nil, err
	}

	newUrl.Path = endpoint

	if query != nil {
		newUrl.RawQuery = query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, newUrl.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Authorizer.Username, c.Authorizer.Password)
	req.Header.Add("Accept", "application/json")

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) Post(ctx context.Context, endpoint string, body interface{}) (*http.Response, error) {
	newUrl, err := url.Parse(c.Authorizer.InstanceURL)
	if err != nil {
		return nil, err
	}

	newUrl.Path = endpoint

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, newUrl.String(), bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Authorizer.Username, c.Authorizer.Password)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) Delete(ctx context.Context, endpoint string, query *url.Values) (*http.Response, error) {
	newUrl, err := url.Parse(c.Authorizer.InstanceURL)
	if err != nil {
		return nil, err
	}

	newUrl.Path = endpoint

	if query != nil {
		newUrl.RawQuery = query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, newUrl.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Authorizer.Username, c.Authorizer.Password)
	req.Header.Add("Accept", "application/json")

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) Put(ctx context.Context, endpoint string, body interface{}) (*http.Response, error) {
	newUrl, err := url.Parse(c.Authorizer.InstanceURL)
	if err != nil {
		return nil, err
	}

	newUrl.Path = endpoint

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, newUrl.String(), bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Authorizer.Username, c.Authorizer.Password)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
