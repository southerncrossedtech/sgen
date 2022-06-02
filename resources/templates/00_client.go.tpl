// Code generated by sgen/v{{.SGenVersion}} (https://github.com/southerncrossedtech/sgen). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.
package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
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

// authorizationTokenPrefix provides a default fallback for jwt style authentication
// but can be overridden with custom values.
const authorizationTokenPrefix = "Bearer"

// HTTPDoer is used for making HTTP requests. This implementation is generally
// a *http.Client.
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client manages communication with the {{.Title}} API.
type Client struct {
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
	// Host is the base url for requests.
	Host          *url.URL
	Version       string
	Authorization Authorization
}

type Authorization struct {
	Prefix string
	Token  string
}

// NewClient returns a new instance of *Client.
func NewClient(opts *Opts) (*Client, error) {
	// Setup a sensible default http client
	httpclient := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}

	// We default to JWT style authentication tokens
	if opts.Authorization.Prefix == "" {
		opts.Authorization.Prefix = authorizationTokenPrefix
	}

	// Setup default client
	client := &Client{
		Client:  httpclient,
		Options: *opts,
		userAgent: fmt.Sprintf(
			"sgen/{{.Title}} %s; Go (%s) [%s-%s]",
			uaVersion,
			runtime.Version(),
			runtime.GOARCH,
			runtime.GOOS,
		),
	}

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

	u, err := c.Options.Host.Parse(path)
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
	switch c.Options.Authorization.Prefix {
	case authorizationTokenPrefix:
	default:
		req.Header.Set("Authorization", fmt.Sprintf("%s: %s", authorizationTokenPrefix, c.Options.Authorization.Token))
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("X-Api-Version", apiVersion)

	if body != nil {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	return req, err
}
