package main

import (
	"log"
	"time"
)

func cbwork(ticker string, size float64) {
	for {
		response, err := getBook(ticker)
		if err != nil {
			log.Printf("Error: %v", err)
			time.Sleep(3 * time.Second)
			continue
		}

		bid, ask, bps := calcSpread(response, size)

		data := InsertRow{
			time.Now().UTC(),
			bid,
			ask,
			size,
			bps,
			ticker,
			"Coinbase",
		}
		result, err := insertData(data)
		if err != nil {
			log.Println(err)
		}
		ra, _ := result.RowsAffected()
		log.Printf("Inserted %d row for LP: %s ticker: %s size: %f.\n", ra, "Coinbase", ticker, size)

		time.Sleep(60 * time.Second)
	}
}
