package metricsmanager

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/target/goalert/engine/processinglock"
	"github.com/target/goalert/util"
)

const engineVersion = 1

// DB handles updating metrics
type DB struct {
	db   *sql.DB
	lock *processinglock.Lock

	setTimeout *sql.Stmt

	findState            *sql.Stmt
	updateState          *sql.Stmt
	findMaxAlertID       *sql.Stmt
	findMinClosedAlertID *sql.Stmt
	findRecentAlert      *sql.Stmt
	insertAlertMetrics   *sql.Stmt
}

// Name returns the name of the module.
func (db *DB) Name() string { return "Engine.MetricsManager" }

// NewDB creates a new DB.
func NewDB(ctx context.Context, db *sql.DB) (*DB, error) {
	lock, err := processinglock.NewLock(ctx, db, processinglock.Config{
		Version: engineVersion,
		Type:    processinglock.TypeMetrics,
	})
	if err != nil {
		return nil, err
	}

	p := &util.Prepare{Ctx: ctx, DB: db}

	return &DB{
		db:   db,
		lock: lock,

		// Abort any cleanup operation that takes longer than 3 seconds
		// error will be logged.
		setTimeout: p.P(`SET LOCAL statement_timeout = 3000`),

		findState: p.P(fmt.Sprintf(`select state -> 'V%d' from engine_processing_versions where type_id = 'metrics'`, engineVersion)),

		updateState: p.P(fmt.Sprintf(`update engine_processing_versions set state = jsonb_set(state, '{V%d}', $1, true) where type_id = 'metrics'`, engineVersion)),

		findMaxAlertID: p.P(`select max(id) from alerts`),

		findMinClosedAlertID: p.P(`select min(id) from alerts where status = 'closed'`),

		findRecentAlert: p.P(`
		select 
			max(a.alert_id) 
		from alert_logs a left join alert_metrics m on m.alert_id = a.id
		where m isnull and event = 'closed' and timestamp > now() - '1 hour'::interval
		`),

		insertAlertMetrics: p.P(`insert into alert_metrics 
		select
			a.id,
			a.service_id,
			(select timestamp - a.created_at from alert_logs where alert_id = a.id and event = 'acknowledged' order by timestamp limit 1),
			(select timestamp - a.created_at from alert_logs where alert_id = a.id and event = 'closed'       order by timestamp limit 1),
			(select count(*) > 1             from alert_logs where alert_id = a.id and event = 'escalated')
		from alerts a
		left join alert_metrics m on m.alert_id = a.id
		where m isnull and a.id between $1 and $2 and a.status='closed'`),
	}, p.Err
}
