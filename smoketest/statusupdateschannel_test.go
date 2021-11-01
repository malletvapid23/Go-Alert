package smoketest

import (
	"testing"

	"github.com/target/goalert/smoketest/harness"
)

// TestStatusUpdatesChannel tests status updates to notification channels.
func TestStatusUpdatesChannel(t *testing.T) {
	t.Parallel()

	sql := `
	insert into escalation_policies (id, name) 
	values
		({{uuid "eid"}}, 'esc policy');
	insert into escalation_policy_steps (id, escalation_policy_id) 
	values
		({{uuid "esid"}}, {{uuid "eid"}});

	insert into notification_channels (id, type, name, value)
	values
		({{uuid "chan"}}, 'SLACK', '#test', {{slackChannelID "test"}});

	insert into escalation_policy_actions (escalation_policy_step_id, channel_id) 
	values 
		({{uuid "esid"}}, {{uuid "chan"}});

	insert into services (id, escalation_policy_id, name) 
	values
		({{uuid "sid"}}, {{uuid "eid"}}, 'service');
`
	h := harness.NewHarness(t, sql, "slack-user-link")
	defer h.Close()

	a := h.CreateAlertWithDetails(h.UUID("sid"), "testing", "details")
	msg := h.Slack().Channel("test").ExpectMessage("testing", "details")
	msg.AssertColor("#862421")

	h.Trigger()

	a.Close()

	updated := msg.ExpectUpdate()
	updated.AssertText("Closed", "testing")
	updated.AssertNotText("details")
	updated.AssertColor("green")
}
