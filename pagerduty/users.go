package pagerduty

import (
	"fmt"
	"net/http"
)

// UsersService provides the API calls to interact with users
type UsersService struct {
	client *Client
}

// User type
type User struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Role      string `json:"role,omitempty"`
	TimeZone  string `json:"time_zone,omitempty"`
	Color     string `json:"color,omitempty"`
	UserURL   string `json:"user_url,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

// Users is a list of users
type Users struct {
	Users []User
}

// UsersOptions provides optional parameters to list requests
type UsersOptions struct {
	Query string `json:"query,omitempty"`
}

// List returns a list of users
func (s *UsersService) List(opt *UsersOptions) ([]User, *http.Response, error) {
	u, err := addOptions("users", opt)
	if err != nil {
		return nil, nil, err
	}

	users := new(Users)
	res, err := s.client.Get(u, users)
	if err != nil {
		return nil, res, err
	}

	return users.Users, res, err
}

// Get returns a User by id if found
func (s *UsersService) Get(id string) (*User, *http.Response, error) {
	u := fmt.Sprintf("users/%v", id)

	user := new(User)
	res, err := s.client.Get(u, user)
	if err != nil {
		return nil, res, err
	}

	return user, res, err
}
