package pagerduty

import "net/http"

// ServicesService type
type ServicesService struct {
	client *Client
}

// Service type
type Service struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	HTMLURL   string `json:"html_url,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
}

// Services is a list of services
type Services struct {
	Services []Service
}

// ServicesOptions provides optional parameters to list requests
type ServicesOptions struct {
	Query string `json:"query,omitempty"`
}

// List returns a list of services
func (s *ServicesService) List(opt *ServicesOptions) ([]Service, *http.Response, error) {
	uri, err := addOptions("services", opt)
	if err != nil {
		return nil, nil, err
	}

	services := new(Services)
	res, err := s.client.Get(uri, services)
	if err != nil {
		return nil, res, err
	}

	return services.Services, res, err
}

// Get returns a single service by id if found
func (s *ServicesService) Get(id string) (*Service, *http.Response, error) {
	service := new(Service)

	res, err := s.client.Get("services/"+id, service)
	if err != nil {
		return nil, res, err
	}

	return service, res, nil
}
