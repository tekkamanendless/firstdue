package firstdue

import (
	"net/http"
)

type Client struct {
	BaseURL    string       // The base URL for API requests; if empty, the default will be used.
	Token      string       // The API "Bearer" token for authentication.
	HTTPClient *http.Client // The HTTP client to use; if nil, http.DefaultClient will be used.
	Debug      bool         // If true, debug information will be printed to the log.
}

func NewClient() *Client {
	return &Client{}
}
