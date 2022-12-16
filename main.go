package main

import (
	"log"
	"sync"
	"time"
)

func init() {
	// https://pkg.go.dev/log#pkg-constants
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

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
		go func(w Work) {
			defer mainwg.Done()
			dowork(w)
		}(w)
		if i < len(WorkList)-1 && WorkList[i].ticker != WorkList[i+1].ticker {
			time.Sleep(2 * time.Second)
			log.Println("Pausing to reduce load on PT API")
		}
	}

	go cbwork("BTC", 1)

	log.Println("Waiting for Main WaitGroup.")
	mainwg.Wait()
	log.Println("Main WaitGroup ended.")
}
