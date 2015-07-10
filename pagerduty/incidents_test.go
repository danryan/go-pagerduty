package pagerduty_test

import (
	"fmt"
	"net/http"
	"testing"
	. "."
)

func TestIncident_empty_marshal(t *testing.T) {
	testJSONMarshal(t, &Incident{}, "{}")
}

func TestIncident_sample_marshal(t *testing.T) {
	t.SkipNow()
	i := &Incident{
		ID:             "ABCDEF",
		IncidentNumber: 1,
		Status:         "resolved",
		CreatedOn:      "2014-08-25T19:11:30Z",
		HTMLURL:        "https://example.pagerduty.com/incidents/ABCDEF",
		IncidentKey:    "79574e08c4bdeaf9d1cdf1f059eba93",
		SService: &Service{
			ID:      "ABCDEF",
			Name:    "Test",
			HTMLURL: "https://example.pagerduty.com/services/ABCDEF",
		},
		EEscalationPolicy: &EscalationPolicy{
			ID:   "ABCDEF",
			Name: "Default",
		},
		AssignedToUser: nil,
		AssignedTo:     make([]*User, 0),
	}

	want := `{
	  "id": "ABCDEF",
	  "incident_number": 1,
	  "created_on": "2014-08-25T19:11:30Z",
	  "status": "resolved",
	  "html_url": "https://example.pagerduty.com/incidents/ABCDEF",
	  "incident_key": "79574e08c4bdeaf9d1cdf1f059eba93",
	  "service": {
	    "id": "ABCDEF",
	    "name": "Test",
	    "html_url": "https://example.pagerduty.com/services/ABCDEF",
	    "deleted_at": null
	  },
	  "escalation_policy": {
	    "id": "ABCDEF",
	    "name": "Default",
	    "deleted_at": null
	  },
	  "assigned_to_user": null,
	  "trigger_summary_data": {
	    "subject": "Test"
	  },
	  "trigger_details_html_url": "https://example.pagerduty.com/incidents/ABCDEF/log_entries/ABCDEF",
	  "trigger_type": "web_trigger",
	  "last_status_change_on": "2014-08-25T20:02:47Z",
	  "last_status_change_by": {
	    "id": "ABCDEF",
	    "name": "Bill Williams",
	    "email": "bill.williams@example.com",
	    "html_url": "https://example.pagerduty.com/users/ABCDEF"
	  },
	  "number_of_escalations": 0,
	  "resolved_by_user": {
	    "id": "ABCDEF",
	    "name": "Bill Williams",
	    "email": "bill.williams@example.com",
	    "html_url": "https://example.pagerduty.com/users/ABCDEF"
	  },
	  "assigned_to": []
	}`

	testJSONMarshal(t, i, want)
}

func TestIncidentsService_Get(t *testing.T) {
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

	want := &User {
		ID       : "ABCDEF",
	}
	if user.ID != want.ID {
		t.Errorf("Users.Get returned:\n%+v\nWant: \n%+v\n", user, want)
	}
}
