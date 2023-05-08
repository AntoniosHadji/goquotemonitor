package main

import (
	"log"
	"time"
)

func cbwork(w Work) {
	for {
		response, err := getBook(w.ticker)
		if err != nil {
			log.Printf("Error: %v", err)
			time.Sleep(3 * time.Second)
			continue
		}

		bid, ask, bps := calcSpread(response, w.size)

		data := InsertRow{
			time.Now().UTC(),
			bid,
			ask,
			w.size,
			bps,
			w.ticker,
			"Coinbase",
		}
		result, err := insertData(data)
		if err != nil {
			log.Println(err)
		}
		ra, _ := result.RowsAffected()
		log.Printf("Inserted %d row for LP: %s ticker: %s size: %f.\n", ra, "Coinbase", w.ticker, w.size)

		time.Sleep(60 * time.Second)
	}
}
