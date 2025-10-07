package firstdue

import (
	"context"
	"fmt"
	"net/http"
)

type NfirsNotification struct {
	DispatchNumber           string    `json:"dispatch_number"`
	IncidentNumber           string    `json:"incident_number"`
	DispatchType             string    `json:"dispatch_type"`
	DispatchIncidentTypeCode string    `json:"dispatch_incident_type_code"`
	AlarmAt                  Timestamp `json:"alarm_at"`
	DispatchNotifiedAt       Timestamp `json:"dispatch_notified_at"`
	Alarms                   int       `json:"alarms"`
	PlaceName                string    `json:"place_name"`
	LocationInfo             string    `json:"location_info"`
	Venue                    string    `json:"venue"`
	Address                  string    `json:"address"`
	Unit                     string    `json:"unit"`
	CrossStreets             string    `json:"cross_streets"`
	City                     string    `json:"city"`
	StateCode                string    `json:"state_code"`
	Latitude                 float64   `json:"latitude"`
	Longitude                float64   `json:"longitude"`
	Narratives               string    `json:"narratives"`
	ShiftName                string    `json:"shift_name"`
	NotificationType         string    `json:"notification_type"`
}

type PostNfirsNotificationsRequest NfirsNotification

type PostNfirsNotificationsResponse struct {
	ID uint64 `json:"id"`
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

func (c *Client) PostNfirsNotificationsIDApparatuses(ctx context.Context, id uint64, input PostNfirsNotificationsIDApparatusesRequest) error {
	err := c.Raw(ctx, http.MethodPost, fmt.Sprintf("/v1/nfirs-notifications/%d/apparatuses", id), input, nil)
	if err != nil {
		return fmt.Errorf("postnfirsnotificationsidapparatuses: %w", err)
	}
	return nil
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
