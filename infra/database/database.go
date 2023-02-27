package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/n0byk/marbu/config"
	"github.com/n0byk/marbu/infra/database/postgres"
)

func NewStorage() {
	if config.Env.POSTGRES_DB != "" {
		log.Println("Postgres storage init.")

		pgPool, err := pgxpool.New(context.Background(), "postgres://marbu:marbu@localhost:5432/marbu?sslmode=disable")
		if err != nil {
			log.Fatalf("Unable to connect to database: %s\n", err)
		}

		storage := postgres.NewRepository(pgPool)

		storage.CheckConnection()

		config.App = config.Service{Storage: storage}
		// defer pgPool.Close()
	}
}
