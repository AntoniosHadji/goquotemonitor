package kraken

import (
	"log"
	"time"

	"github.com/antonioshadji/goquotemonitor/db"
)

var provider = "Kraken"

func Work(w db.Work) {
	for {
		response, err := getBook(w.Ticker)
		if err != nil {
			log.Printf("Error: %v", err)
			time.Sleep(3 * time.Second)
			continue
		}

		bid, ask, bps := calcSpread(response, w.Size)

		data := db.InsertRow{
			TS:     time.Now().UTC(),
			Bid:    bid,
			Ask:    ask,
			Size:   w.Size,
			Width:  bps,
			Ticker: w.Ticker,
			LP:     provider,
		}
		result, err := db.InsertData(data)
		if err != nil {
			log.Println(err)
		}
		ra, _ := result.RowsAffected()
		log.Printf("Inserted %d row for LP: %s ticker: %s size: %f.\n", ra, provider, w.Ticker, w.Size)

		time.Sleep(60 * time.Second)
	}
}
