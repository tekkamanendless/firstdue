package firstdue

import (
	"net/http"
)

// BaseURL is the default base URL for the FirstDue API.
const BaseURL = "https://sizeup.firstduesizeup.com/fd-api"

type ClientConfig struct {
	BaseURL    string       // The base URL for API requests; if empty, the default will be used.
	Token      string       // The API "Bearer" token for authentication.
	Debug      bool         // If true, debug information will be printed to the log.
	HTTPClient *http.Client // The HTTP client to use.
}

// Client is a client for the FirstDue API.
type Client struct {
	config ClientConfig
}

// ClientOption is a function that configures a Client.
type ClientOption func(*ClientConfig)

// NewClient returns a new client for the FirstDue API.
func NewClient(opts ...ClientOption) *Client {
	config := ClientConfig{}
	for _, opt := range opts {
		opt(&config)
	}
	if config.BaseURL == "" {
		config.BaseURL = BaseURL
	}
	// Do *not* set the HTTP client if one wasn't provided.
	// At run-time, we'll just use the default net/http client, but we won't save it to the config.

	c := &Client{
		config: config,
	}
	return c
}

// WithHTTPClient sets the HTTP client to use for requests.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *ClientConfig) {
		c.HTTPClient = httpClient
	}
}

// WithBaseURL sets the base URL for API requests.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *ClientConfig) {
		c.BaseURL = baseURL
	}
}

// WithToken sets the API "Bearer" token for authentication.
func WithToken(token string) ClientOption {
	return func(c *ClientConfig) {
		c.Token = token
	}
}

// WithDebug sets the debug flag.
func WithDebug(debug bool) ClientOption {
	return func(c *ClientConfig) {
		c.Debug = debug
	}
}

func (c *Client) HTTPClient() *http.Client {
	return c.config.HTTPClient
}

func (c *Client) BaseURL() string {
	return c.config.BaseURL
}

func (c *Client) Token() string {
	return c.config.Token
}

func (c *Client) Debug() bool {
	return c.config.Debug
}
