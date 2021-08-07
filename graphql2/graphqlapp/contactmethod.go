package graphqlapp

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"strings"

	"github.com/target/goalert/config"
	"github.com/target/goalert/graphql2"
	"github.com/target/goalert/notification"
	"github.com/target/goalert/user/contactmethod"
	"github.com/target/goalert/util/log"
	"github.com/target/goalert/validation"
	"github.com/target/goalert/validation/validate"
	"github.com/ttacon/libphonenumber"
)

type ContactMethod App

func (a *App) UserContactMethod() graphql2.UserContactMethodResolver {
	return (*ContactMethod)(a)
}

func (a *ContactMethod) Value(ctx context.Context, obj *contactmethod.ContactMethod) (string, error) {
	if obj.Type != contactmethod.TypeWebhook {
		return obj.Value, nil
	}

	u, err := url.Parse(obj.Value)
	if err != nil {
		return "", err
	}
	return maskURLPass(u), nil
}

func maskURLPass(u *url.URL) string {
	if u.User == nil {
		return u.String()
	}

	_, ok := u.User.Password()
	if !ok {
		return u.String()
	}

	u.User = url.UserPassword(u.User.Username(), "")

	parts := strings.SplitN(u.String(), "@", 2)
	parts[0] += "***"
	return strings.Join(parts, "@")
}

func (a *ContactMethod) FormattedValue(ctx context.Context, obj *contactmethod.ContactMethod) (string, error) {
	formatted, err := a.Value(ctx, obj)
	if err != nil {
		return "", err
	}

	switch obj.Type {
	case contactmethod.TypeSMS, contactmethod.TypeVoice:
		num, err := libphonenumber.Parse(formatted, "")
		if err != nil {
			log.Log(ctx, fmt.Errorf("parse number for formatting: %w", err))
			break
		}
		formatted = libphonenumber.Format(num, libphonenumber.INTERNATIONAL)
	case contactmethod.TypeWebhook:
		u, err := url.Parse(formatted)
		if err != nil {
			log.Log(ctx, fmt.Errorf("parse url for formatting: %w", err))
			break
		}
		u.RawQuery = ""
		if len(u.Path) > 15 {
			u.Path = u.Path[:12] + "..."
		}
		formatted = maskURLPass(u)
	}
	return formatted, nil
}

func (a *ContactMethod) LastTestMessageState(ctx context.Context, obj *contactmethod.ContactMethod) (*graphql2.NotificationState, error) {
	t := obj.LastTestVerifyAt()
	if t.IsZero() {
		return nil, nil
	}

	status, _, err := a.NotificationStore.LastMessageStatus(ctx, notification.MessageTypeTest, obj.ID, t)
	if err != nil {
		return nil, err
	}
	if status == nil {
		return nil, nil
	}

	return notificationStateFromStatus(status.Status), nil
}
func (a *ContactMethod) LastVerifyMessageState(ctx context.Context, obj *contactmethod.ContactMethod) (*graphql2.NotificationState, error) {
	t := obj.LastTestVerifyAt()
	if t.IsZero() {
		return nil, nil
	}

	status, _, err := a.NotificationStore.LastMessageStatus(ctx, notification.MessageTypeVerification, obj.ID, t)
	if err != nil {
		return nil, err
	}
	if status == nil {
		return nil, nil
	}

	return notificationStateFromStatus(status.Status), nil
}

func (q *Query) UserContactMethod(ctx context.Context, id string) (*contactmethod.ContactMethod, error) {
	return (*App)(q).FindOneCM(ctx, id)
}

func (m *Mutation) CreateUserContactMethod(ctx context.Context, input graphql2.CreateUserContactMethodInput) (*contactmethod.ContactMethod, error) {
	var cm *contactmethod.ContactMethod
	cfg := config.FromContext(ctx)

	if input.Type == contactmethod.TypeWebhook && !cfg.ValidWebhookURL(input.Value) {
		return nil, validation.NewFieldError("value", "URL not allowed by administrator")
	}

	err := withContextTx(ctx, m.DB, func(ctx context.Context, tx *sql.Tx) error {
		var err error
		cm, err = m.CMStore.CreateTx(ctx, tx, &contactmethod.ContactMethod{
			Name:     input.Name,
			Type:     input.Type,
			UserID:   input.UserID,
			Value:    input.Value,
			Disabled: true,
		})
		if err != nil {
			return err
		}

		if input.NewUserNotificationRule != nil {
			input.NewUserNotificationRule.UserID = &input.UserID
			input.NewUserNotificationRule.ContactMethodID = &cm.ID

			_, err = m.CreateUserNotificationRule(ctx, *input.NewUserNotificationRule)

			if err != nil {
				return validation.AddPrefix("newUserNotificationRule.", err)
			}
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	return cm, nil
}

func (m *Mutation) UpdateUserContactMethod(ctx context.Context, input graphql2.UpdateUserContactMethodInput) (bool, error) {

	err := withContextTx(ctx, m.DB, func(ctx context.Context, tx *sql.Tx) error {
		cm, err := m.CMStore.FindOneTx(ctx, tx, input.ID)
		if err != nil {
			return err
		}
		if input.Name != nil {
			cm.Name = *input.Name
		}
		if input.Value != nil {
			cm.Value = *input.Value
		}

		return m.CMStore.UpdateTx(ctx, tx, cm)
	})
	return err == nil, err
}

func (m *Mutation) SendContactMethodVerification(ctx context.Context, input graphql2.SendContactMethodVerificationInput) (bool, error) {
	err := m.NotificationStore.SendContactMethodVerification(ctx, input.ContactMethodID)
	return err == nil, err
}

func (m *Mutation) VerifyContactMethod(ctx context.Context, input graphql2.VerifyContactMethodInput) (bool, error) {
	err := validate.Range("Code", input.Code, 100000, 999999)
	if err != nil {
		// return "must be 6 digits" error as we care about # of digits, not the code's actual value
		return false, validation.NewFieldError("Code", "must be 6 digits")
	}

	err = m.NotificationStore.VerifyContactMethod(ctx, input.ContactMethodID, input.Code)
	return err == nil, err
}
