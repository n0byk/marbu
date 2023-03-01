package postgre

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/n0byk/marbu/infra/database/repository"
)

type pgRepository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) repository.Repository {
	return &pgRepository{pool}
}

func (db *pgRepository) CheckConnection() {
	if err := db.pool.Ping(context.Background()); err != nil {
		log.Fatalf("PostgreSQL connection error: %s\n", err)
	}
}

func (db *pgRepository) CloseConnection() {
	db.pool.Close()
}
