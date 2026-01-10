package firstdue

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type StringFloat64 float64

func (f StringFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%f", f))
}

func (f *StringFloat64) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}
	var v float64
	if err := json.Unmarshal([]byte(s), &v); err != nil {
		return err
	}
	*f = StringFloat64(v)
	return nil
}

type StringUint64 uint64

func (n StringUint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%d", n))
}

func (n *StringUint64) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}
	var v int
	if err := json.Unmarshal([]byte(s), &v); err != nil {
		return err
	}
	*n = StringUint64(v)
	return nil
}

type NfirsNotification struct {
	ID                       uint64         `json:"id,omitempty"`
	DispatchNumber           string         `json:"dispatch_number"`
	IncidentNumber           string         `json:"incident_number"`
	DispatchType             string         `json:"dispatch_type"`
	DispatchIncidentTypeCode string         `json:"dispatch_incident_type_code"`
	AlarmAt                  Timestamp      `json:"alarm_at"`
	DispatchNotifiedAt       Timestamp      `json:"dispatch_notified_at"`
	Alarms                   int            `json:"alarms"`
	CADPriority              *string        `json:"cad_priority"`
	PlaceName                *string        `json:"place_name"`
	BusinessName             *string        `json:"business_name"`
	LocationInfo             *string        `json:"location_info"`
	Venue                    *string        `json:"venue"`
	Address                  string         `json:"address"`
	Unit                     *string        `json:"unit"`
	CrossStreets             string         `json:"cross_streets"`
	City                     string         `json:"city"`
	StateCode                string         `json:"state_code"`
	ZipCode                  *string        `json:"zip_code"`
	Latitude                 *StringFloat64 `json:"latitude"`
	Longitude                *StringFloat64 `json:"longitude"`
	Narratives               *string        `json:"narratives"`
	ShiftName                *string        `json:"shift_name"`
	NotificationType         *string        `json:"notification_type"`
	AidTypeCode              *string        `json:"aid_type_code"`
	AidFDIDNumber            *string        `json:"aid_fdid_number"`
	AidFDIDNumbers           []string       `json:"aid_fdid_numbers"`
	ControlledAt             Timestamp      `json:"controlled_at"`
	OfficerInCharge          *string        `json:"officer_in_charge"`
	CallCompletedAt          Timestamp      `json:"call_completed_at"`
	Zone                     *string        `json:"zone"`
	HouseNum                 *string        `json:"house_num"`
	PrefixDirection          *string        `json:"prefix_direction"`
	StreetName               *string        `json:"street_name"`
	StreetType               *string        `json:"street_type"`
	SuffixDirection          *string        `json:"suffix_direction"`
	EMSIncidentNumber        *string        `json:"ems_incident_number"`
	EMSResponseNumber        *string        `json:"ems_response_number"`
	Station                  *string        `json:"station"`
	EMDCardNumber            *string        `json:"emd_card_number"`
	PSAPAnsweredAt           *Timestamp     `json:"psap_answered_at"`
}

type PostNfirsNotificationsRequest NfirsNotification

type PostNfirsNotificationsResponse struct {
	ID StringUint64 `json:"id"`
}

func (c *Client) PostNfirsNotifications(ctx context.Context, input PostNfirsNotificationsRequest) (output PostNfirsNotificationsResponse, err error) {
	err = c.Raw(ctx, http.MethodPost, "/v1/nfirs-notifications", input, &output)
	if err != nil {
		return output, fmt.Errorf("postnfirsnotifications: %w", err)
	}
	return output, nil
}

func (c *Client) DeleteNfirsNotificationsID(ctx context.Context, id uint64) error {
	err := c.Raw(ctx, http.MethodDelete, fmt.Sprintf("/v1/nfirs-notifications/%d", id), nil, nil)
	if err != nil {
		return fmt.Errorf("deletenfirsnotifications: %w", err)
	}
	return nil
}

type GetNfirsNotificationsIDResponse NfirsNotification

func (c *Client) GetNfirsNotificationsID(ctx context.Context, id uint64) (output GetNfirsNotificationsIDResponse, err error) {
	err = c.Raw(ctx, http.MethodGet, fmt.Sprintf("/v1/nfirs-notifications/%d", id), nil, &output)
	if err != nil {
		return output, fmt.Errorf("getnfirsnotificationsid: %w", err)
	}
	return output, nil
}

type PutNfirsNotificationsIDRequest NfirsNotification

func (c *Client) PutNfirsNotificationsID(ctx context.Context, id uint64, input PutNfirsNotificationsIDRequest) error {
	err := c.Raw(ctx, http.MethodPut, fmt.Sprintf("/v1/nfirs-notifications/%d", id), input, nil)
	if err != nil {
		return fmt.Errorf("putnfirsnotificationsid: %w", err)
	}
	return nil
}

func (c *Client) DeleteNfirsNotificationsNumberID(ctx context.Context, id string) error {
	err := c.Raw(ctx, http.MethodDelete, fmt.Sprintf("/v1/nfirs-notifications/number/%s", id), nil, nil)
	if err != nil {
		return fmt.Errorf("deletenfirsnotificationsnumberid: %w", err)
	}
	return nil
}

type PutNfirsNotificationsNumberIDRequest NfirsNotification

func (c *Client) PutNfirsNotificationsNumberID(ctx context.Context, id string, input PutNfirsNotificationsNumberIDRequest) error {
	err := c.Raw(ctx, http.MethodPut, fmt.Sprintf("/v1/nfirs-notifications/number/%s", id), input, nil)
	if err != nil {
		return fmt.Errorf("putnfirsnotificationsnumberid: %w", err)
	}
	return nil
}

type NfirsNotificationApparatus struct {
	UnitCode               string    `json:"unit_code"`
	IsAid                  bool      `json:"is_aid"`
	DispatchAt             Timestamp `json:"dispatch_at"`
	ArriveAt               Timestamp `json:"arrive_at"`
	DispatchAcknowledgedAt Timestamp `json:"dispatch_acknowledged_at"`
	EnrouteAt              Timestamp `json:"enroute_at"`
	ClearAt                Timestamp `json:"clear_at"`
	BackInServiceAt        Timestamp `json:"back_in_service_at"`
	CanceledAt             Timestamp `json:"canceled_at"`
	CanceledStageCode      string    `json:"canceled_stage_code"`
}

type PostNfirsNotificationsIDApparatusesRequest NfirsNotificationApparatus

type PostNfirsNotificationsIDApparatusesResponse struct {
	ID StringUint64 `json:"id"`
}

func (c *Client) PostNfirsNotificationsIDApparatuses(ctx context.Context, id uint64, input PostNfirsNotificationsIDApparatusesRequest) (output PostNfirsNotificationsIDApparatusesResponse, err error) {
	err = c.Raw(ctx, http.MethodPost, fmt.Sprintf("/v1/nfirs-notifications/%d/apparatuses", id), input, &output)
	if err != nil {
		return output, fmt.Errorf("postnfirsnotificationsidapparatuses: %w", err)
	}
	return output, nil
}

type PutNfirsNotificationsIDApparatusesIDRequest NfirsNotificationApparatus

func (c *Client) PutNfirsNotificationsIDApparatusesID(ctx context.Context, id uint64, apparatusID uint64, input PutNfirsNotificationsIDApparatusesIDRequest) error {
	err := c.Raw(ctx, http.MethodPut, fmt.Sprintf("/v1/nfirs-notifications/%d/apparatuses/%d", id, apparatusID), input, nil)
	if err != nil {
		return fmt.Errorf("putnfirsnotificationsidapparatusesid: %w", err)
	}
	return nil
}

func (c *Client) DeleteNfirsNotificationsIDApparatusesID(ctx context.Context, id uint64, apparatusID uint64) error {
	err := c.Raw(ctx, http.MethodDelete, fmt.Sprintf("/v1/nfirs-notifications/%d/apparatuses/%d", id, apparatusID), nil, nil)
	if err != nil {
		return fmt.Errorf("deletenfirsnotificationsidapparatusesid: %w", err)
	}
	return nil
}

type GetNfirsNotificationsDispatchNumberIDRequest struct{}

type GetNfirsNotificationsDispatchNumberIDResponse NfirsNotification

func (c *Client) GetNfirsNotificationsDispatchNumberID(ctx context.Context, dispatchNumber string, input GetNfirsNotificationsDispatchNumberIDRequest) (output GetNfirsNotificationsDispatchNumberIDResponse, err error) {
	path := "/v1/nfirs-notifications/dispatch-number/" + url.PathEscape(dispatchNumber)
	err = c.Raw(ctx, http.MethodGet, path, nil, &output)
	if err != nil {
		return output, fmt.Errorf("getnfirsnotificationsdispatchnumberid: %w", err)
	}
	return output, nil
}

type PostNfirsNotificationsNumberIDApparatusesRequest NfirsNotificationApparatus

func (c *Client) PostNfirsNotificationsNumberIDApparatuses(ctx context.Context, id string, input PostNfirsNotificationsNumberIDApparatusesRequest) error {
	err := c.Raw(ctx, http.MethodPost, fmt.Sprintf("/v1/nfirs-notifications/number/%s/apparatuses", id), input, nil)
	if err != nil {
		return fmt.Errorf("postnfirsnotificationsnumberidapparatuses: %w", err)
	}
	return nil
}

type PutNfirsNotificationsNumberIDApparatusesCodeIDRequest NfirsNotificationApparatus

func (c *Client) PutNfirsNotificationsNumberIDApparatusesCodeID(ctx context.Context, id string, apparatusID string, input PutNfirsNotificationsNumberIDApparatusesCodeIDRequest) error {
	err := c.Raw(ctx, http.MethodPut, fmt.Sprintf("/v1/nfirs-notifications/number/%s/apparatuses/code/%s", id, apparatusID), input, nil)
	if err != nil {
		return fmt.Errorf("putnfirsnotificationsnumberidapparatusescodeid: %w", err)
	}
	return nil
}

func (c *Client) DeleteNfirsNotificationsNumberIDApparatusesCodeID(ctx context.Context, id string, apparatusID string) error {
	err := c.Raw(ctx, http.MethodDelete, fmt.Sprintf("/v1/nfirs-notifications/number/%s/apparatuses/code/%s", id, apparatusID), nil, nil)
	if err != nil {
		return fmt.Errorf("deletenfirsnotificationsnumberidapparatusescodeid: %w", err)
	}
	return nil
}
