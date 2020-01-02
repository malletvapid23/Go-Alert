package calendarsubscription

import (
	"context"
	"database/sql"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"github.com/target/goalert/permission"
	"github.com/target/goalert/util"
	"github.com/target/goalert/util/sqlutil"
	"github.com/target/goalert/validation/validate"
)

type Store struct {
	db *sql.DB
	findOne *sql.Stmt
	create  *sql.Stmt
	update  *sql.Stmt
	delete  *sql.Stmt
	findAll *sql.Stmt
	getConfig *sql.Stmt
}

type Config struct {
	NotificationMinutes []int `json:"notification_minutes"`
}

func NewStore(ctx context.Context, db *sql.DB) (*Store, error) {
	p := &util.Prepare{DB: db, Ctx: ctx}

	return &Store{
		db: db,
		findOne: p.P(`
				SELECT
					cs.id,
					cs.name,
					cs.user_id,
					cs.last_access,
					cs.disabled
				FROM user_calendar_subscriptions cs
				WHERE cs.id = $1
			`),
		create: p.P(`
				INSERT INTO user_calendar_subscriptions (user_id, id, name, config, schedule_id)
				VALUES ($1, $2, $3, $4, $5)
			`),
		update: p.P(`UPDATE user_calendar_subscriptions SET name = $2, disabled = $3, config = $4 WHERE id = $1`),
		delete: p.P(`DELETE FROM user_calendar_subscriptions WHERE id = any($1)`),
		findAll: p.P(`SELECT id,user_id,name FROM user_calendar_subscriptions WHERE user_id = $1`),
		getConfig: p.P(`SELECT config from user_calendar_subscriptions WHERE id = $1`),
	}, p.Err
}

func (b *Store) FindOne(ctx context.Context, id string) (*CalendarSubscription, error) {
	err := permission.LimitCheckAny(ctx, permission.All)
	if err != nil {
		return nil, err
	}

	err = validate.UUID("CalendarSubscriptionID", id)
	if err != nil {
		return nil, err
	}

	var cs CalendarSubscription
	err = b.findOne.QueryRowContext(ctx, id).Scan(&cs.ID, &cs.Name, &cs.UserID, &cs.LastAccess, &cs.Disabled)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

func (b *Store) CreateSubscriptionTx(ctx context.Context, tx *sql.Tx, cs *CalendarSubscription) (*CalendarSubscription, error) {
	err := permission.LimitCheckAny(ctx, permission.Admin, permission.User)
	if err != nil {
		return nil, err
	}

	cs, err = cs.Normalize()
	if err != nil {
		return nil, err
	}

	stmt := b.create
	if tx != nil {
		stmt = tx.StmtContext(ctx, stmt)
	}

	cs.ID = uuid.NewV4().String()
	cs.UserID = permission.UserID(ctx)

	_, err = stmt.ExecContext(ctx, cs.UserID, cs.ID, cs.Name, cs.Config, cs.ScheduleID)
	if err != nil {
		return nil, err
	}
	return cs, nil
}

func (b *Store) UpdateSubscriptionTx(ctx context.Context, tx *sql.Tx, cs *CalendarSubscription) (err error) {
	err = permission.LimitCheckAny(ctx, permission.Admin, permission.User)
	if err != nil {
		return err
	}
	err = validate.UUID("CalendarSubscriptionID", cs.ID)
	if err != nil {
		return err
	}
	n, err := cs.Normalize()
	if err != nil {
		return err
	}

	stmt := b.getConfig
	if tx != nil {
		stmt = tx.StmtContext(ctx, stmt)
	}
	err = stmt.QueryRowContext(ctx, n.ID).Scan(&n.Config)
	if err != nil {
		return err
	}
	var config Config
	err = json.Unmarshal(n.Config, &config)
	if err != nil {
		return err
	}
	configJson, err := json.Marshal(config)
	if err != nil {
		return err
	}


	stmt = b.update
	if tx != nil {
		stmt = tx.StmtContext(ctx, stmt)
	}

	_, err = stmt.ExecContext(ctx, n.ID, n.Name, n.Disabled, configJson)
	if err != nil {
		return err
	}

	return err
}
func (b *Store) FindAll(ctx context.Context, csID string) ([]CalendarSubscription, error) {
	err := permission.LimitCheckAny(ctx, permission.Admin, permission.User)
	if err != nil {
		return nil, err
	}
	err = validate.UUID("CalendarSubscriptionID", csID)
	if err != nil {
		return nil, err
	}

	rows, err := b.findAll.QueryContext(ctx, csID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	calendarsubscriptions := []CalendarSubscription{}
	for rows.Next() {
		var cs CalendarSubscription
		err = rows.Scan(&cs.ID, &cs.UserID, &cs.Name)
		if err != nil {
			return nil, err
		}
		calendarsubscriptions = append(calendarsubscriptions, cs)
	}

	return calendarsubscriptions, nil
}

func (b *Store) DeleteSubscriptionsTx(ctx context.Context, tx *sql.Tx, ids []string) error {
	err := permission.LimitCheckAny(ctx, permission.Admin, permission.User)
	if err != nil {
		return err
	}

	if len(ids) == 0 {
		return nil
	}

	err = validate.ManyUUID("CalendarSubscriptionID", ids, 50)
	if err != nil {
		return err
	}

	cs := b.delete
	if tx != nil {
		cs = tx.StmtContext(ctx, cs)
	}
	_, err = cs.ExecContext(ctx, sqlutil.UUIDArray(ids))
	return err
}
