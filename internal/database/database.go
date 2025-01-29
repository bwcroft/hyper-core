package database

import (
	"context"
	"fmt"

	"github.com/bwcroft/hyper-core/internal/config"
	"github.com/bwcroft/hyper-core/utils"
  "github.com/jackc/pgx/v5/pgxpool"
)

type DBConfig struct {
	User string
	Pass string
	Host string
	Name string
	Port uint16
}

func GetDBConfig() (c DBConfig) {
	c.Host = utils.GetEnvString(config.DBHost, "")
	c.User = utils.GetEnvString(config.DBUser, "")
	c.Pass = utils.GetEnvString(config.DBPass, "")
	c.Name = utils.GetEnvString(config.DBName, "")
	c.Port = utils.GetEnvUint16(config.DBPort, 5432)
	return
}

func DBConnect(c DBConfig) (db *pgxpool.Pool, err error) {
  ctx := context.Background()
	cs := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", c.User, c.Pass, c.Host, c.Port, c.Name)
	db, err = pgxpool.New(ctx, cs)
	return
}
