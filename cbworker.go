package main

import (
	"log"
	"time"
)

func cbwork(w Work) {
	for {
		response, err := getBook(w.Ticker)
		if err != nil {
			log.Printf("Error: %v", err)
			time.Sleep(3 * time.Second)
			continue
		}

		bid, ask, bps := calcSpread(response, w.Size)

		data := InsertRow{
			time.Now().UTC(),
			bid,
			ask,
			w.Size,
			bps,
			w.Ticker,
			"Coinbase",
		}
		result, err := insertData(data)
		if err != nil {
			log.Println(err)
		}
		ra, _ := result.RowsAffected()
		log.Printf("Inserted %d row for LP: %s ticker: %s size: %f.\n", ra, "Coinbase", w.Ticker, w.Size)

		time.Sleep(60 * time.Second)
	}
}
