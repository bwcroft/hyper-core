package config

const (
	DBHost     = "DB_HOST"
	DBUser     = "DB_USER"
	DBPass     = "DB_PASS"
	DBName     = "DB_NAME"
	DBPort     = "DB_PORT"
	JWTSecret  = "JWT_SECRET"
	ServerPort = "SERVER_PORT"
)

func ServerEnvs() []string {
	return []string{
		DBHost,
		DBUser,
		DBPass,
		DBName,
		DBPort,
	}
}
