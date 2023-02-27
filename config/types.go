package config

type AppConfig struct {
	API_SERVER_NAME string

	APP_PORT string
	APP_ENV  string

	POSTGRES_DB       string
	POSTGRES_HOST     string
	POSTGRES_PORT     int
	POSTGRES_USER     string
	POSTGRES_PASSWORD string

	// DB              string `env:"DATABASE_DSN" envDefault:"postgres://developer:developer@localhost:5432/app?sslmode=disable"`
}
