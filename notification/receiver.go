package notification

import (
	"context"
	"errors"
)

// A Receiver processes incoming messages and responses.
type Receiver interface {
	// SetMessageStatus can be used to update the state of a message.
	SetMessageStatus(ctx context.Context, externalID string, status *Status) error

	// Receive records a response to a previously sent message.
	Receive(ctx context.Context, callbackID string, result Result) error

	// ReceiveSubject records a response to a previously sent message from a provider/subject (e.g. Slack user).
	ReceiveSubject(ctx context.Context, providerID, subjectID, callbackID string, result Result) error

	// Start indicates a user has opted-in for notifications to this contact method.
	Start(context.Context, Dest) error

	// Stop indicates a user has opted-out of notifications from a contact method.
	Stop(context.Context, Dest) error

	// IsKnownDest checks if the given destination is known/not disabled.
	IsKnownDest(ctx context.Context, value string) (bool, error)
}

// ErrUnknownSubject is returned from ReceiveSubject when the subject is unknown.
var ErrUnknownSubject = errors.New("unknown subject for that provider")
