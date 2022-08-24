package smoketest

import (
	"testing"
	"time"

	"github.com/target/goalert/test/smoketest/harness"
)

// TestOnCallNotifyTOD will test that time-of-day on-call notifications work.
func TestOnCallNotifyTOD(t *testing.T) {
	t.Parallel()

	sql := `
	insert into users (id, name, email)
	values
		({{uuid "uid"}}, 'bob', 'bob@example.com');

	insert into schedules (id, name, time_zone) 
	values
		({{uuid "sid"}}, 'testschedule', 'UTC');

	insert into schedule_rules (id, schedule_id, sunday, monday, tuesday, wednesday, thursday, friday, saturday, start_time, end_time, tgt_user_id)
	values
		({{uuid "ruleID"}}, {{uuid "sid"}}, true, true, true, true, true, true, true, '00:00:00', '00:00:00', {{uuid "uid"}});

	insert into notification_channels (id, type, name, value)
	values
		({{uuid "chan1"}}, 'SLACK', '#test1', {{slackChannelID "test1"}}),
		({{uuid "chan2"}}, 'SLACK', '#test2', {{slackChannelID "test2"}});
	
	insert into schedule_data (schedule_id, data)
	values
		({{uuid "sid"}}, '{"V1":{"OnCallNotificationRules": [{"ChannelID": {{uuidJSON "chan1"}}, "Time": "00:00" },{"ChannelID": {{uuidJSON "chan1"}}, "Time": "01:00" },{"ChannelID": {{uuidJSON "chan2"}}, "Time": "00:00" }]}}');
`
	h := harness.NewHarness(t, sql, "outgoing-messages-schedule-id")
	defer h.Close()

	h.Trigger()

	h.FastForward(24 * time.Hour)

	// should only send 1 message to each channel
	h.Slack().Channel("test1").ExpectMessage("on-call", "testschedule", "bob")
	h.Slack().Channel("test2").ExpectMessage("on-call", "testschedule", "bob")
}
