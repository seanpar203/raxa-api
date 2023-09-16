package db

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/seanpar203/go-api/internal/common"
)

var (
	pg     *sql.DB
	pgOnce sync.Once
)

// Returns a connection to PG database based on env variables.
func Postgres() (*sql.DB, error) {
	var err error

	pgOnce.Do(func() {

		var (
			host    = common.GetEnv("PSQL_HOST", "localhost")
			port    = common.GetEnvAsInt("PSQL_PORT", 5432)
			user    = common.GetEnv("PSQL_USER", "postgres")
			pass    = common.GetEnv("PSQL_PASS", "postgres")
			dbname  = common.GetEnv("PSQL_DBNAME", "go_api")
			sslmode = common.GetEnv("PSQL_SSLMODE", "disable")

			pgInfo = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
		)

		pg, err = sql.Open("postgres", fmt.Sprintf(pgInfo, host, port, user, pass, dbname, sslmode))
	})

	return pg, err
}
