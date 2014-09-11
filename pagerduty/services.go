package pagerduty

// Service type
type Service struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	HTMLURL   string `json:"html_url,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
}
