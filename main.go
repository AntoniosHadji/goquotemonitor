package main

import (
	"log"
	"sync"
)

func init() {
	// https://pkg.go.dev/log#pkg-constants
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	defer db.Close()
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	var mainwg sync.WaitGroup
	for i, w := range WorkList {
		log.Printf("%02d: %#v", i, w)

		go func(w Work) {
			mainwg.Add(1)
			defer mainwg.Done()

			if w.lp == "Coinbase" {
				cbwork(w.ticker, w.size)
			} else {
				dowork(w)
			}

		}(w)
	}

	// TODO: add mechanism for stopping go routines
	log.Println("Waiting for Main WaitGroup.")
	mainwg.Wait()
	log.Println("Main WaitGroup ended.")
}
