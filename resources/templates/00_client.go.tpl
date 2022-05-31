 package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"
)

// apiVersion is the API version in use by this client.
const apiVersion = "{{.Version}}"

// uaVersion is the userAgent version sent to your API so they can track usage
// of this library.
const uaVersion = "{{.Title}}-{{.Version}}"

const authorizationBearer = "Bearer"

// HTTPDoer is used for making HTTP requests. This implementation is generally
// a *http.Client.
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client manages communication with the Example API.
type Client struct {
	// baseURL is the base url for requests.
	baseURL *url.URL

	// userAgent sets the User-Agent header for requests so you can
	// track usage of the client.
	userAgent string

	// Client is the HTTP Client used to communicate with the API.
	// By default this uses http.DefaultClient, so there are no timeouts
	// configured. It's recommended you set your own HTTP client with
	// reasonable timeouts for your application.
	Client HTTPDoer

	// Client options allow the caller to configure various parts of the
	// request.
	Options Opts

	// Services used for talking with different parts of the API
	// ...
}

type serviceImpl struct {
	client *Client
}

type Opts struct {
	Host          string
	Version       string
	Timeout       time.Duration
	Authorization Authorization
}

type Authorization struct {
	Type  string
	Token string
}

// NewClient returns a new instance of *Client.
func NewClient(opts *Opts) (*Client, error) {
	httpclient := &http.Client{
		Timeout: opts.Timeout,
	}

	// We default to JWT style authentication tokens
	if opts.Authorization.Type == "" {
		opts.Authorization.Type = authorizationBearer
	}

	// Setup default client
	client := &Client{
		Client:  httpclient,
		Options: *opts,
		userAgent: fmt.Sprintf(
			"sgen/%s; Go (%s) [%s-%s]",
			uaVersion,
			runtime.Version(),
			runtime.GOARCH,
			runtime.GOOS,
		),
	}

	//client.Accounts = &accountsImpl{client: client}
	return client, nil
}

// newRequest creates an authenticated API request that is ready to send.
func (c *Client) newRequest(ctx context.Context, method string, path string, body interface{}) (*http.Request, error) {
	switch c.Options.Version {
	case "":
		path = fmt.Sprintf("/%s/%s", c.Options.Version, strings.TrimPrefix(path, "/"))
	default:
		path = fmt.Sprintf("/%s", strings.TrimPrefix(path, "/"))
	}

	u, err := c.baseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	// Request body
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// We default to JWT Bearer: <token> types
	switch c.Options.Authorization.Type {
	case authorizationBearer:
	default:
		req.Header.Set("Authorization", fmt.Sprintf("%s: %s", authorizationBearer, c.Options.Authorization.Token))
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("X-Api-Version", apiVersion)

	if body != nil {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	return req, err
}
