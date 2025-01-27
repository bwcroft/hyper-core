package main

import (
	"flag"

	"github.com/bwcroft/hyper-core/internal/config"
	"github.com/bwcroft/hyper-core/internal/database"
	"github.com/bwcroft/hyper-core/internal/server"
	"github.com/bwcroft/hyper-core/utils"
)

type ApiFlags struct {
	EnvFilePath string
	EnvValidate bool
}

func GetApiFlags() (flags ApiFlags) {
	envFilePath := flag.String("env", "", "path to load env from")
	envValidate := flag.Bool("env-validate", true, "validate server envs")
	flag.Parse()
	flags.EnvFilePath = *envFilePath
	flags.EnvValidate = *envValidate
	return
}

func main() {
	flags := GetApiFlags()

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
