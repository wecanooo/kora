package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	Exists(ctx context.Context, tableName string, field string, value string) bool
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new store
func NewSQLStore(db *sql.DB) Store {
	return &SQLStore{
		db: db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

const exists = `
SELECT COUNT(id) as count
FROM $1
WHERE $2 = $3
`

func (store *SQLStore) Exists(ctx context.Context, tableName string, field string, value string) bool {
	rows := store.db.QueryRowContext(ctx, exists, tableName, field, value)
	var count int
	err := rows.Scan(&count)
	if err != nil {
		return false
	}

	if count > 0 {
		return true
	}

	return false
}
