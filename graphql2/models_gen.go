// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql2

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/target/goalert/alert"
	alertlog "github.com/target/goalert/alert/log"
	"github.com/target/goalert/assignment"
	"github.com/target/goalert/escalation"
	"github.com/target/goalert/label"
	"github.com/target/goalert/notification/slack"
	"github.com/target/goalert/override"
	"github.com/target/goalert/schedule"
	"github.com/target/goalert/schedule/rotation"
	"github.com/target/goalert/schedule/rule"
	"github.com/target/goalert/service"
	"github.com/target/goalert/user"
	"github.com/target/goalert/user/contactmethod"
)

type AlertConnection struct {
	Nodes    []alert.Alert `json:"nodes"`
	PageInfo *PageInfo     `json:"pageInfo"`
}

type AlertLogEntryConnection struct {
	Nodes    []alertlog.Entry `json:"nodes"`
	PageInfo *PageInfo        `json:"pageInfo"`
}

type AlertRecentEventsOptions struct {
	Limit *int    `json:"limit"`
	After *string `json:"after"`
}

type AlertSearchOptions struct {
	FilterByStatus    []AlertStatus `json:"filterByStatus"`
	FilterByServiceID []string      `json:"filterByServiceID"`
	Search            *string       `json:"search"`
	First             *int          `json:"first"`
	After             *string       `json:"after"`
	FavoritesOnly     *bool         `json:"favoritesOnly"`
	Omit              []int         `json:"omit"`
}

type AuthSubjectConnection struct {
	Nodes    []user.AuthSubject `json:"nodes"`
	PageInfo *PageInfo          `json:"pageInfo"`
}

type ConfigHint struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type ConfigValue struct {
	ID          string     `json:"id"`
	Description string     `json:"description"`
	Value       string     `json:"value"`
	Type        ConfigType `json:"type"`
	Password    bool       `json:"password"`
}

type ConfigValueInput struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type CreateAlertInput struct {
	Summary   string  `json:"summary"`
	Details   *string `json:"details"`
	ServiceID string  `json:"serviceID"`
}

type CreateCalendarSubscriptionInput struct {
	Name string `json:"name"`
}

type CreateEscalationPolicyInput struct {
	Name        string                            `json:"name"`
	Description *string                           `json:"description"`
	Repeat      *int                              `json:"repeat"`
	Steps       []CreateEscalationPolicyStepInput `json:"steps"`
}

type CreateEscalationPolicyStepInput struct {
	EscalationPolicyID *string                `json:"escalationPolicyID"`
	DelayMinutes       int                    `json:"delayMinutes"`
	Targets            []assignment.RawTarget `json:"targets"`
	NewRotation        *CreateRotationInput   `json:"newRotation"`
	NewSchedule        *CreateScheduleInput   `json:"newSchedule"`
}

type CreateHeartbeatMonitorInput struct {
	ServiceID      string `json:"serviceID"`
	Name           string `json:"name"`
	TimeoutMinutes int    `json:"timeoutMinutes"`
}

type CreateIntegrationKeyInput struct {
	ServiceID *string            `json:"serviceID"`
	Type      IntegrationKeyType `json:"type"`
	Name      string             `json:"name"`
}

type CreateRotationInput struct {
	Name        string        `json:"name"`
	Description *string       `json:"description"`
	TimeZone    string        `json:"timeZone"`
	Start       time.Time     `json:"start"`
	Favorite    *bool         `json:"favorite"`
	Type        rotation.Type `json:"type"`
	ShiftLength *int          `json:"shiftLength"`
	UserIDs     []string      `json:"userIDs"`
}

type CreateScheduleInput struct {
	Name        string                `json:"name"`
	Description *string               `json:"description"`
	TimeZone    string                `json:"timeZone"`
	Favorite    *bool                 `json:"favorite"`
	Targets     []ScheduleTargetInput `json:"targets"`
}

type CreateServiceInput struct {
	Name                 string                        `json:"name"`
	Description          *string                       `json:"description"`
	Favorite             *bool                         `json:"favorite"`
	EscalationPolicyID   *string                       `json:"escalationPolicyID"`
	NewEscalationPolicy  *CreateEscalationPolicyInput  `json:"newEscalationPolicy"`
	NewIntegrationKeys   []CreateIntegrationKeyInput   `json:"newIntegrationKeys"`
	Labels               []SetLabelInput               `json:"labels"`
	NewHeartbeatMonitors []CreateHeartbeatMonitorInput `json:"newHeartbeatMonitors"`
}

type CreateUserContactMethodInput struct {
	UserID                  string                           `json:"userID"`
	Type                    contactmethod.Type               `json:"type"`
	Name                    string                           `json:"name"`
	Value                   string                           `json:"value"`
	NewUserNotificationRule *CreateUserNotificationRuleInput `json:"newUserNotificationRule"`
}

type CreateUserNotificationRuleInput struct {
	UserID          *string `json:"userID"`
	ContactMethodID *string `json:"contactMethodID"`
	DelayMinutes    int     `json:"delayMinutes"`
}

type CreateUserOverrideInput struct {
	ScheduleID   string    `json:"scheduleID"`
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	AddUserID    *string   `json:"addUserID"`
	RemoveUserID *string   `json:"removeUserID"`
}

type EscalationPolicyConnection struct {
	Nodes    []escalation.Policy `json:"nodes"`
	PageInfo *PageInfo           `json:"pageInfo"`
}

type EscalationPolicySearchOptions struct {
	First  *int     `json:"first"`
	After  *string  `json:"after"`
	Search *string  `json:"search"`
	Omit   []string `json:"omit"`
}

type LabelConnection struct {
	Nodes    []label.Label `json:"nodes"`
	PageInfo *PageInfo     `json:"pageInfo"`
}

type LabelKeySearchOptions struct {
	First  *int     `json:"first"`
	After  *string  `json:"after"`
	Search *string  `json:"search"`
	Omit   []string `json:"omit"`
}

type LabelSearchOptions struct {
	First      *int     `json:"first"`
	After      *string  `json:"after"`
	Search     *string  `json:"search"`
	UniqueKeys *bool    `json:"uniqueKeys"`
	Omit       []string `json:"omit"`
}

type LabelValueSearchOptions struct {
	Key    string   `json:"key"`
	First  *int     `json:"first"`
	After  *string  `json:"after"`
	Search *string  `json:"search"`
	Omit   []string `json:"omit"`
}

type PageInfo struct {
	EndCursor   *string `json:"endCursor"`
	HasNextPage bool    `json:"hasNextPage"`
}

type RotationConnection struct {
	Nodes    []rotation.Rotation `json:"nodes"`
	PageInfo *PageInfo           `json:"pageInfo"`
}

type RotationSearchOptions struct {
	First          *int     `json:"first"`
	After          *string  `json:"after"`
	Search         *string  `json:"search"`
	Omit           []string `json:"omit"`
	FavoritesOnly  *bool    `json:"favoritesOnly"`
	FavoritesFirst *bool    `json:"favoritesFirst"`
}

type ScheduleConnection struct {
	Nodes    []schedule.Schedule `json:"nodes"`
	PageInfo *PageInfo           `json:"pageInfo"`
}

type ScheduleRuleInput struct {
	ID            *string     `json:"id"`
	Start         *rule.Clock `json:"start"`
	End           *rule.Clock `json:"end"`
	WeekdayFilter []bool      `json:"weekdayFilter"`
}

type ScheduleSearchOptions struct {
	First          *int     `json:"first"`
	After          *string  `json:"after"`
	Search         *string  `json:"search"`
	Omit           []string `json:"omit"`
	FavoritesOnly  *bool    `json:"favoritesOnly"`
	FavoritesFirst *bool    `json:"favoritesFirst"`
}

type ScheduleTarget struct {
	ScheduleID string                `json:"scheduleID"`
	Target     *assignment.RawTarget `json:"target"`
	Rules      []rule.Rule           `json:"rules"`
}

type ScheduleTargetInput struct {
	ScheduleID  *string               `json:"scheduleID"`
	Target      *assignment.RawTarget `json:"target"`
	NewRotation *CreateRotationInput  `json:"newRotation"`
	Rules       []ScheduleRuleInput   `json:"rules"`
}

type SendContactMethodVerificationInput struct {
	ContactMethodID string `json:"contactMethodID"`
}

type ServiceConnection struct {
	Nodes    []service.Service `json:"nodes"`
	PageInfo *PageInfo         `json:"pageInfo"`
}

type ServiceSearchOptions struct {
	First          *int     `json:"first"`
	After          *string  `json:"after"`
	Search         *string  `json:"search"`
	Omit           []string `json:"omit"`
	FavoritesOnly  *bool    `json:"favoritesOnly"`
	FavoritesFirst *bool    `json:"favoritesFirst"`
}

type SetFavoriteInput struct {
	Target   *assignment.RawTarget `json:"target"`
	Favorite bool                  `json:"favorite"`
}

type SetLabelInput struct {
	Target *assignment.RawTarget `json:"target"`
	Key    string                `json:"key"`
	Value  string                `json:"value"`
}

type SlackChannelConnection struct {
	Nodes    []slack.Channel `json:"nodes"`
	PageInfo *PageInfo       `json:"pageInfo"`
}

type SlackChannelSearchOptions struct {
	First  *int     `json:"first"`
	After  *string  `json:"after"`
	Search *string  `json:"search"`
	Omit   []string `json:"omit"`
}

type StringConnection struct {
	Nodes    []string  `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type TimeZone struct {
	ID string `json:"id"`
}

type TimeZoneConnection struct {
	Nodes    []TimeZone `json:"nodes"`
	PageInfo *PageInfo  `json:"pageInfo"`
}

type TimeZoneSearchOptions struct {
	First  *int     `json:"first"`
	After  *string  `json:"after"`
	Search *string  `json:"search"`
	Omit   []string `json:"omit"`
}

type UpdateAlertsInput struct {
	AlertIDs  []int       `json:"alertIDs"`
	NewStatus AlertStatus `json:"newStatus"`
}

type UpdateEscalationPolicyInput struct {
	ID          string   `json:"id"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Repeat      *int     `json:"repeat"`
	StepIDs     []string `json:"stepIDs"`
}

type UpdateEscalationPolicyStepInput struct {
	ID           string                 `json:"id"`
	DelayMinutes *int                   `json:"delayMinutes"`
	Targets      []assignment.RawTarget `json:"targets"`
}

type UpdateHeartbeatMonitorInput struct {
	ID             string  `json:"id"`
	Name           *string `json:"name"`
	TimeoutMinutes *int    `json:"timeoutMinutes"`
}

type UpdateRotationInput struct {
	ID              string         `json:"id"`
	Name            *string        `json:"name"`
	Description     *string        `json:"description"`
	TimeZone        *string        `json:"timeZone"`
	Start           *time.Time     `json:"start"`
	Type            *rotation.Type `json:"type"`
	ShiftLength     *int           `json:"shiftLength"`
	ActiveUserIndex *int           `json:"activeUserIndex"`
	UserIDs         []string       `json:"userIDs"`
}

type UpdateScheduleInput struct {
	ID          string  `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	TimeZone    *string `json:"timeZone"`
}

type UpdateServiceInput struct {
	ID                 string  `json:"id"`
	Name               *string `json:"name"`
	Description        *string `json:"description"`
	EscalationPolicyID *string `json:"escalationPolicyID"`
}

type UpdateUserContactMethodInput struct {
	ID    string  `json:"id"`
	Name  *string `json:"name"`
	Value *string `json:"value"`
}

type UpdateUserInput struct {
	ID                          string    `json:"id"`
	Name                        *string   `json:"name"`
	Email                       *string   `json:"email"`
	Role                        *UserRole `json:"role"`
	StatusUpdateContactMethodID *string   `json:"statusUpdateContactMethodID"`
}

type UpdateUserOverrideInput struct {
	ID           string     `json:"id"`
	Start        *time.Time `json:"start"`
	End          *time.Time `json:"end"`
	AddUserID    *string    `json:"addUserID"`
	RemoveUserID *string    `json:"removeUserID"`
}

type UserConnection struct {
	Nodes    []user.User `json:"nodes"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

type UserOverrideConnection struct {
	Nodes    []override.UserOverride `json:"nodes"`
	PageInfo *PageInfo               `json:"pageInfo"`
}

type UserOverrideSearchOptions struct {
	First              *int       `json:"first"`
	After              *string    `json:"after"`
	Omit               []string   `json:"omit"`
	ScheduleID         *string    `json:"scheduleID"`
	FilterAddUserID    []string   `json:"filterAddUserID"`
	FilterRemoveUserID []string   `json:"filterRemoveUserID"`
	FilterAnyUserID    []string   `json:"filterAnyUserID"`
	Start              *time.Time `json:"start"`
	End                *time.Time `json:"end"`
}

type UserSearchOptions struct {
	First  *int     `json:"first"`
	After  *string  `json:"after"`
	Search *string  `json:"search"`
	Omit   []string `json:"omit"`
}

type VerifyContactMethodInput struct {
	ContactMethodID string `json:"contactMethodID"`
	Code            int    `json:"code"`
}

type AlertStatus string

const (
	AlertStatusStatusAcknowledged   AlertStatus = "StatusAcknowledged"
	AlertStatusStatusClosed         AlertStatus = "StatusClosed"
	AlertStatusStatusUnacknowledged AlertStatus = "StatusUnacknowledged"
)

var AllAlertStatus = []AlertStatus{
	AlertStatusStatusAcknowledged,
	AlertStatusStatusClosed,
	AlertStatusStatusUnacknowledged,
}

func (e AlertStatus) IsValid() bool {
	switch e {
	case AlertStatusStatusAcknowledged, AlertStatusStatusClosed, AlertStatusStatusUnacknowledged:
		return true
	}
	return false
}

func (e AlertStatus) String() string {
	return string(e)
}

func (e *AlertStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AlertStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AlertStatus", str)
	}
	return nil
}

func (e AlertStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ConfigType string

const (
	ConfigTypeString     ConfigType = "string"
	ConfigTypeStringList ConfigType = "stringList"
	ConfigTypeInteger    ConfigType = "integer"
	ConfigTypeBoolean    ConfigType = "boolean"
)

var AllConfigType = []ConfigType{
	ConfigTypeString,
	ConfigTypeStringList,
	ConfigTypeInteger,
	ConfigTypeBoolean,
}

func (e ConfigType) IsValid() bool {
	switch e {
	case ConfigTypeString, ConfigTypeStringList, ConfigTypeInteger, ConfigTypeBoolean:
		return true
	}
	return false
}

func (e ConfigType) String() string {
	return string(e)
}

func (e *ConfigType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ConfigType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ConfigType", str)
	}
	return nil
}

func (e ConfigType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type IntegrationKeyType string

const (
	IntegrationKeyTypeGeneric  IntegrationKeyType = "generic"
	IntegrationKeyTypeGrafana  IntegrationKeyType = "grafana"
	IntegrationKeyTypeSite24x7 IntegrationKeyType = "site24x7"
	IntegrationKeyTypeEmail    IntegrationKeyType = "email"
)

var AllIntegrationKeyType = []IntegrationKeyType{
	IntegrationKeyTypeGeneric,
	IntegrationKeyTypeGrafana,
	IntegrationKeyTypeSite24x7,
	IntegrationKeyTypeEmail,
}

func (e IntegrationKeyType) IsValid() bool {
	switch e {
	case IntegrationKeyTypeGeneric, IntegrationKeyTypeGrafana, IntegrationKeyTypeSite24x7, IntegrationKeyTypeEmail:
		return true
	}
	return false
}

func (e IntegrationKeyType) String() string {
	return string(e)
}

func (e *IntegrationKeyType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IntegrationKeyType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IntegrationKeyType", str)
	}
	return nil
}

func (e IntegrationKeyType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserRole string

const (
	UserRoleUnknown UserRole = "unknown"
	UserRoleUser    UserRole = "user"
	UserRoleAdmin   UserRole = "admin"
)

var AllUserRole = []UserRole{
	UserRoleUnknown,
	UserRoleUser,
	UserRoleAdmin,
}

func (e UserRole) IsValid() bool {
	switch e {
	case UserRoleUnknown, UserRoleUser, UserRoleAdmin:
		return true
	}
	return false
}

func (e UserRole) String() string {
	return string(e)
}

func (e *UserRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserRole", str)
	}
	return nil
}

func (e UserRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
