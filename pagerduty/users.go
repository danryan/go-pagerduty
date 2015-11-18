package pagerduty

import (
	"errors"
	"fmt"
	"net/http"
)

// UsersService provides the API calls to interact with users
type UsersService struct {
	client *Client
}

// User type
type User struct {
	ID                string    `json:"id,omitempty"`
	Name              string    `json:"name,omitempty"`
	Email             string    `json:"email,omitempty"`
	Role              string    `json:"role,omitempty"`
	TimeZone          string    `json:"time_zone,omitempty"`
	Color             string    `json:"color,omitempty"`
	UserURL           string    `json:"user_url,omitempty"`
	AvatarURL         string    `json:"avatar_url,omitempty"`
	Description       string    `json:"description,omitempty"`
	InvitationSent    bool      `json:"invitation_sent,omitempty"`
	MarketingOptOut   bool      `json:"marketing_opt_out,omitempty"`
	ContactMethods    []Contact `json:"contact_methods,omitempty"`
	NotificationRules []Rules   `json:"notification_rules,omitempty"`
}

// Rules type
type Rules struct {
	ID                  string    `json:"id,omitempty"`
	Urgency             string    `json:"urgency,omitempty"`
	StartDelayInMinutes int       `json:"start_delay_in_minutes,omitempty"`
	CreatedAt           string    `json:"created_at,omitempty"`
	ContactMethods      []Contact `json:"contact_methods,omitempty"`
}

// Sound type
type Sound struct {
	Type string `json:"type,omitempty"`
	File string `json:"file,omitempty"`
}

// Contact type
type Contact struct {
	ID             string  `json:"id,omitempty"`
	Label          string  `json:"label,omitempty"`
	Address        string  `json:"address,omitempty"`
	UserID         string  `json:"user_id,omitempty"`
	Type           string  `json:"type,omitempty"`
	Email          string  `json:"email,omitempty"`
	SendShortEmail bool    `json:"send_short_email,omitempty"`
	CountryCode    int     `json:"country_code,omitempty"`
	PhoneNumber    string  `json:"phone_number,omitempty"`
	DeviceToken    string  `json:"device_token,omitempty"`
	DeviceType     string  `json:"device_type,omitempty"`
	Sound          string  `json:"sound,omitempty"`
	Sounds         []Sound `json:"sounds,omitempty"`
	Blacklisted    bool    `json:"blacklisted,omitempty"`
	Enabled        bool    `json:"enabled,omitempty"`
}

// Users is a list of users
type Users struct {
	Users []User
}

// UsersOptions provides optional parameters to list requests
type UsersOptions struct {
	Query   string   `url:"query,omitempty"`
	Include []string `url:"include[],omitempty"`
}

// List returns a list of users
func (s *UsersService) List(opt *UsersOptions) ([]User, *http.Response, error) {
	uri, err := addOptions("users", opt)
	if err != nil {
		return nil, nil, err
	}

	users := new(Users)
	res, err := s.client.Get(uri, users)
	if err != nil {
		return nil, res, err
	}

	return users.Users, res, err
}

// Get returns a User by id if found
func (s *UsersService) Get(id string) (*User, *http.Response, error) {
	uri := fmt.Sprintf("users/%v", id)

	user := new(User)
	res, err := s.client.Get(uri, user)
	if err != nil {
		return nil, res, err
	}

	return user, res, err
}

// FindContactMethod Seaches a users contact methods for a type
func (s *User) FindContactMethod(methodType string) (Contact, error) {

	for _, c := range s.ContactMethods {
		if c.Type == methodType {
			return c, nil
		}
	}

	return Contact{}, errors.New(fmt.Sprintf("contact of type '%s' was not found", methodType))

}
