package firstdue

import (
	"net/http"
)

// Client is a client for the FirstDue API.
type Client struct {
	BaseURL    string       // The base URL for API requests; if empty, the default will be used.
	Token      string       // The API "Bearer" token for authentication.
	Debug      bool         // If true, debug information will be printed to the log.
	httpClient *http.Client // The HTTP client to use.
}

// ClientOption is a function that configures a Client.
type ClientOption func(*Client)

// NewClient returns a new client for the FirstDue API.
func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		httpClient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// WithHTTPClient sets the HTTP client to use for requests.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}
