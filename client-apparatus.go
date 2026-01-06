package firstdue

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type GetApparatusesRequest struct {
}

type GetApparatusesResponse struct {
	List  []GetApparatusesResponseApparatus `json:"list"`
	Total int                               `json:"total"`
}

type GetApparatusesResponseApparatus struct {
	UUID     string  `json:"uuid"`
	Name     string  `json:"name"`
	UnitCode *string `json:"unit_code"` // The unit code that maps back to the dispatch information.
	UseCode  string  `json:"use_code"`
	UseName  string  `json:"use_name"`
}

func (c *Client) GetApparatuses(ctx context.Context, input GetApparatusesRequest) (output GetApparatusesResponse, err error) {
	values, err := query.Values(input)
	if err != nil {
		return output, err
	}
	path := "/v1/apparatuses"
	if q := values.Encode(); q != "" {
		path += "?" + q
	}
	err = c.Raw(ctx, http.MethodGet, path, nil, &output)
	if err != nil {
		return output, fmt.Errorf("getdispatches: %w", err)
	}
	return output, nil
}
