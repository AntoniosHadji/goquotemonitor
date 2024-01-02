package main

import (
	"flag"
	"log"
	"sync"
	"time"

	"github.com/antonioshadji/goquotemonitor/db"
)

var port = flag.String("port", "8080", "Port to listen on for web ui.")

func init() {
	// https://pkg.go.dev/log#pkg-constants
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	flag.Parse()
	defer db.DB.Close()
	defer db.Stmt.Close() // Prepared statements take up server resources and should be closed after use.

	var mainwg sync.WaitGroup
	for i, w := range db.WorkList {
		log.Printf("%02d: %#v", i, w)

		go func(w db.Work) {
			mainwg.Add(1)
			defer mainwg.Done()

			if w.LP == "Coinbase" {
				cbwork(w)
				// TODO: finish work to integrate sFOX quotes ~7bps on new years day
				// } else if w.LP == "sFOX" {
				// 	sfox.Work(w.LP)
			} else {
				dowork(w)
			}

		}(w)
		// separate quotes to reduce load
		time.Sleep(1 * time.Second)
	}

	// TODO: work in progress - display config settings and edit
	go db.Webui(*port)

	// TODO: add mechanism for stopping go routines
	log.Println("Waiting for Main WaitGroup.")
	mainwg.Wait()
	log.Println("Main WaitGroup ended.")
}
