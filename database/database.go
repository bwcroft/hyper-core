package database

import (
	"context"
	"fmt"

	"github.com/bwcroft/hypercore/env"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const (
	EnvHost = "DB_HOST"
	EnvPort = "DB_PORT"
	EnvUser = "DB_USER"
	EnvPass = "DB_PASS"
	EnvDB   = "DB_Name"
)

// MakeConnStr generates a database connection string.
//
// The function uses default values unless environment variables are set.
func MakeConnString() string {
	db := env.GetEnvString(EnvDB, "")
	host := env.GetEnvString(EnvHost, "")
	user := env.GetEnvString(EnvUser, "")
	pass := env.GetEnvString(EnvPass, "")
	port := env.GetEnvUint16(EnvPort, 5432)
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, pass, host, port, db)
}

// Connect to database
func Connect(cs string) (*sqlx.DB, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cs)
	if err != nil {
		return nil, err
	}
	db := sqlx.NewDb(stdlib.OpenDBFromPool(pool), "pgx")
	return db, nil
}
