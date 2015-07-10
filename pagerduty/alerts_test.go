package pagerduty_test

import (
	"testing"
	. "."
)

func TestAlerts_empty_marshal(t *testing.T) {
	testJSONMarshal(t, &Alerts{}, "{}")
}

func TestAlerts_sample_marshal(t *testing.T) {
	i := &Alerts{
	 	Pagination: &Pagination{
			Limit: 100,
			Offset: 0,
			Total: 1,
		},
		Alerts: []Alert {
			Alert {
				ID: "PWL7QXS",
				Type: "Phone",
				StartedAt: "2013-03-06T15:28:50-05:00",
				Address: "+15555551234",
				UUser: User {
					ID       : "PT23IWX",
					Name     : "Tim Wright",
					Email    : "tim@acme.com",
					Role     : "owner",
					TimeZone : "Eastern Time (US & Canada)",
					Color    : "purple",
					UserURL  : "/users/PT23IWX",
					AvatarURL: "https://secure.gravatar.com/avatar/923a2b907dc04244e9bb5576a42e70a7.png?d=mm&r=PG",
					InvitationSent: false,
					MarketingOptOut: false,
},
			},
		},
	}

	want := `{
  "limit": 100,
  "offset": 0,
  "total": 1,
  "alerts": [
    {
      "id": "PWL7QXS",
      "type": "Phone",
      "started_at": "2013-03-06T15:28:50-05:00",
      "user": {
        "id": "PT23IWX",
        "name": "Tim Wright",
        "email": "tim@acme.com",
        "role": "owner",
        "time_zone": "Eastern Time (US \u0026 Canada)",
        "color": "purple",
        "user_url": "/users/PT23IWX",
        "avatar_url": "https://secure.gravatar.com/avatar/923a2b907dc04244e9bb5576a42e70a7.png?d=mm\u0026r=PG",
        "invitation_sent": false,
        "marketing_opt_out": false
      },
      "address": "+15555551234"
    }
  ]
}`
	testJSONMarshal(t, i, want)
}

