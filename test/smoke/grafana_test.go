package smoke

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/target/goalert/test/smoke/harness"
)

func TestGrafana(t *testing.T) {
	t.Parallel()

	const sql = `
	insert into users (id, name, email)
	values
		({{uuid "user"}}, 'bob', 'joe');

	insert into user_contact_methods (id, user_id, name, type, value)
	values
		({{uuid "cm1"}}, {{uuid "user"}}, 'personal', 'SMS', {{phone "1"}});

	insert into user_notification_rules (user_id, contact_method_id, delay_minutes)
	values
		({{uuid "user"}}, {{uuid "cm1"}}, 0);

	insert into escalation_policies (id, name)
	values
		({{uuid "eid"}}, 'esc policy');

	insert into escalation_policy_steps (id, escalation_policy_id)
	values
		({{uuid "esid"}}, {{uuid "eid"}});

	insert into escalation_policy_actions (escalation_policy_step_id, user_id)
	values
		({{uuid "esid"}}, {{uuid "user"}});

	insert into services (id, escalation_policy_id, name)
	values
		({{uuid "sid"}}, {{uuid "eid"}}, 'service');

	insert into integration_keys (id, type, name, service_id)
	values
		({{uuid "int_key"}}, 'grafana', 'my key', {{uuid "sid"}});
`
	h := harness.NewHarness(t, sql, "ids-to-uuids")
	defer h.Close()

	url := h.URL() + "/v1/webhooks/grafana?integration_key=" + h.UUID("int_key")

	resp, err := http.Post(url, "application/json", bytes.NewBufferString(`
		{
			"ruleName": "bob",
			"ruleId": 1,
			"message": "test",
			"state": "alerting",
			"title": "woot",
			"ruleUrl": "dontcare"
		}
	`))
	if err != nil {
		t.Fatal("post to grafana endpoint failed:", err)
	} else if resp.StatusCode != 200 {
		t.Error("non-200 response:", resp.Status)
	}
	resp.Body.Close()

	h.Twilio(t).Device(h.Phone("1")).ExpectSMS("bob")
}
