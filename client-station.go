package firstdue

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type GetStationsRequest struct {
}

type GetStationsResponse struct {
	List  []GetStationsResponseStation `json:"list"`
	Total int                          `json:"total"`
}

type GetStationsResponseStation struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

func (c *Client) GetStations(ctx context.Context, input GetStationsRequest) (output GetStationsResponse, err error) {
	values, err := query.Values(input)
	if err != nil {
		return output, err
	}
	path := "/v1/stations"
	if q := values.Encode(); q != "" {
		path += "?" + q
	}
	err = c.Raw(ctx, http.MethodGet, path, nil, &output)
	if err != nil {
		return output, fmt.Errorf("getdispatches: %w", err)
	}
	return output, nil
}
