package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/madsaune/snowy/auth"
)

type Client struct {
	UserAgent  string
	Authorizer *auth.Authorizer
	HttpClient *http.Client
}

func New(authorizer *auth.Authorizer) *Client {
	defaultHttpClient := &http.Client{
		Timeout: time.Second * 30,
	}
	return NewWithClient(authorizer, defaultHttpClient)
}

func NewWithClient(authorizer *auth.Authorizer, client *http.Client) *Client {
	return &Client{
		UserAgent:  "Snowy (go-http-client/1.1)",
		Authorizer: authorizer,
		HttpClient: client,
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

	if res.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("error: Unauthorized - Have you set SN_USERNAME and SN_PASSWORD as environment variables?")
	}

	// FIXME: Add additional checks for for most common HTTP Status codes

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: request failed unexpectedly\ncode: %d\nurl: %s\nerror: %v", res.StatusCode, newUrl.String(), err)
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
