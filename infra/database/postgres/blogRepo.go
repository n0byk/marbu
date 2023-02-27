package postgres

import (
	"context"
	"fmt"
	"os"
)

func (db *pgRepository) BlogList(ctx context.Context, id string) error {
	var greeting string
	err := db.pool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
	return nil
}
func (db *pgRepository) BlogCreate(ctx context.Context, id string) error {
	var greeting string
	err := db.pool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	return nil
}
func (db *pgRepository) BlogGet(ctx context.Context, id string) error {
	return nil
}
func (db *pgRepository) BlogUpdate(ctx context.Context, id string) error {
	return nil
}
func (db *pgRepository) BlogDelete(ctx context.Context, id string) error {
	return nil
}
