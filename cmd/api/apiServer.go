package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/n0byk/marbu/app/httpApi"
	"github.com/n0byk/marbu/config"
	"github.com/n0byk/marbu/dataservice/postgres"
)

func init() {
	config.NewConfig()

	if config.Env.POSTGRES_DB != "" {
		log.Println("Postgres storage init.")

		pgpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatalf("Unable to connect to database: %s\n", err)
		}
		postgres.NewDBRepository(pgpool)
		defer pgpool.Close()
	}
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	httpApi.NewHttpApi(r)

	if err := http.ListenAndServe(config.Env.APP_PORT, r); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Listen: %s\n", err)
	}
}
