package firstdue

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

type Timestamp time.Time

var _ json.Marshaler = (*Timestamp)(nil)
var _ json.Unmarshaler = (*Timestamp)(nil)
var _ query.Encoder = (*Timestamp)(nil)

// IsZero returns true if the timestamp is the zero value.
//
// This is used for "omitempty" support in query parameters.
func (t Timestamp) IsZero() bool {
	return time.Time(t).IsZero()
}

func (t Timestamp) EncodeValues(key string, v *url.Values) error {
	s := time.Time(t).Format(time.RFC3339)
	v.Set(key, s)
	return nil
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	s := time.Time(t).Format(time.RFC3339)
	return json.Marshal(s)
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}
	*t = Timestamp(parsed)
	return nil
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Errors  []struct {
		Field   string `json:"field"`
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors,omitempty"`
}
