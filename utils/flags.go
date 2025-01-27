package utils

import "flag"

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
