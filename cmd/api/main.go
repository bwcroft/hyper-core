package main

import (
	"github.com/bwcroft/hyper-core/internal/config"
	"github.com/bwcroft/hyper-core/internal/database"
	"github.com/bwcroft/hyper-core/internal/routes"
	"github.com/bwcroft/hyper-core/utils"
)

func main() {
  /** Env Validation */
	flags := utils.GetApiFlags()
	if err := utils.InitEnvs(config.ServerEnvs(), flags.EnvFilePath); err != nil && flags.EnvValidate {
		panic(err)
	}

  /** Database Connection */
	db, err := database.Connect(database.GetConfig())
	if err != nil {
		panic(err)
	}
  defer db.Close()

  /** Http Server */
	port := utils.GetEnvUint16(config.ServerPort, 8080)
  server := routes.InitRoutes(db)
	if err := server.Listen(port); err != nil {
		panic(err)
	}
}
