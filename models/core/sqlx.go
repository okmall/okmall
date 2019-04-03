package core

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// Query Mode
const (
	QueryGet uint8 = 1 + iota
	QuerySelect
	QueryExec
)

var (
	errNilQuery = fmt.Errorf("nil query")
)

// Query indicate sqlx query info
type Query interface {
	ID() uint32
	Mode() uint8
	SQL() string
	Args() []interface{}
}

// Queries indicate sqlx query slice info
type Queries interface {
	Query
	Queries() []Query
}

// Sqlx indicate sqlx exec engine
type Sqlx struct {
	*sqlx.DB
}

type dbx interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// RunQueries run query slice
func (db *Sqlx) RunQueries(q Queries) error {
	if q == nil {
		return errNilQuery
	}
	queries := q.Queries()
	if len(queries) != 0 {
		tx, err := db.Beginx()
		for _, query := range queries {
			if err = runQuery(tx, query); err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
	}
	return nil
}

// RunQuery run query by sqlx
func (db *Sqlx) RunQuery(q Query) error {
	if q == nil {
		return errNilQuery
	}
	return runQuery(db, q)
}

func runQuery(db dbx, q Query) error {
	var err error
	query := q.SQL()
	args := q.Args()
	switch q.Mode() {
	case QueryGet:
		err = db.Get(q, query, args...)
	case QuerySelect:
		err = db.Select(q, query, args...)
	case QueryExec:
		_, err = db.Exec(query, args...)
	}
	return err
}
