package pagerduty

import "net/http"

// IncidentsService type
type IncidentsService struct {
	client *Client
}

// Incident type
type Incident struct {
	ID             string           `json:"id,omitempty"`
	Status         string           `json:"status,omitempty"`
	IncidentNumber int              `json:"incident_number,omitempty"`
	CreatedOn      string           `json:"created_on,omitempty"`
	Summary        *IncidentSummary `json:"trigger_summary_data,omitempty"`
	User           *User            `json:"assigned_to_user,omitempty"`
}

// Incidents is a list of incidents
type Incidents struct {
	Incidents []Incident
}

// IncidentSummary type
type IncidentSummary struct {
	Subject     string //`json:"subject,omitempty"`
	Description string //`json:"description,omitempty"`
}

// Get returns a single incident by id if found
func (s *IncidentsService) Get(id string) (*Incident, *http.Response, error) {
	incident := new(Incident)

	res, err := s.client.Get("incidents/"+id, incident)
	if err != nil {
		return nil, res, err
	}

	return incident, res, nil
}

// IncidentsOptions provides optional parameters to list requests
type IncidentsOptions struct {
	Status string `url:"status,omitempty"`
	SortBy string `url:"sort_by,omitempty"`
}

// List returns a list of incidents
func (s *IncidentsService) List(opt *IncidentsOptions) ([]Incident, *http.Response, error) {
	u, err := addOptions("incidents", opt)
	if err != nil {
		return nil, nil, err
	}

	incidents := new(Incidents)

	res, err := s.client.Get(u, incidents)
	if err != nil {
		return nil, res, err
	}

	return incidents.Incidents, res, err
}
