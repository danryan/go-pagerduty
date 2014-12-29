package pagerduty_test

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	. "github.com/danryan/go-pagerduty/pagerduty"
)

func TestSchedule_marshal(t *testing.T) {
	testJSONMarshal(t, &Schedule{}, "{}")

	s := &Schedule{
		ID:                   "SCHED1",
		Name:                 "Primary",
		Timezone:             "Pacific Time (US & Canada)",
		Today:                "2014-12-27",
		EscalationPolicies:   []*EscalationPolicy{(&EscalationPolicy{ID: "ESCPOL1", Name: "Escalation Policy 1"})},
		ScheduleLayers:       []*ScheduleLayer{(&ScheduleLayer{Name: "Schedule Layer 1"})},
		OverridesSubschedule: &ScheduleLayer{Name: "Override Schedule Layer"},
		FinalSchedule:        &ScheduleLayer{Name: "Final Schedule Layer"},
	}

	want := `{
		"id": "SCHED1",
		"name": "Primary",
                "time_zone": "Pacific Time (US & Canada)",
                "today": "2014-12-27",
                "escalation_policies": [{"id": "ESCPOL1", "name": "Escalation Policy 1"}],
                "schedule_layers": [{"name": "Schedule Layer 1"}],
                "overrides_subschedule": {"name": "Override Schedule Layer"},
                "final_schedule": {"name": "Final Schedule Layer"}
	}`

	testJSONMarshal(t, s, want)
}

func TestSchedulesService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/schedules", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"schedules": [{"id": "SCHED1", "name": "Schedule 1"}], "limit": 100, "offset": 0, "total": 1}`)
	})

	schedules, _, err := client.Schedules.List(&SchedulesOptions{})
	if err != nil {
		t.Errorf("Schedules.List returned error: %v", err)
	}

	want := &Schedules{Schedules: []*Schedule{(&Schedule{ID: "SCHED1", Name: "Schedule 1"})}, Limit: 100, Offset: 0, Total: 1}
	if !reflect.DeepEqual(schedules, want) {
		t.Errorf("Schedules.List returned %+v, want %%+v", schedules, want)
	}
}

func TestSchedulesService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/schedules/s", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"schedule": {"id": "ABCDEF"}}`)
	})

	schedule, _, err := client.Schedules.Get("s")
	if err != nil {
		t.Errorf("Schedules.Get returned error: %v", err)
	}

	want := &Schedule{ID: "ABCDEF"}
	if !reflect.DeepEqual(schedule, want) {
		t.Errorf("Schedules.Get returned %+v, want %%+v", schedule, want)
	}
}

func TestSchedulesService_Entries(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/schedules/s/entries", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"total": 1, "entries": [{"user": {"id": "USR1", "name": "Test User"}, "end": "2012-08-19T12:00:00-04:00", "start": "2012-08-19T00:00:00-04:00"}]}`)
	})

	entries, _, err := client.Schedules.Entries("s", &ScheduleEntriesOptions{})
	if err != nil {
		t.Errorf("Schedules.Entries returned error: %v", err)
	}

	want := &ScheduleEntries{ScheduleEntries: []*ScheduleEntry{(&ScheduleEntry{User: &User{ID: "USR1", Name: "Test User"}, Start: "2012-08-19T00:00:00-04:00", End: "2012-08-19T12:00:00-04:00"})}, Total: 1}
	if !reflect.DeepEqual(entries, want) {
		t.Errorf("Schedules.Entries returned %+v, want %%+v", entries, want)
	}
}
