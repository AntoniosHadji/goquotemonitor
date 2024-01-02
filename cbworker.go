package main

import (
	"log"
	"time"

	"github.com/antonioshadji/goquotemonitor/db"
)

func cbwork(w db.Work) {
	for {
		response, err := getBook(w.Ticker)
		if err != nil {
			log.Printf("Error: %v", err)
			time.Sleep(3 * time.Second)
			continue
		}

		bid, ask, bps := calcSpread(response, w.Size)

		// TODO: unkeyed struct warning
		data := db.InsertRow{
			time.Now().UTC(),
			bid,
			ask,
			w.Size,
			bps,
			w.Ticker,
			"Coinbase",
		}
		result, err := db.InsertData(data)
		if err != nil {
			log.Println(err)
		}
		ra, _ := result.RowsAffected()
		log.Printf("Inserted %d row for LP: %s ticker: %s size: %f.\n", ra, "Coinbase", w.Ticker, w.Size)

		time.Sleep(60 * time.Second)
	}
}
