package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestStdLibSQL(t *testing.T) {
	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	var greeting string
	err = db.QueryRow("select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}

func TestInsert(t *testing.T) {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	fmt.Printf("drivers: %v\n", sql.Drivers())

	stmt, err := db.Prepare("INSERT INTO spreads (ts,bid,ask,size,width_bps,ticker,lp) VALUES($1,$2,$3,$4,$5,$6,$7)")
	if err != nil {
		t.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	ts := time.Now().UTC()
	bid := 17123.50
	ask := 17150.50
	size := 1.0
	width := 26.5
	ticker := "BTC"
	lp := "DV"

	result, err := stmt.Exec(ts, bid, ask, size, width, ticker, lp)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}
