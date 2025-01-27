package main

import (
	"github.com/bwcroft/hyper-core/internal/config"
	"github.com/bwcroft/hyper-core/internal/database"
	"github.com/bwcroft/hyper-core/internal/server"
	"github.com/bwcroft/hyper-core/utils"
)

func main() {
	flags := utils.GetApiFlags()
	if err := utils.InitEnvs(config.ServerEnvs(), flags.EnvFilePath); err != nil && flags.EnvValidate {
		panic(err)
	}
	_, err := database.DBConnect(database.GetDBConfig())
	if err != nil {
		panic(err)
	}
  if err := server.StartServer(); err != nil {
    panic(err)
  }
}
