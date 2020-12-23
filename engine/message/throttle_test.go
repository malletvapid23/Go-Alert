package message_test

import (
	"github.com/stretchr/testify/require"
	"github.com/target/goalert/engine/message"
	"github.com/target/goalert/notification"
	"testing"
	"time"
)

func TestThrottle_InCooldown(t *testing.T) {
	n := time.Now()

	var cfg = message.ThrottleConfig{
		notification.DestTypeSMS: {
			{Count: 1, Per: time.Minute},
		},
	}

	throttle := message.NewThrottle(cfg, n, false)

	msg := message.Message{
		ID:        "0",
		Type:      message.TypeAlertNotification,
		UserID:    "User A",
		ServiceID: "Service A",
		Dest:      notification.Dest{Type: notification.DestTypeSMS, ID: "SMS A"},
		SentAt:    n,
	}

	pending := throttle.InCooldown(msg)
	require.Equal(t, pending, false)

	throttle.Record(msg)

	pending = throttle.InCooldown(msg)
	require.Equal(t, true, pending)
}
