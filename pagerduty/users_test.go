package pagerduty_test

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	. "git.int.enstratius.com/infraengprivate/dcm-hal/pagerduty"
)

func TestUser_marshal(t *testing.T) {
	testJSONMarshal(t, &User{}, "{}")

	u := &User{
		ID:      "ABCDEF",
		Name:    "Bill Williams",
		Email:   "bill.williams@example.com",
		UserURL: "/users/ABCDEF",
	}

	want := `{
		"id": "ABCDEF",
		"name": "Bill Williams",
		"email": "bill.williams@example.com",
		"user_url": "/users/ABCDEF"
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
