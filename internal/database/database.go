package database

import (
	"context"
	"fmt"

	"github.com/bwcroft/hyper-core/internal/config"
	"github.com/bwcroft/hyper-core/utils"
	"github.com/jackc/pgx/v5"
)

type DBConfig struct {
	User string
	Pass string
	Host string
	Name string
	Port uint16
}

func GetDBConfig() (c DBConfig) {
	c.Host = utils.GetEnvString(config.DBHost, "localhost")
	c.User = utils.GetEnvString(config.DBUser, "root")
	c.Pass = utils.GetEnvString(config.DBPass, "root")
	c.Name = utils.GetEnvString(config.DBName, "hypercore")
	c.Port = utils.GetEnvUint16(config.DBPort, 5432)
	return
}

func DBConnect(c DBConfig) (db *pgx.Conn, err error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", c.User, c.Pass, c.Host, c.Port, c.Name)
	db, err = pgx.Connect(context.Background(), url)
	return
}
