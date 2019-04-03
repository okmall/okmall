package sqlite3

import (
	"github.com/jmoiron/sqlx"
	"github.com/okmall/okmall/models/core"
	"github.com/okmall/okmall/models/model"

	_ "github.com/mattn/go-sqlite3" // sqlite3 db driver
)

type sqlite3Repository struct {
	*core.Sqlx
}

// NewRepository build a new core.Repository that backend by mysql database
func NewRepository(config *model.Config) (core.Repository, error) {
	_, dsn := config.Dsn()
	sqliteDb, err := sqlx.Connect("sqlite3", dsn)
	return &sqlite3Repository{Sqlx: &core.Sqlx{DB: sqliteDb}}, err
}
