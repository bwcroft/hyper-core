package main

import (
	"github.com/bwcroft/hyper-core/internal/config"
	"github.com/bwcroft/hyper-core/internal/database"
	"github.com/bwcroft/hyper-core/internal/server"
	"github.com/bwcroft/hyper-core/utils"
)

func main() {
  /** Env Validation */
	flags := utils.GetApiFlags()
	if err := utils.InitEnvs(config.ServerEnvs(), flags.EnvFilePath); err != nil && flags.EnvValidate {
		panic(err)
	}

  /** Database Connection */
	db, err := database.DBConnect(database.GetDBConfig())
	if err != nil {
		panic(err)
	}
  defer db.Close()

  /** Http Server */
	if err := server.StartServer(db); err != nil {
		panic(err)
	}
}
