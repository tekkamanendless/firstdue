package firstdue

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/tekkamanendless/httperror"
)

// Raw performs a raw HTTP request to the FirstDue API.
// It handles authentication, request creation, and response parsing.
//
// Input and output are expected to be JSON-serializable structures.  If omitted, they will not be sent or parsed.
func (c *Client) Raw(ctx context.Context, method string, path string, input any, output any) error {
	fullURL := strings.TrimRight(c.config.BaseURL, "/") + "/" + strings.TrimLeft(path, "/")

	var inputReader io.Reader
	if input != nil {
		inputContents, err := json.Marshal(input)
		if err != nil {
			return fmt.Errorf("failed to marshal input: %w", err)
		}
		inputReader = strings.NewReader(string(inputContents))
	}
	request, err := http.NewRequest(strings.ToUpper(method), fullURL, inputReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	if c.config.Token != "" {
		request.Header.Set("Authorization", "Bearer "+c.config.Token)
	}
	if input != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	if c.config.Debug {
		contents, err := httputil.DumpRequest(request, true)
		if err != nil {
			// Oh well; we can't dump the request.
		} else {
			slog.Debug("HTTP request:\n" + string(contents))
		}
	}

	httpClient := c.config.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	response, err := httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer response.Body.Close()

	if c.config.Debug {
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

		var errs []error
		for _, err := range errorResponse.Errors {
			errs = append(errs, fmt.Errorf("%s: %s: %s", err.Field, err.Code, err.Message))
		}
		return fmt.Errorf("%w: %s: %w", httperror.ErrorFromStatus(response.StatusCode), errorResponse.Message, errors.Join(errs...))
	}

	if output != nil {
		if err := json.NewDecoder(response.Body).Decode(output); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
