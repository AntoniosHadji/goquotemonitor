package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB
var stmt *sql.Stmt

// InsertRow struct to hold data from calculations
type InsertRow struct {
	ts     time.Time
	bid    float64
	ask    float64
	size   float64
	width  float64
	ticker string
	lp     string
}

func init() {
	var err error
	// Opening a driver typically will not attempt to connect to the database.
	db, err = sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		fmt.Fprintf(os.Stderr, "Unable to open database: %v\n", err)
		os.Exit(1)
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	stmt, err = db.Prepare("INSERT INTO spreads (ts,bid,ask,size,width_bps,ticker,lp) VALUES($1,$2,$3,$4,$5,$6,$7)")
	if err != nil {
		log.Fatal(err)
	}

	WorkList, err = getWork()
	if err != nil {
		log.Fatal(err)
	}
}

func insertData(d InsertRow) (sql.Result, error) {
	result, err := stmt.Exec(d.ts, d.bid, d.ask, d.size, d.width, d.ticker, d.lp)
	if err != nil {
		log.Println(err)
	}

	return result, err

}

func getWork() ([]Work, error) {

	var worklist []Work

	rows, err := db.Query("SELECT * FROM work order by ticker,size,lp")
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var w Work
		if err := rows.Scan(&w.lp, &w.ticker, &w.size); err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		worklist = append(worklist, w)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return worklist, nil
}
