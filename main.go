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
// loaded in init func in ./dbsql.go
var WorkList []Work

func main() {
	defer db.Close()
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	var mainwg sync.WaitGroup
	for i, w := range WorkList {
		log.Printf("%d: %#v", i, w)

		go func(w Work) {
			mainwg.Add(1)
			defer mainwg.Done()

			if w.lp == "Coinbase" {
				cbwork(w.ticker, w.size)
			}
			dowork(w)

		}(w)
	}

	// TODO: add mechanism for stopping go routines
	log.Println("Waiting for Main WaitGroup.")
	mainwg.Wait()
	log.Println("Main WaitGroup ended.")
}
