package firstdue

import (
	"context"
	"fmt"
	"net/http"
)

type GetLogsSettingsRequest struct{}

type GetLogsSettingsResponse []struct {
	IsFdapiLogEnabled bool `json:"is_fdapi_log_enabled"`
}

func (c *Client) GetLogsSettings(ctx context.Context, input GetLogsSettingsRequest) (output GetLogsSettingsResponse, err error) {
	err = c.Raw(ctx, http.MethodGet, "/v1/logs/settings", nil, &output)
	if err != nil {
		return output, fmt.Errorf("getlogssettings: %w", err)
	}
	return output, nil
}

type PostLogsRequest struct {
	Message   string `json:"message"`
	LevelCode string `json:"level_code"`
	Category  string `json:"category"`
}

func (c *Client) PostLogs(ctx context.Context, input PostLogsRequest) error {
	err := c.Raw(ctx, http.MethodPost, "/v1/logs", input, nil)
	if err != nil {
		return fmt.Errorf("postlogs: %w", err)
	}
	return nil
}

type PostLogsBatchRequest []struct {
	Message   string `json:"message"`
	LevelCode string `json:"level_code"`
	Category  string `json:"category"`
}

func (c *Client) PostLogsBatch(ctx context.Context, input PostLogsBatchRequest) error {
	err := c.Raw(ctx, http.MethodPost, "/v1/logs/batch", input, nil)
	if err != nil {
		return fmt.Errorf("postlogsbatch: %w", err)
	}
	return nil
}
