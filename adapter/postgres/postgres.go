package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

// IDB is an alias to bun.IDB.
type IDB = bun.IDB

// New returns a new instance of bun.DB
func New() *bun.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	bundb := bun.NewDB(db, pgdialect.New())
	bundb.AddQueryHook(
		bundebug.NewQueryHook(
			// disable the hook
			bundebug.WithEnabled(false),

			// BUNDEBUG=1 logs failed queries
			// BUNDEBUG=2 logs all queries
			bundebug.FromEnv("BUNDEBUG"),
		),
	)

	return bundb
}
