package main

import (
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestStdLibSQL(t *testing.T) {
	var err error

	var greeting string
	err = db.QueryRow("select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

}

func TestInsert(t *testing.T) {
	var err error

	stmt, err := db.Prepare("INSERT INTO spreads (ts,bid,ask,size,width_bps,ticker,lp) VALUES($1,$2,$3,$4,$5,$6,$7)")
	if err != nil {
		t.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	data := InsertRow{
		ts:     time.Now().UTC(),
		bid:    17123.50,
		ask:    17150.50,
		size:   1.0,
		width:  26.5,
		ticker: "BTC",
		lp:     "TEST",
	}

	result, err := stmt.Exec(data.ts, data.bid, data.ask, data.size, data.width, data.ticker, data.lp)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)

}
