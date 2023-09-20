package db

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/seanpar203/go-api/internal/common"
)

var (
	pg     *sql.DB
	pgOnce sync.Once

	pgTest     *sql.DB
	pgTestOnce sync.Once
)

// Our config struct
type pgConfig struct {
	host    string
	port    int
	user    string
	pass    string
	dbname  string
	sslmode string

	maxOpenConns    int
	maxIdleConns    int
	connMaxLifetime time.Duration
}

// connect establishes a connection to the PostgreSQL database.
//
// It takes the PostgreSQL configuration parameters as input and returns a pointer to the sql.DB object and an error.
func (cfg *pgConfig) connect() (*sql.DB, error) {
	pgInfo := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"

	pg, err := sql.Open("postgres", fmt.Sprintf(pgInfo, cfg.host, cfg.port, cfg.user, cfg.pass, cfg.dbname, cfg.sslmode))

	if err != nil {
		return pg, err
	}

	pg.SetMaxIdleConns(cfg.maxIdleConns)
	pg.SetMaxOpenConns(cfg.maxOpenConns)
	pg.SetConnMaxLifetime(cfg.connMaxLifetime)

	return pg, nil
}

// getPgConfigFromEnv returns the PostgreSQL configuration from environment variables.
//
// No parameters.
// Returns a pgConfig struct.
func getPgConfigFromEnv() pgConfig {
	return pgConfig{
		host:    common.GetEnv("PSQL_HOST", "localhost"),
		port:    common.GetEnvAsInt("PSQL_PORT", 5432),
		user:    common.GetEnv("PSQL_USER", "postgres"),
		pass:    common.GetEnv("PSQL_PASS", "postgres"),
		dbname:  common.GetEnv("PSQL_DBNAME", "go_api"),
		sslmode: common.GetEnv("PSQL_SSLMODE", "disable"),

		maxOpenConns:    common.GetEnvAsInt("PSQL_MAX_OPEN_CONNS", 10),
		maxIdleConns:    common.GetEnvAsInt("PSQL_MAX_IDLE_CONNS", 10),
		connMaxLifetime: time.Duration(common.GetEnvAsInt("PSQL_CONN_MAX_LIFETIME", 60)) * time.Second,
	}
}


// Postgres returns a pointer to a sql.DB object and an error.
//
// The function connects to a Postgres database using the configuration obtained from the environment variables. It ensures that the connection is established only once by utilizing sync.Once. If the connection is successful, it returns the sql.DB object and a nil error. Otherwise, it returns a nil sql.DB object and the error encountered during the connection process.
func Postgres() (*sql.DB, error) {
	var err error

	pgOnce.Do(func() {
		cfg := getPgConfigFromEnv()

		pg, err = cfg.connect()
	})

	return pg, err
}

// PostgresTest is a function that returns a *sql.DB and an error. 
//
// It initializes a Postgres test database connection and returns the connection and any potential errors that occurred during the process.
// 
// Return:
// - *sql.DB: The Postgres test database connection.
// - error: Any errors that occurred during the process of initializing the connection.
func PostgresTest() (*sql.DB, error) {
	var err error

	pgTestOnce.Do(func() {
		cfg := getPgConfigFromEnv()
		cfg.dbname = "test"

		pgTest, err = cfg.connect()
	})

	return pgTest, err
}
