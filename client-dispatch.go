package firstdue

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type GetDispatchesRequest struct {
	Page  int       `url:"page,omitempty"`
	Since Timestamp `url:"since,omitempty"`
}

type GetDispatchesResponse []struct {
	ID               int       `json:"id"`
	Type             string    `json:"type"`
	Message          string    `json:"message"`
	Address          string    `json:"address"`
	Address2         string    `json:"address2"`
	City             string    `json:"city"`
	StateCode        string    `json:"state_code"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
	UnitCodes        []string  `json:"unit_codes"`
	IncidentTypeCode string    `json:"incident_type_code"`
	StatusCode       string    `json:"status_code"`
	XrefID           string    `json:"xref_id"`
	CreatedAt        Timestamp `json:"created_at"`
}

func (c *Client) GetDispatches(ctx context.Context, input GetDispatchesRequest) (output GetDispatchesResponse, err error) {
	values, err := query.Values(input)
	if err != nil {
		return output, err
	}
	path := "/v1/dispatches"
	if q := values.Encode(); q != "" {
		path += "?" + q
	}
	err = c.Raw(ctx, http.MethodGet, path, nil, &output)
	if err != nil {
		return output, fmt.Errorf("getdispatches: %w", err)
	}
	return output, nil
}
