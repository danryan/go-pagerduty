package pagerduty

// EscalationPolicy type
type EscalationPolicy struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
}
