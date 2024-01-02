package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB
var Stmt *sql.Stmt

// Work struct containing details of each quote pair
type Work struct {
	ID     int
	LP     string
	Ticker string
	Size   float64
}

// WorkList - list of work to do
var WorkList []Work

// Config - configuration data for application
type Config struct {
	ID       int
	Datatype string
	Key      string
	Value    string
}

// ConfigList - array of data from config db table
var ConfigList []Config

var TradeDesk = make(map[string]string)
var Assets = make(map[string]string)
var Account string

// InsertRow struct to hold data from calculations
type InsertRow struct {
	TS     time.Time
	Bid    float64
	Ask    float64
	Size   float64
	Width  float64
	Ticker string
	LP     string
}

func init() {
	var err error
	// Opening a driver typically will not attempt to connect to the database.
	DB, err = sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		fmt.Fprintf(os.Stderr, "Unable to open database: %v\n", err)
		os.Exit(1)
	}

	DB.SetConnMaxLifetime(0)
	DB.SetMaxIdleConns(10)
	// default Postgres maximum
	DB.SetMaxOpenConns(100)

	err = DB.Ping()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	Stmt, err = DB.Prepare("INSERT INTO spreads (ts,bid,ask,size,width_bps,ticker,lp) VALUES($1,$2,$3,$4,$5,$6,$7)")
	if err != nil {
		log.Fatal(err)
	}

	WorkList, err = getWork()
	if err != nil {
		log.Fatal(err)
	}

	ConfigList, err = getConfig()
	if err != nil {
		log.Fatal(err)
	}
	processConfig()

}

func InsertData(d InsertRow) (sql.Result, error) {
	result, err := Stmt.Exec(d.TS, d.Bid, d.Ask, d.Size, d.Width, d.Ticker, d.LP)
	if err != nil {
		log.Println(err)
	}

	return result, err

}

func getWork() ([]Work, error) {

	var worklist []Work

	rows, err := DB.Query("SELECT * FROM work order by ticker,size,lp")
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var w Work
		if err := rows.Scan(&w.LP, &w.Ticker, &w.Size, &w.ID); err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		worklist = append(worklist, w)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return worklist, nil
}

func getConfig() ([]Config, error) {
	var cfglist []Config

	rows, err := DB.Query("SELECT * FROM config order by data_type, key")
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var c Config
		if err := rows.Scan(&c.Datatype, &c.Key, &c.Value, &c.ID); err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		cfglist = append(cfglist, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return cfglist, nil
}

func processConfig() {
	for i, c := range ConfigList {
		log.Printf("%02d: %#v", i, c)
		switch c.Datatype {
		case "account":
			Account = c.Value
		case "assets":
			Assets[c.Key] = c.Value
		case "tradeDesk":
			TradeDesk[c.Key] = c.Value
		default:
			log.Printf("[WARNING] Not implemented config data type: %s\n", c.Datatype)
		}
	}

	log.Println("Configured trade desks:")
	for k, v := range TradeDesk {
		log.Printf("%6s: %s", k, v)
	}
	log.Println("Configured assets:")
	for k, v := range Assets {
		log.Printf("%4s: %s", k, v)
	}

}
