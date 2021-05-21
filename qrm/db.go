package qrm

import (
	"context"
	"database/sql"
)

// DB is common database interface used by query result mapping
// Both *sql.DB and *sql.Tx implements DB interface
type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}
