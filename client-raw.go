package firstdue

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/tekkamanendless/httperror"
)

// BaseURL is the default base URL for the FirstDue API.
const BaseURL = "https://sizeup.firstduesizeup.com/fd-api"

// Raw performs a raw HTTP request to the FirstDue API.
// It handles authentication, request creation, and response parsing.
//
// Input and output are expected to be JSON-serializable structures.  If omitted, they will not be sent or parsed.
func (c *Client) Raw(ctx context.Context, method string, path string, input any, output any) error {
	baseURL := c.BaseURL
	if baseURL == "" {
		baseURL = BaseURL
	}
	fullURL := strings.TrimRight(baseURL, "/") + "/" + strings.TrimLeft(path, "/")

	var inputReader io.Reader
	if input != nil {
		inputContents, err := json.Marshal(input)
		if err != nil {
			return fmt.Errorf("failed to marshal input: %w", err)
		}
		inputReader = strings.NewReader(string(inputContents))
	}
	request, err := http.NewRequest(method, fullURL, inputReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	if c.Token != "" {
		request.Header.Set("Authorization", "Bearer "+c.Token)
	}
	if input != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	if c.Debug {
		contents, err := httputil.DumpRequest(request, true)
		if err != nil {
			// Oh well; we can't dump the request.
		} else {
			slog.Debug("HTTP request:\n" + string(contents))
		}
	}

	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	response, err := httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer response.Body.Close()

	if c.Debug {
		contents, err := httputil.DumpResponse(response, true)
		if err != nil {
			// Oh well; we can't dump the response.
		} else {
			slog.Debug("HTTP response:\n" + string(contents))
		}
	}

	if response.StatusCode < 200 {
		return httperror.ErrorFromStatus(response.StatusCode)
	}

	if response.StatusCode >= 300 {
		var errorResponse ErrorResponse
		if err := json.NewDecoder(response.Body).Decode(&errorResponse); err != nil {
			return fmt.Errorf("%w: (failed to decode error response)", httperror.ErrorFromStatus(response.StatusCode))
		}
		return fmt.Errorf("%w: %s", httperror.ErrorFromStatus(response.StatusCode), errorResponse.Message)
	}

	if output != nil {
		if err := json.NewDecoder(response.Body).Decode(output); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
