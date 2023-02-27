package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Env AppConfig

func NewConfig() {
	err := godotenv.Load()
	handleError(err)

	postgresPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	handleError(err)

	Env = AppConfig{
		API_SERVER_NAME: os.Getenv("API_SERVER_NAME"),
		APP_PORT:        os.Getenv("APP_PORT"),
		APP_ENV:         os.Getenv("APP_ENV"),

		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT:     postgresPort,
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
	}

}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
