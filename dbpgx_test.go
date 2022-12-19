package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestPGX(t *testing.T) {

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	defer dbpool.Close()
	if err != nil {
		t.Fatal("Unable to create connection pool:", err)
	}

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		t.Fatal("Unable to create connection pool:", err)
	}

	fmt.Println(greeting)
}
