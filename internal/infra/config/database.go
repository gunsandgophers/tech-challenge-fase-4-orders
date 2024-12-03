package config

var (
	DB_HOST     = GetEnv("DB_HOST", "mongodb")
	DB_PORT     = GetEnv("DB_PORT", "27017")
	DB_USER     = GetEnv("DB_USER", "admin")
	DB_PASSWORD = GetEnv("DB_PASSWORD", "senha123")
	DB_NAME     = GetEnv("DB_NAME", "orders")
)
