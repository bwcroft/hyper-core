package database

import (
	"context"
	"fmt"

	"github.com/bwcroft/hyper-core/internal/config"
	"github.com/bwcroft/hyper-core/utils"
  "github.com/jackc/pgx/v5/pgxpool"
  "github.com/jackc/pgx/v5/stdlib"
  "github.com/jmoiron/sqlx"
)

type DBConfig struct {
	User string
	Pass string
	Host string
	Name string
	Port uint16
}

func GetConfig() (c DBConfig) {
	c.Host = utils.GetEnvString(config.DBHost, "")
	c.User = utils.GetEnvString(config.DBUser, "")
	c.Pass = utils.GetEnvString(config.DBPass, "")
	c.Name = utils.GetEnvString(config.DBName, "")
	c.Port = utils.GetEnvUint16(config.DBPort, 5432)
	return
}

func Connect(c DBConfig) (*sqlx.DB, error) {
  ctx := context.Background()
	cs := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", c.User, c.Pass, c.Host, c.Port, c.Name)
  pool, err := pgxpool.New(ctx, cs)
  if err != nil {
    return nil, err
  }
  db := sqlx.NewDb(stdlib.OpenDBFromPool(pool), "pgx")
  return db, nil
}
