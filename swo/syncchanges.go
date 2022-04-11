package swo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// syncChanges will apply all changes recorded in the change_log table to the next DB.
func (e *Execute) syncChanges(ctx context.Context, srcTx, dstTx pgxQueryer) ([]int, error) {
	type rowID struct {
		table string
		id    string
	}

	var r rowID
	var changeIDs []int
	var changeID int
	changes := make(map[rowID]struct{})
	rowIDs := make(map[string][]string)
	_, err := srcTx.QueryFunc(ctx, "select id, table_name, row_id from change_log", nil, []interface{}{&changeID, &r.table, &r.id}, func(pgx.QueryFuncRow) error {
		if _, ok := changes[r]; ok {
			return nil
		}
		changes[r] = struct{}{}
		rowIDs[r.table] = append(rowIDs[r.table], r.id)
		changeIDs = append(changeIDs, changeID)

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fetch changes: %w", err)
	}
	if len(changes) == 0 {
		return nil, nil
	}

	var applyChanges pgx.Batch
	applyChanges.Queue("SET CONSTRAINTS ALL DEFERRED")

	type pendingDelete struct {
		query string
		idArg interface{}
		count int
	}
	var deletes []pendingDelete

	var queryChanges pgx.Batch
	var changedTables []Table
	for _, table := range e.tables {
		if table.SkipSync() {
			continue
		}
		ids := rowIDs[table.Name]
		if len(ids) == 0 {
			continue
		}

		queryChanges.Queue(table.SelectRowsQuery(), table.IDs(ids))
		changedTables = append(changedTables, table)
	}

	res := srcTx.SendBatch(ctx, &queryChanges)
	defer res.Close()

	// go in insert order for fetching updates/inserts, note deleted rows
	for _, table := range changedTables {
		sd, err := e.readChanges(ctx, table, res, rowIDs[table.Name])
		if err != nil {
			return changeIDs, fmt.Errorf("fetch changed rows: %w", err)
		}
		if len(sd.toDelete) > 0 {
			deletes = append(deletes, pendingDelete{table.DeleteRowsQuery(), table.IDs(sd.toDelete), len(sd.toDelete)})
		}

		err = e.queueChanges(&applyChanges, table.UpdateRowsQuery(), sd.toUpdate)
		if err != nil {
			return changeIDs, fmt.Errorf("apply updates: %w", err)
		}

		err = e.queueChanges(&applyChanges, table.InsertRowsQuery(), sd.toInsert)
		if err != nil {
			return changeIDs, fmt.Errorf("apply inserts: %w", err)
		}
	}

	// handle pendingDeletes in reverse table order
	for i := len(deletes) - 1; i >= 0; i-- {
		applyChanges.Queue(deletes[i].query, deletes[i].idArg)
	}

	if applyChanges.Len() == 1 {
		// no changes (just defer constraints)
		return nil, nil
	}

	err = dstTx.SendBatch(ctx, &applyChanges).Close()
	if err != nil {
		return changeIDs, fmt.Errorf("apply changes: %w", err)
	}

	return changeIDs, nil
}
