package main

import (
	"log"
	"sync"
	"time"
)

// Work struct containing details of each quote pair
type Work struct {
	lp     string
	ticker string
	size   float64
}

// WorkList - list of work to do
var WorkList = []Work{
	{"Enigma", "BTC", 1.0},
	{"DV", "BTC", 1.0},
	{"Enigma", "ETH", 1.0},
	{"DV", "ETH", 1.0},
	{"DV", "USDC", 100},
	{"DV", "USDT", 100},
	{"DV", "LTC", 1.0},
	{"DV", "AVAX", 10},
	{"DV", "SOL", 10},
	{"Enigma", "BTC", 0.01},
	{"DV", "BTC", 0.01},
}

func main() {
	var err error
	defer db.Close()
	stmt, err = db.Prepare("INSERT INTO spreads (ts,bid,ask,size,width_bps,ticker,lp) VALUES($1,$2,$3,$4,$5,$6,$7)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	var mainwg sync.WaitGroup
	for i, w := range WorkList {
		log.Printf("%d: %#v", i, w)
		mainwg.Add(1)
		go func() {
			defer mainwg.Done()
			dowork(w)
		}()
		time.Sleep(127 * time.Millisecond)
	}
	log.Println("Waiting for Main WaitGroup.")
	mainwg.Wait()
	log.Println("Main WaitGroup ended.")
}
