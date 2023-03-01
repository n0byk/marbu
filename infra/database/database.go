package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/n0byk/marbu/config"
	"github.com/n0byk/marbu/infra/database/postgre"
)

func New() {
	if config.Env.POSTGRES_DB != "" {
		log.Println("PostgreSQL storage init.")

		pgPool, err := pgxpool.New(context.Background(), "postgres://marbu:marbu@localhost:5432/marbu?sslmode=disable")
		if err != nil {
			log.Fatalf("Unable to connect to database: %s\n", err)
		}

		storage := postgre.NewRepository(pgPool)

		storage.CheckConnection()

		config.App = config.Service{Storage: storage}
	}
}
