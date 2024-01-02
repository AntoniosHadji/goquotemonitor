package db

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
	err = DB.QueryRow("select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

}

func TestInsert(t *testing.T) {
	var err error

	Stmt, err := DB.Prepare("INSERT INTO spreads (ts,bid,ask,size,width_bps,ticker,lp) VALUES($1,$2,$3,$4,$5,$6,$7)")
	if err != nil {
		t.Fatal(err)
	}
	defer Stmt.Close() // Prepared statements take up server resources and should be closed after use.

	data := InsertRow{
		TS:     time.Now().UTC(),
		Bid:    17123.50,
		Ask:    17150.50,
		Size:   1.0,
		Width:  26.5,
		Ticker: "XXX",
		LP:     "TEST",
	}

	result, err := Stmt.Exec(data.TS, data.Bid, data.Ask, data.Size, data.Width, data.Ticker, data.LP)
	if err != nil {
		t.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		//t.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		//t.Fatal(err)
	}
	fmt.Printf("id: %d , rows: %d\n", id, rows)

	result, err = DB.Exec("DELETE FROM spreads WHERE lp = $1;", "TEST")
	if err != nil {
		t.Fatal("Failed delete query.", err)
	}
	id, err = result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		//t.Fatal(err)
	}
	rows, err = result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		//t.Fatal(err)
	}
	fmt.Printf("id: %d , rows: %d\n", id, rows)

}

func TestQuerySelect(t *testing.T) {

	var worklist []Work

	rows, err := DB.Query("SELECT * FROM work order by ticker,size,lp")
	if err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var w Work
		if err := rows.Scan(&w.LP, &w.Ticker, &w.Size); err != nil {
			fmt.Println(fmt.Errorf("%v", err))
		}
		worklist = append(worklist, w)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}

	fmt.Println(worklist)

	for i, item := range worklist {
		fmt.Println(i, item)
	}
}
