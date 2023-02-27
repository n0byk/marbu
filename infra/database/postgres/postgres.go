package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/n0byk/marbu/infra/repository"
)

type pgRepository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) repository.Repository {
	return &pgRepository{pool}
}
