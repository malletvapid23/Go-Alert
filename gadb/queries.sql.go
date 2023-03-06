// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: queries.sql

package gadb

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const alertHasEPState = `-- name: AlertHasEPState :one
SELECT EXISTS (
        SELECT 1
        FROM escalation_policy_state
        WHERE alert_id = $1
    ) AS has_ep_state
`

func (q *Queries) AlertHasEPState(ctx context.Context, alertID int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, alertHasEPState, alertID)
	var has_ep_state bool
	err := row.Scan(&has_ep_state)
	return has_ep_state, err
}

const calSubAuthUser = `-- name: CalSubAuthUser :one
UPDATE user_calendar_subscriptions
SET last_access = now()
WHERE NOT disabled
    AND id = $1
    AND date_trunc('second', created_at) = $2 RETURNING user_id
`

type CalSubAuthUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
}

func (q *Queries) CalSubAuthUser(ctx context.Context, arg CalSubAuthUserParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, calSubAuthUser, arg.ID, arg.CreatedAt)
	var user_id uuid.UUID
	err := row.Scan(&user_id)
	return user_id, err
}

const createCalSub = `-- name: CreateCalSub :one
INSERT INTO user_calendar_subscriptions (
        id,
        NAME,
        user_id,
        disabled,
        schedule_id,
        config
    )
VALUES ($1, $2, $3, $4, $5, $6) RETURNING created_at
`

type CreateCalSubParams struct {
	ID         uuid.UUID
	Name       string
	UserID     uuid.UUID
	Disabled   bool
	ScheduleID uuid.UUID
	Config     json.RawMessage
}

func (q *Queries) CreateCalSub(ctx context.Context, arg CreateCalSubParams) (time.Time, error) {
	row := q.db.QueryRowContext(ctx, createCalSub,
		arg.ID,
		arg.Name,
		arg.UserID,
		arg.Disabled,
		arg.ScheduleID,
		arg.Config,
	)
	var created_at time.Time
	err := row.Scan(&created_at)
	return created_at, err
}

const deleteManyCalSub = `-- name: DeleteManyCalSub :exec
DELETE FROM user_calendar_subscriptions
WHERE id = ANY($1::uuid [ ])
    AND user_id = $2
`

type DeleteManyCalSubParams struct {
	Column1 []uuid.UUID
	UserID  uuid.UUID
}

func (q *Queries) DeleteManyCalSub(ctx context.Context, arg DeleteManyCalSubParams) error {
	_, err := q.db.ExecContext(ctx, deleteManyCalSub, pq.Array(arg.Column1), arg.UserID)
	return err
}

const findManyCalSubByUser = `-- name: FindManyCalSubByUser :many
SELECT id,
    NAME,
    user_id,
    disabled,
    schedule_id,
    config,
    last_access
FROM user_calendar_subscriptions
WHERE user_id = $1
`

type FindManyCalSubByUserRow struct {
	ID         uuid.UUID
	Name       string
	UserID     uuid.UUID
	Disabled   bool
	ScheduleID uuid.UUID
	Config     json.RawMessage
	LastAccess sql.NullTime
}

func (q *Queries) FindManyCalSubByUser(ctx context.Context, userID uuid.UUID) ([]FindManyCalSubByUserRow, error) {
	rows, err := q.db.QueryContext(ctx, findManyCalSubByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindManyCalSubByUserRow
	for rows.Next() {
		var i FindManyCalSubByUserRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.UserID,
			&i.Disabled,
			&i.ScheduleID,
			&i.Config,
			&i.LastAccess,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findOneCalSub = `-- name: FindOneCalSub :one
SELECT id,
    NAME,
    user_id,
    disabled,
    schedule_id,
    config,
    last_access
FROM user_calendar_subscriptions
WHERE id = $1
`

type FindOneCalSubRow struct {
	ID         uuid.UUID
	Name       string
	UserID     uuid.UUID
	Disabled   bool
	ScheduleID uuid.UUID
	Config     json.RawMessage
	LastAccess sql.NullTime
}

func (q *Queries) FindOneCalSub(ctx context.Context, id uuid.UUID) (FindOneCalSubRow, error) {
	row := q.db.QueryRowContext(ctx, findOneCalSub, id)
	var i FindOneCalSubRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.Disabled,
		&i.ScheduleID,
		&i.Config,
		&i.LastAccess,
	)
	return i, err
}

const findOneCalSubForUpdate = `-- name: FindOneCalSubForUpdate :one
SELECT id,
    NAME,
    user_id,
    disabled,
    schedule_id,
    config,
    last_access
FROM user_calendar_subscriptions
WHERE id = $1 FOR
UPDATE
`

type FindOneCalSubForUpdateRow struct {
	ID         uuid.UUID
	Name       string
	UserID     uuid.UUID
	Disabled   bool
	ScheduleID uuid.UUID
	Config     json.RawMessage
	LastAccess sql.NullTime
}

func (q *Queries) FindOneCalSubForUpdate(ctx context.Context, id uuid.UUID) (FindOneCalSubForUpdateRow, error) {
	row := q.db.QueryRowContext(ctx, findOneCalSubForUpdate, id)
	var i FindOneCalSubForUpdateRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.Disabled,
		&i.ScheduleID,
		&i.Config,
		&i.LastAccess,
	)
	return i, err
}

const lockOneAlertService = `-- name: LockOneAlertService :one
SELECT maintenance_expires_at notnull::bool AS is_maint_mode,
    alerts.status
FROM services svc
    JOIN alerts ON alerts.service_id = svc.id
WHERE alerts.id = $1
`

type LockOneAlertServiceRow struct {
	IsMaintMode bool
	Status      EnumAlertStatus
}

func (q *Queries) LockOneAlertService(ctx context.Context, id int64) (LockOneAlertServiceRow, error) {
	row := q.db.QueryRowContext(ctx, lockOneAlertService, id)
	var i LockOneAlertServiceRow
	err := row.Scan(&i.IsMaintMode, &i.Status)
	return i, err
}

const now = `-- name: Now :one
SELECT now()::timestamptz
`

func (q *Queries) Now(ctx context.Context) (time.Time, error) {
	row := q.db.QueryRowContext(ctx, now)
	var column_1 time.Time
	err := row.Scan(&column_1)
	return column_1, err
}

const requestAlertEscalationByTime = `-- name: RequestAlertEscalationByTime :one
UPDATE escalation_policy_state
SET force_escalation = TRUE
WHERE alert_id = $1
    AND last_escalation <= $2::timestamptz RETURNING TRUE
`

type RequestAlertEscalationByTimeParams struct {
	AlertID int64
	Column2 time.Time
}

func (q *Queries) RequestAlertEscalationByTime(ctx context.Context, arg RequestAlertEscalationByTimeParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, requestAlertEscalationByTime, arg.AlertID, arg.Column2)
	var column_1 bool
	err := row.Scan(&column_1)
	return column_1, err
}

const updateCalSub = `-- name: UpdateCalSub :exec
UPDATE user_calendar_subscriptions
SET NAME = $1,
    disabled = $2,
    config = $3,
    last_update = now()
WHERE id = $4
    AND user_id = $5
`

type UpdateCalSubParams struct {
	Name     string
	Disabled bool
	Config   json.RawMessage
	ID       uuid.UUID
	UserID   uuid.UUID
}

func (q *Queries) UpdateCalSub(ctx context.Context, arg UpdateCalSubParams) error {
	_, err := q.db.ExecContext(ctx, updateCalSub,
		arg.Name,
		arg.Disabled,
		arg.Config,
		arg.ID,
		arg.UserID,
	)
	return err
}
