package postgre

import (
	"context"
	"fmt"
	"os"
)

func (db *pgRepository) FeedList(ctx context.Context, id string) (string, error) {
	var greeting string
	err := db.pool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
	return greeting, nil
}
