package main

import (
	"log"
	"sync"
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
// TODO: move this to database, with access via http server
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
	defer db.Close()
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	var mainwg sync.WaitGroup
	for i, w := range WorkList {
		log.Printf("%d: %#v", i, w)

		go func(w Work) {
			mainwg.Add(1)
			defer mainwg.Done()
			dowork(w)
		}(w)
	}

	go func() {
		mainwg.Add(1)
		defer mainwg.Done()
		cbwork("BTC", 1)
	}()

	// TODO: add mechanism for stopping go routines
	log.Println("Waiting for Main WaitGroup.")
	mainwg.Wait()
	log.Println("Main WaitGroup ended.")
}
