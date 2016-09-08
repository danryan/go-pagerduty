package pagerduty_test

import (
	"fmt"
	. "github.com/danryan/go-pagerduty/pagerduty"
	"net/http"
	"reflect"
	"testing"
)

func TestUser_marshal(t *testing.T) {
	testJSONMarshal(t, &User{}, "{}")

	u := &User{
		ID:      "ABCDEF",
		Name:    "Bill Williams",
		Email:   "bill.williams@example.com",
		UserURL: "/users/ABCDEF",
		ContactMethods: []Contact{
			Contact{
				ID:          "PNB9UG1",
				Label:       "Mobile",
				Address:     "4072559655",
				UserID:      "PSXBF4M",
				Type:        "SMS",
				CountryCode: 1,
				PhoneNumber: "4072559655",
				Enabled:     true,
			},
		},
	}

	want := `{
		"id": "ABCDEF",
		"name": "Bill Williams",
		"email": "bill.williams@example.com",
		"user_url": "/users/ABCDEF",
		"contact_methods" : [
 			{
        "id": "PNB9UG1",
        "label": "Mobile",
        "address": "4072559655",
        "user_id": "PSXBF4M",
        "type": "SMS",
        "country_code": 1,
        "phone_number": "4072559655",
        "enabled": true
      }
		]
	}`

	testJSONMarshal(t, u, want)
}

func TestUsersService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/u", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id": "ABCDEF"}`)
	})

	user, _, err := client.Users.Get("u")
	if err != nil {
		t.Errorf("Users.Get returned error: %v", err)
	}

	want := &User{ID: "ABCDEF"}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("Users.Get returned %+v, want %%+v", user, want)
	}
}
