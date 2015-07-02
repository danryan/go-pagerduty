package pagerduty

import "net/http"

// AlertsService type
type AlertsService struct {
	client *Client
}

// Alert type
type Alert struct {
	ID                    string            `json:"id,omitempty"`
	Type                  string            `json:"type,omitempty"`
	StartedAt             string            `json:"started_at,omitempty"`
	UUser                 User              `json:"user,omitempty"` // This is conflicting with the package name on assignment in test. Not sure of the soltuion
	Address               string            `json:"address,omitempty"`
}

// Alerts is a list of alerts
type Alerts struct {
	*Pagination
	Alerts []Alert `json:"alerts,omitempty"`
}

// AlertsOptions provides optional parameters to list requests
type AlertsOptions struct {
	Pagination
	Type     string `url:"filter[type],omitempty"` // Can be one of SMS, Email, Phone, or Push
	Timezone string `url:"time_zone,omitempty"`
	Since    string `url:"since,omitempty"`        // Format 2006-01-02T15:04:05-07:00
	Until    string `url:"until,omitempty"`
}

// List returns a list of alerts
func (s *AlertsService) List(opt *AlertsOptions) (Alerts, *http.Response, error) {
	var alerts Alerts

	u, err := addOptions("alerts", opt)
	if err != nil {
		return alerts, nil, err
	}

	res, err := s.client.Get(u, &alerts)
	if err != nil {
		return alerts, res, err
	}

	return alerts, res, err
}
