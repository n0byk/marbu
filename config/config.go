package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Config AppConfig

func InitConfig() {
	err := godotenv.Load()
	handleError(err)

	appPort, err := strconv.Atoi(os.Getenv("PORT"))
	handleError(err)

	postgresPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	handleError(err)

	Config = AppConfig{
		API_SERVER_NAME: os.Getenv("API_SERVER_NAME"),
		PORT:            appPort,
		ENV:             os.Getenv("ENV"),

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
		os.Exit(1)
	}
}
