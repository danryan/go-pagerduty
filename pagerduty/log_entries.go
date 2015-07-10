package pagerduty

import (
    "time"
    "net/http"
)

// The SchedulesService struct is instantiated in the pagerduty
// struct instantiation and contains a reference back to the pagerduty
// client
type LogEntriesService struct {
    client *Client
}

type Channel struct {
    Subject string `json:"subject,omitempty"`
    Details string `json:"details,omitempty"`
    Summary string `json:"summary,omitempty"`
    Type string `json:"type,omitempty"`
}

type Context struct {
    Type string `json:"type,omitempty"`
    Href string `json:"href,omitempty"`
}

type Notification struct {
    Type string `json:"type,omitempty"`
    Address string `json:"address,omitempty"`
    Status string `json:"status,omitempty"`
}

type LogEntry struct {
    ID string `json:"id"`
    Type string `json:"type,omitempty"`
    CreatedAt time.Time `json:"created_at"`
    Note string `json:"note,omitempty"`
    Agent *User `json:"agent,omitempty"`
    Channel *Channel `json:"channel,omitempty"`
    Contexts []Context `json:"contexts"`
    AssignedUser *User `json:"assigned_user,omitempty"`
    User *User `json:"user,omitempty"`
    Notification *Notification `json:"notification,omitempty"`
}

// Group
type LogEntries struct {
    Pagination
    LogEntries []LogEntry `json:"log_entries"`
}

// Individual
type LogEntryResponse struct {
    LogEntry LogEntry `json:"log_entry"`
}

// ScheduleEntriesOptions provides optional parameters to entries requests
type LogEntriesOptions struct {
    Pagination
    Timezone string `url:"time_zone,omitempty"`
    Since    string `url:"since,omitempty"`
    Until    string `url:"until,omitempty"`
    is_overview bool   `url:"is_overview,omitempty"`
    include []string `url:"include,omitempty"`
}

// ScheduleEntriesOptions provides optional parameters to entries requests
type LogEntryOptions struct {
    Pagination
    Timezone string `url:"time_zone,omitempty"`
    Since    string `url:"since,omitempty"`
    Until    string `url:"until,omitempty"`
    is_overview bool   `url:"is_overview,omitempty"`
    include []string `url:"include,omitempty"`
}

// LogEntries returns a list of log entries for the search
func (s *LogEntriesService) LogEntries(opt *LogEntriesOptions) (*LogEntries, *http.Response, error) {
    uri, err := addOptions("log_entries", opt)
    if err != nil {
        return nil, nil, err
    }

    entries := new(LogEntries)
    res, err := s.client.Get(uri, entries)
    if err != nil {
        return nil, res, err
    }

    return entries, res, err
}

// LogEntries returns a list of log entries for the search
func (s *LogEntriesService) UserLogEntries(user_id string, opt *LogEntriesOptions) (*LogEntries, *http.Response, error) {
    uri, err := addOptions("users/"+user_id+"/log_entries", opt)
    if err != nil {
        return nil, nil, err
    }

    entries := new(LogEntries)
    res, err := s.client.Get(uri, entries)
    if err != nil {
        return nil, res, err
    }

    return entries, res, err
}

// LogEntries returns a list of log entries for the search
func (s *LogEntriesService) IncidentLogEntries(incident_id string, opt *LogEntriesOptions) (*LogEntries, *http.Response, error) {
    uri, err := addOptions("incidents/"+incident_id+"/log_entries", opt)
    if err != nil {
        return nil, nil, err
    }

    entries := new(LogEntries)
    res, err := s.client.Get(uri, entries)
    if err != nil {
        return nil, res, err
    }

    return entries, res, err
}

// LogEntries returns a list of log entries for the search
func (s *LogEntriesService) LogEntry(id string, opt *LogEntryOptions) (*LogEntry, *http.Response, error) {
    uri, err := addOptions("log_entries/" + id, opt)
    if err != nil {
        return nil, nil, err
    }

    entries := new(LogEntry)
    res, err := s.client.Get(uri, entries)
    if err != nil {
        return nil, res, err
    }

    return entries, res, err
}

