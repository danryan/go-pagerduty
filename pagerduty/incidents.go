package pagerduty

import "net/http"

// IncidentsService type
type IncidentsService struct {
	client *Client
}

// Incident type
type Incident struct {
	ID                    string            `json:"id,omitempty"`
	IncidentNumber        int               `json:"incident_number,omitempty"`
	Status                string            `json:"status,omitempty"`
	CreatedOn             string            `json:"created_on,omitempty"`
	Summary               *IncidentSummary  `json:"trigger_summary_data,omitempty"`
	User                  *User             `json:"assigned_to_user,omitempty"`
	SService               *Service          `json:"service,omitempty"` // This is conflicting with the package name on assignment in test. Not sure of the soltuion
	EEscalationPolicy      *EscalationPolicy `json:"escalation_policy,omitempty"` // This is conflicting with the package name on assignment in test. Not sure of the soltuion
	HTMLURL               string            `json:"html_url,omitempty"`
	IncidentKey           string            `json:"incident_key,omitempty"`
	TriggerDetailsHTMLURL string            `json:"trigger_details_html_url,omitempty"`
	TriggerType           string            `json:"trigger_type,omitempty"`
	LastStatusChangeOn    string            `json:"last_status_change_on,omitempty"`
	LastStatusChangeBy    *User             `json:"last_status_change_by,omitempty"`
	NumberOfEscalations   int               `json:"number_of_escalations,omitempty"`
	ResolvedByUser        *User             `json:"resolved_by_user,omitempty"`
	AssignedToUser        *User             `json:"assigned_to_user,omitempty"`
	AssignedTo            []*User           `json:"assigned_to,omitempty"`
}

// Incidents is a list of incidents
type Incidents struct {
	Pagination
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
	Pagination
	Status string `url:"status,omitempty"`
	SortBy string `url:"sort_by,omitempty"`
	Since  string `url:"since,omitempty"`
	Until  string `url:"until,omitempty"`
}

// List returns a list of incidents
func (s *IncidentsService) List(opt *IncidentsOptions) (Incidents, *http.Response, error) {
	var incidents Incidents
	u, err := addOptions("incidents", opt)
	if err != nil {
		return incidents, nil, err
	}


	res, err := s.client.Get(u, &incidents)
	if err != nil {
		return incidents, res, err
	}

	return incidents, res, err
}
