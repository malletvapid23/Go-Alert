package label

import (
	"context"
	"database/sql"

	"github.com/target/goalert/assignment"
	"github.com/target/goalert/permission"
	"github.com/target/goalert/util"
	"github.com/target/goalert/validation/validate"

	"github.com/pkg/errors"
)

// Store allows the lookup and management of Labels.
type Store struct {
	db *sql.DB

	upsert           *sql.Stmt
	delete           *sql.Stmt
	findAllByService *sql.Stmt
	uniqueKeys       *sql.Stmt
}

// NewStore will Set a DB backend from a sql.DB. An error will be returned if statements fail to prepare.
func NewStore(ctx context.Context, db *sql.DB) (*Store, error) {
	p := &util.Prepare{DB: db, Ctx: ctx}
	return &Store{
		db: db,
		upsert: p.P(`
			INSERT INTO labels (tgt_service_id, key, value)
			VALUES ($1, $2, $3)
			ON CONFLICT (key, tgt_service_id) DO UPDATE
			SET value = $3
		`),
		delete: p.P(`
			DELETE FROM labels
			WHERE tgt_service_id = $1 
			AND key = $2 
		`),
		findAllByService: p.P(`
			SELECT key, value
			FROM labels
			WHERE tgt_service_id = $1
			ORDER BY key ASC
		`),
		uniqueKeys: p.P(`
			SELECT DISTINCT (key) 
			FROM labels
			ORDER BY key ASC
		`),
	}, p.Err
}

// SetTx will set a label for the service. It can be used to set the key-value pair for the label,
// delete a label or update the value given the label's key.
func (s *Store) SetTx(ctx context.Context, tx *sql.Tx, label *Label) error {
	err := permission.LimitCheckAny(ctx, permission.System, permission.User)
	if err != nil {
		return err
	}

	n, err := label.Normalize()
	if err != nil {
		return err
	}

	if n.Value == "" {
		// Delete Operation
		stmt := s.delete
		if tx != nil {
			stmt = tx.StmtContext(ctx, stmt)
		}

		_, err = stmt.ExecContext(ctx, n.Target.TargetID(), n.Key)
		return errors.Wrap(err, "delete label")
	}

	stmt := s.upsert
	if tx != nil {
		stmt = tx.StmtContext(ctx, stmt)
	}

	_, err = stmt.ExecContext(ctx, n.Target.TargetID(), n.Key, n.Value)
	if err != nil {
		return errors.Wrap(err, "set label")
	}

	return nil
}

// FindAllByService finds all labels for a particular service. It returns all key-value pairs.
func (s *Store) FindAllByService(ctx context.Context, serviceID string) ([]Label, error) {
	err := permission.LimitCheckAny(ctx, permission.System, permission.User)
	if err != nil {
		return nil, err
	}

	err = validate.UUID("ServiceID", serviceID)
	if err != nil {
		return nil, err
	}
	rows, err := s.findAllByService.QueryContext(ctx, serviceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var labels []Label
	var l Label

	for rows.Next() {
		err = rows.Scan(
			&l.Key,
			&l.Value,
		)
		if err != nil {
			return nil, errors.Wrap(err, "scan row")
		}

		l.Target = assignment.ServiceTarget(serviceID)

		labels = append(labels, l)
	}

	return labels, nil
}

func (s *Store) UniqueKeysTx(ctx context.Context, tx *sql.Tx) ([]string, error) {
	err := permission.LimitCheckAny(ctx, permission.System, permission.User)
	if err != nil {
		return nil, err
	}

	stmt := s.uniqueKeys
	if tx != nil {
		stmt = tx.StmtContext(ctx, stmt)
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var keys []string

	for rows.Next() {
		var k string
		err = rows.Scan(&k)
		if err != nil {
			return nil, errors.Wrap(err, "scan row")
		}

		keys = append(keys, k)
	}
	return keys, nil
}

func (s *Store) UniqueKeys(ctx context.Context) ([]string, error) {
	return s.UniqueKeysTx(ctx, nil)
}
