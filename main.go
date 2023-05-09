package main

import (
	"flag"
	"log"
	"sync"
)

var port = flag.String("port", "8080", "Port to listen on for web ui.")

func init() {
	// https://pkg.go.dev/log#pkg-constants
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	flag.Parse()
	defer db.Close()
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	var mainwg sync.WaitGroup
	for i, w := range WorkList {
		log.Printf("%02d: %#v", i, w)

		go func(w Work) {
			mainwg.Add(1)
			defer mainwg.Done()

			if w.LP == "Coinbase" {
				cbwork(w)
			} else {
				dowork(w)
			}

		}(w)
	}

	go webui(*port)

	// TODO: add mechanism for stopping go routines
	log.Println("Waiting for Main WaitGroup.")
	mainwg.Wait()
	log.Println("Main WaitGroup ended.")
}
