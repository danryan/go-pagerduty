package pagerduty

import (
	"errors"
	"net/http"
)

// The SchedulesService struct is instantiated in the pagerduty
// struct instantiation and contains a reference back to the pagerduty
// client
type SchedulesService struct {
	client *Client
}

// The ScheduleWrapper type is necessary because the GET /schedules/:id
// API returns a nested json object in the form: {"schedule": {"id": ...}}
type ScheduleWrapper struct {
	Schedule *Schedule `json:"schedule,omitempty"`
}

type Schedule struct {
	ID                   string              `json:"id,omitempty"`
	Name                 string              `json:"name,omitempty"`
	Timezone             string              `json:"time_zone,omitempty"`
	Today                string              `json:"today,omitempty"`
	EscalationPolicies   []*EscalationPolicy `json:"escalation_policies,omitempty"`
	ScheduleLayers       []*ScheduleLayer    `json:"schedule_layers,omitempty"`
	OverridesSubschedule *ScheduleLayer      `json:"overrides_subschedule,omitempty"`
	FinalSchedule        *ScheduleLayer      `json:"final_schedule,omitempty"`
}

type ScheduleLayer struct {
	Name                       string   `json:"name,omitempty"`
	RenderedScheduleEntries    []string `json:"rendered_schedule_entries,omitempty"`
	RestrictionType            string   `json:"restriction_type,omitempty"`
	Restrictions               []string `json:"restrictions,omitempty"`
	Priority                   int      `json:"priority,omitempty"`
	Start                      string   `json:"start,omitempty"`
	End                        string   `json:"end,omitempty"`
	RenderedCoveragePercentage int      `json:"rendered_coverage_percentage,omitempty"`
	RotationTurnLengthSeconds  int      `json:"rotation_turn_length_seconds,omitempty"`
	RotationVirtualStart       string   `json:"rotation_virtual_start,omitempty"`
	Users                      []*User  `json:"users,omitempty"`
}

type Schedules struct {
	Schedules []*Schedule `json:"schedules,omitempty"`
	Limit     int         `json:"limit,omitempty"`
	Offset    int         `json:"offset,omitempty"`
	Total     int         `json:"total,omitempty"`
}

// SchedulesOptions provides optional parameters to list requests
type SchedulesOptions struct {
	Query       string `url:"query,omitempty"`
	RequesterId string `url:"requester_id,omitempty"`
}

// List returns a list of schedules
func (s *SchedulesService) List(opt *SchedulesOptions) (*Schedules, *http.Response, error) {
	uri, err := addOptions("schedules", opt)
	if err != nil {
		return nil, nil, err
	}

	schedules := new(Schedules)
	res, err := s.client.Get(uri, schedules)
	if err != nil {
		return nil, res, err
	}

	return schedules, res, err
}

// Get returns a single schedule by id if found
func (s *SchedulesService) Get(id string) (*Schedule, *http.Response, error) {
	wrapper := new(ScheduleWrapper)

	res, err := s.client.Get("schedules/"+id, wrapper)
	if err != nil {
		return nil, res, err
	}

	if wrapper.Schedule == nil {
		return nil, res, errors.New("pagerduty: schedule json object nil")
	}

	return wrapper.Schedule, res, nil
}

// ScheduleEntriesOptions provides optional parameters to entries requests
type ScheduleEntriesOptions struct {
	Since    string `url:"since,omitempty"`
	Until    string `url:"until,omitempty"`
	Overflow bool   `url:"overflow,omitempty"`
	Timezone string `url:"time_zone,omitempty"`
	UserId   string `url:"user_id,omitempty"`
}

type ScheduleEntries struct {
	ScheduleEntries []*ScheduleEntry `json:"entries,omitempty"`
	Total           int              `json:"total,omitempty"`
}

type ScheduleEntry struct {
	User  *User  `json:"user,omitempty"`
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

// Entries returns a list of schedule entries for a schedule by id
func (s *SchedulesService) Entries(id string, opt *ScheduleEntriesOptions) (*ScheduleEntries, *http.Response, error) {
	uri, err := addOptions("schedules/"+id+"/entries", opt)
	if err != nil {
		return nil, nil, err
	}

	entries := new(ScheduleEntries)
	res, err := s.client.Get(uri, entries)
	if err != nil {
		return nil, res, err
	}

	return entries, res, err
}
